package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/travas-io/travas/model"
	"github.com/travas-io/travas/pkg/config"
	"github.com/travas-io/travas/pkg/hash"
	"github.com/travas-io/travas/query"
	"github.com/travas-io/travas/query/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"reflect"
)

type Travas struct {
	App *config.Tools
	DB  query.TravasDBRepo
}

func NewTravas(app *config.Tools, db *mongo.Client) *Travas {
	return &Travas{
		App: app,
		DB:  repo.NewTravasDB(app, db),
	}
}

// todo -> this is where all our handler/ controller logic will be done

func (tr *Travas) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Todo : render the home page of the application
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (tr *Travas) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
func (tr *Travas) ProcessRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.Tourist

		if err := ctx.Request.ParseForm(); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}
		user.Email = ctx.Request.Form.Get("email")
		user.FirstName = ctx.Request.Form.Get("first_name")
		user.LastName = ctx.Request.Form.Get("last_name")
		user.Phone = ctx.Request.Form.Get("phone")
		user.Password, _ = hash.Encrypt(ctx.Request.Form.Get("password"))
		user.CheckPassword, _ = hash.Encrypt(ctx.Request.Form.Get("check_password"))

		ok := reflect.DeepEqual(user.Password, user.CheckPassword)
		if !ok {
			_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("passwords did not match"))
		}

		if err := tr.App.Validator.Struct(&user); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); !ok {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
				log.Println(err)
				return
			}
		}
		track, userID, err := tr.DB.InsertUser(user)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("error while adding new user"))
			return
		}

		ctx.Header("Content-Type", "application/json")

		data := model.IntraData{
			ID:       userID,
			Email:    user.Email,
			Password: user.Password,
		}
		tr.App.Session.Put(ctx.Request.Context(), "data", data)

		switch track {
		case 1:
			//add the user id to session
			//redirect to the home page of the application
			ctx.JSON(http.StatusSeeOther, gin.H{
				"message": "Existing Account, Go to the Login page",
			})
		case 0:
			//	after inserting new user to the database
			//  notify the user to verify their  details via mail
			//  OR
			//  Send notification message on the page for them to login
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Registered Successfully",
			})
		}

	}
}
