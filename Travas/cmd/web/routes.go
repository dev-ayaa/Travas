package main

import (
	"github.com/gin-gonic/gin"
	"github.com/travas-io/travas/Travas/pkg/controller"
)

func Routes(r *gin.Engine, h controller.TravasHandler){
	router :=r.Use(gin.Logger(), gin.Recovery())
	router.
}
