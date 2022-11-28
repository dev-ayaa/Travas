package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/travas-io/travas/model"
	"github.com/travas-io/travas/pkg/config"
	"github.com/travas-io/travas/pkg/hash"
	"github.com/travas-io/travas/pkg/token"
	"github.com/travas-io/travas/query"
	"github.com/travas-io/travas/query/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// Welcome : This method render the welcome page of the flutter mobile application
func (tr *Travas) Welcome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Todo : render the home page of the application
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// Register : this Handler will render and show the register page for user
func (tr *Travas) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// ProcessRegister : As the name implies , this method will help to process all the registration process
// of the user
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
		user.Password = ctx.Request.Form.Get("password")
		user.CheckPassword = ctx.Request.Form.Get("check_password")
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Token = ""
		user.NewToken = ""
		user.BookedTours = []model.Tour{}
		user.TaggedTourist = []model.TaggedTourist{}
		user.RequestTours = []model.Tour{}

		if user.Password != user.CheckPassword {
			_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("passwords did not match"))
		}

		user.Password, _ = hash.Encrypt(user.Password)
		user.CheckPassword, _ = hash.Encrypt(user.CheckPassword)

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
		cookieData := sessions.Default(ctx)

		userInfo := model.UserInfo{
			ID:       userID,
			Email:    user.Email,
			Password: user.Password,
		}
		cookieData.Set("info", userInfo)

		if err := cookieData.Save(); err != nil {
			log.Println("error from the session storage")
			_ = ctx.AbortWithError(http.StatusNotFound, gin.Error{Err: err})
			return
		}
		switch track {
		case 1:
			// add the user id to session
			// redirect to the home page of the application
			ctx.JSON(http.StatusOK, gin.H{
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

// LoginPage : this will show the login page for user
func (tr *Travas) LoginPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

// ProcessLogin : this method will help to parse, verify, and as well as authenticate the user
// login details, and it also helps to generate jwt token for restricted routers

func (tr *Travas) ProcessLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.Request.ParseForm(); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}
		email := ctx.Request.Form.Get("email")
		password := ctx.Request.Form.Get("password")

		cookieData := sessions.Default(ctx)
		userInfo := cookieData.Get("info").(model.UserInfo)

		verified, err := hash.Verify(password, userInfo.Password)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot verify user input password"))
		}
		if verified {
			switch {
			case email == userInfo.Email:
				_, checkErr := tr.DB.CheckForUser(userInfo.ID)

				if checkErr != nil {
					_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("unregistered user %v", checkErr))
				}
				// generate the jwt token
				t1, t2, err := token.Generate(userInfo.Email, userInfo.ID)
				if err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("token no generated : %v ", err))
				}

				var tk map[string]string
				tk = map[string]string{"t1": t1, "t2": t2}

				// update the database adding the token to user database
				_, updateErr := tr.DB.UpdateInfo(userInfo.ID, tk)
				if updateErr != nil {
					_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("unregistered user %v", updateErr))
				}

				ctx.SetCookie("authorization", t1, 60*60*24*7, "/", "localhost", false, true)
				ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to user homepage"})
			}
		}
	}
}

// UserMainPage : what this method does :  it extracts data from the tour operator database for all the tour
// packages added by the tour operator to be display on the user home/main pages for user to view all
// available packages

func (tr *Travas) UserMainPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tourOp []model.Tour

		// validate the struct tags of the Tours model
		if err := tr.App.Validator.Struct(&tourOp); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); !ok {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
				log.Println(err)
				return
			}
		}
		// call the database queries to fetch all the tour packages
		data, err := tr.DB.LoadTourPackage(tourOp)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("error no data found %v", err))
			return
		}

		if len(data) >= 1 {
			// if tour packages are available
			ctx.JSON(http.StatusOK, gin.H{"tours": data})
		} else {
			// else if tour package data is not available
			ctx.JSON(http.StatusOK, gin.H{"tour": "No data available"})
		}
	}
}

func (tr *Travas) SelectTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (tr *Travas) BookTour() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	Get all the information from the frontend
		var tour model.Tour
		err := json.NewDecoder(ctx.Request.Body).Decode(&tour)
		if err != nil {
			tr.App.ErrorLogger.Fatalf("no values in the request body : %v ", err)
		}
		// put the tour info in session
		cookieData := sessions.Default(ctx)
		cookieData.Set("booked_tour", tour)
		cookieData.Set("tour_id", tour.ID)

		if err := cookieData.Save(); err != nil {
			log.Println("error from the session storage")
			_ = ctx.AbortWithError(http.StatusNotFound, gin.Error{Err: err})
			return
		}

		userInfo := cookieData.Get("info").(model.UserInfo)
		userID := userInfo.ID

		// add the data to the tourist database
		ok, err := tr.DB.AddTourPackage(userID, tour)
		if !ok {
			_ = ctx.AbortWithError(http.StatusNotModified, fmt.Errorf("error cannot add user tour package in db : %v ", err))
		}

		ctx.JSON(http.StatusOK, gin.H{"message": " You have successfully booked for a tour packages"})
	}
}

func (tr *Travas) ProcessCheckOut () gin.HandlerFunc{
	return func(ctx *gin.Context){
		if err := ctx.Request.ParseForm(); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}
		var taggedTourist []model.TaggedTourist
		err := json.NewDecoder(ctx.Request.Body).Decode(&taggedTourist)
		if err != nil {
			tr.App.ErrorLogger.Fatalf("no values in the request body : %v ", err)
		}

		cookieData := sessions.Default(ctx)

		userInfo := cookieData.Get("info").(model.UserInfo)
		tourID:= cookieData.Get("tour_id").(primitive.ObjectID)
		userID := userInfo.ID
		
		
		ok, _ :=  tr.DB.UpdateTourPlans(userID, tourID, taggedTourist)
		if !ok{
			_ = ctx.AbortWithError(http.StatusNotModified, fmt.Errorf("error cannot number of tagged tourist : %v ", err))
		}

		ctx.JSON(http.StatusOK, gin.H{"message": " You have successfully booked for a tour packages"})

	}
}