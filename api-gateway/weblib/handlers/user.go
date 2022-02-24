package handlers

import "C"
import (
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context)  {
	var userReq services.UserRequest
	PanicIfUserError(C.Bind(&userReq))
	userService:=c.Keys["userService"].(services.UserService)
	userReqMERR，err:=userService.UserRegister((context.Background(),&userReq))
	PanicIfUserError
}