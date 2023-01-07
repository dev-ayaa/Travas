package db

import (
	"context"
	"time"

	"github.com/travas-io/travas/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app config.Tools

func Connection(uri string) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions))
	if err != nil {
		app.ErrorLogger.Panicln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}

	return client
}
