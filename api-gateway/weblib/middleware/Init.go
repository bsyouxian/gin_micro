package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 接受服务实例，并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make([map[string]]interface{})
		context.Keys["userService"] = service[0]
		context.Next()
	}
}
//处理错误的中间件
func ErrMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover();r!=nil{
				context.JSON(200,gin.H{
					"code":404,
					"msg":fmt.Sprintf("%s",r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}