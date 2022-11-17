package repo

import (
	"github.com/travas-io/travas/pkg/config"
	"github.com/travas-io/travas/query"
	"go.mongodb.org/mongo-driver/mongo"
)

type TravasDB struct {
	App *config.TravasConfig
	DB  *mongo.Client
}

func NewTravasDB(app *config.TravasConfig, db *mongo.Client) query.TravasDBRepo {
	return &TravasDB{
		App: app,
		DB:  db,
	}
}
