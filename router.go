package main

import (
	"BAT-douyin/controller/userhandler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	douyin := r.Group("/douyin")
	{
		user := douyin.Group("/user")
		{
			user.GET("/", userhandler.GetMes)
			user.POST("/register/", userhandler.Register)
			user.POST("/login/", userhandler.Login)
		}
	}
}
