package main

import (
	"github.com/gin-gonic/gin"
	"github.com/travas-io/travas/pkg/controller"
)

func Routes(r *gin.Engine, t controller.Travas) {
	router := r.Use(gin.Logger(), gin.Recovery())

	r.Use(LoadAndSave(r.Handler()))
	router.GET("/", t.Welcome())
	router.GET("/api/register", t.Register())
	router.POST("/api/user/register", t.ProcessRegister())
	router.GET("/api/user/sign-in", t.LoginPage())
	router.POST("/api/user/sign-in", t.ProcessLogin())

	protectRouter := r.Group("api/auth")
	protectRouter.Use(Authorization())
	{
		protectRouter.GET("/user/home", t.Main())
	}
	return
}
