package router

import (
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getlist", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.GET("/user/updateUser", service.UpdateUser)
	r.GET("/user/loginVal", service.LoginValidate)
	r.GET("/user/sendMsg", service.SendMsg)

	return r
}
