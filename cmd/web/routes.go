package main

import (
	"github.com/gin-gonic/gin"
	"github.com/travas-io/travas/pkg/controller"
)

func Routes(r *gin.Engine, c controller.TravasHandler) {
	router := r.Use(gin.Logger(), gin.Recovery())
	router.GET("/", c.HomePage())
	return
}
