package repo

import "go.mongodb.org/mongo-driver/mongo"

// todo -> this file is where the application will be connected to collection in database

func TouristsData(db *mongo.Client, collection string) *mongo.Collection {
	var touristData = db.Database("travasdb").Collection(collection)
	return touristData
}

func OperatorsData(db *mongo.Client, collection string) *mongo.Collection {
	var operatorData = db.Database("travasdb").Collection(collection)
	return operatorData
}

func ToursData(db *mongo.Client, collection string) *mongo.Collection {
	var tourData = db.Database("travasdb").Collection(collection)
	return tourData
}