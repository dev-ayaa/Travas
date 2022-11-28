package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"github.com/gorilla/mux"
	"github.com/travas-io/travas/pkg/controller"
)

func Routes(r *gin.Engine, t controller.Travas) {
	router := r.Use(gin.Logger(), gin.Recovery())
	router.Use(cors.Default())
	cookieData := cookie.NewStore([]byte("travas"))
	router.Use(sessions.Sessions("session", cookieData))
	router.GET("/", t.Welcome())
	router.GET("/api/register", t.Register())
	router.POST("/api/user/register", t.ProcessRegister())
	router.GET("/api/user/sign-in", t.LoginPage())
	router.POST("/api/user/sign-in", t.ProcessLogin())
	//router.POST("/api/tour/create", t.CreateTour())
	//router.DELETE("/api/tour/delete/:id", t.DeleteTour())
	//router.PUT("/api/tour/update/:id", t.UpdateTour())
	//router.GET("/api/tour/tours", t.GetAllTours())
	//router.GET("/api/tour/:id", t.GetTour())
	protectRouter := r.Group("/api/auth")
	protectRouter.Use(Authorization())
	{
		protectRouter.GET("/user/home", t.UserMainPage())
		protectRouter.GET("/user/select/package", t.SelectTour())
		protectRouter.POST("/user/book/package", t.BookTour())

	}

}
