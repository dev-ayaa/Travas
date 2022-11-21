package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/travas-io/travas/pkg/token"
	"net/http"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var authToken string
		value := ctx.GetHeader("Authorization")
		if value == "" {
			_ = ctx.AbortWithError(http.StatusNoContent, errors.New("no value for authorization header"))
		}
		valSlices := strings.Split(value, ",")
		if valSlices[0] == "Bearer" && len(valSlices[1]) > 1 {
			authToken = valSlices[1]
		}
		parse, err := token.Parse(authToken)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusUnauthorized, gin.Error{Err: err})
		}
		ctx.Set("pass", authToken)
		ctx.Set("id", parse.ID)
		ctx.Set("email", parse.Email)
		ctx.Next()
	}
}

func LoadAndSave(next http.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session.LoadAndSave(next)
	}
}
