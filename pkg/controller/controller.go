package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/travas-io/travas/model"
	"github.com/travas-io/travas/pkg/config"
	"github.com/travas-io/travas/pkg/hash"
	"github.com/travas-io/travas/pkg/token"
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

func (tr *Travas) Welcome() gin.HandlerFunc {
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

func (tr *Travas) LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
func (tr *Travas) ProcessLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.Request.ParseForm(); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}
		email := ctx.Request.Form.Get("email")
		password := ctx.Request.Form.Get("password")

		data := tr.App.Session.Get(ctx.Request.Context(), "data").(model.IntraData)
		verified, err := hash.Verify(password, data.Password)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot verify user input password"))
		}
		if verified {
			switch {
			case email == data.Email:
				_, checkErr := tr.DB.CheckForUser(data.ID)

				if checkErr != nil {
					_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("unregistered user %v", checkErr))
				}
				//generate the jwt token
				t1, t2, err := token.Generate(data.Email, data.ID)
				if err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token no generated : %v ", err))

				}

				var tk map[string]string
				tk = map[string]string{"t1": t1, "t2": t2}

				//update the database adding the token to user database
				_, updateErr := tr.DB.UpdateInfo(data.ID, tk)
				if updateErr != nil {
					_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("unregistered user %v", updateErr))

				}
				ctx.Header("Content-Type", "application/json")
				ctx.Header("Authorization", fmt.Sprintf("Bearer,%v", t1))

				ctx.Redirect(http.StatusSeeOther, "/user/home")
			}
		}

	}
}

func (tr *Travas) Main() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
