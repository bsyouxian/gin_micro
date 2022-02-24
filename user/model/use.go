package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PassWordDigest string
}

const (
	PassWordCost=12
)
//加密密码
func (u *User) SetPassWord(PassWord string) error {
	byte,err:=bcrypt.GenerateFromPassword([]byte(PassWord),PassWordCost)
	if err != nil {
		panic(err)
	}
	u.PassWordDigest=string(byte)
	return nil
}
//检验密码
func (u *User) CheckPassWord(PassWord string) bool {
	err:=bcrypt.CompareHashAndPassword([]byte(u.PassWordDigest),[]byte(PassWord))
	return err==nil
}