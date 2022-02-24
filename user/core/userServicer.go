package core

import (
	"errors"
	"context"
	"github.com/jinzhu/gorm"
	"user/model"
	"user/servies"
)

func (u *UserServicer) UserLogin(ctx context.Context, req *servies.UserRequest, resp *servies.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassWord(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
func BuildUser(item model.User) *servies.UserModel {
	userModel := servies.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
func (*UserServicer) UserRegister(ctx context.Context, req *servies.UserRequest, resp *servies.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	connt := 0
	if err := model.DB.Model(&model.User{}).Where("user_Name=?", req.UserName).Count(&connt).Error; err != nil {
		return err //查到了相同的ID 会让COUNT+1
	}
	if connt > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	//加密密码
if err :=user.SetPassWord(req.Password);err!=nil{
	return err
}
if err:= model.DB.Create(&user).Error;err!=nil{
	return err
}
resp.UserDetail=BuildUser(user)
	return nil
}
