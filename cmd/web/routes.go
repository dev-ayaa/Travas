package main

import (
	"github.com/gin-gonic/gin"
	"github.com/travas-io/travas/pkg/controller"
)

func Routes(r *gin.Engine, t controller.Travas) {
	router := r.Use(gin.Logger(), gin.Recovery())
	router.GET("/", t.Home())
	router.GET("/api/register", t.Register())
	router.POST("/api/user/register", t.ProcessRegister())
	router.GET("/api/user/sign-in", t.Login())
	return
}
