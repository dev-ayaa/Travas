package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/travas-io/travas/Travas/pkg/config"
	"github.com/travas-io/travas/Travas/query"
	"github.com/travas-io/travas/Travas/query/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

type TravasHandler struct {
	App *config.TravasConfig
	DB  query.TravasDBRepo
}

func NewTravasHandler(app *config.TravasConfig, db *mongo.Client) *TravasHandler {
	return &TravasHandler{
		App: app,
		DB:  repo.NewTravasDB(app, db),
	}
}

// todo -> this is where all our handler/ controller logic will be done

func (tv *TravasHandler) HomePage() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
