package db

import (
	"context"

	"github.com/travas-io/travas/pkg/config"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app config.Tools

func Connection(uri string) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions))
	/*
		if err != nil {
			app.ErrorLogger.Panicln(err)
		}
	*/
	_ = client.Ping(ctx, nil)
	/*
		if err != nil {
			app.ErrorLogger.Fatalln(err)
		}
	*/
	return client
}
