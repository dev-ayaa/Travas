package repo

import (
	"context"
	"github.com/travas-io/travas/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// database queries is done in this file

func (td *TravasDB) InsertUser(user model.Tourist) (int, primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var id primitive.ObjectID

	filter := bson.D{{"email", user.Email}}
	var res bson.M
	err := TouristData(td.DB, "tourist").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			user.ID = primitive.NewObjectID()
			doc := bson.D{
				{"email", user.Email},
				{"first_name", user.FirstName},
				{"last_name", user.LastName},
				{"phone", user.Phone},
				{"password", user.Password},
			}
			_, insertErr := TouristData(td.DB, "tourist").InsertOne(ctx, doc)
			if insertErr != nil {
				td.App.ErrorLogger.Fatalf("cannot add user to the database : %v ", insertErr)
			}
			return 0, user.ID, nil
		}
		td.App.ErrorLogger.Fatal(err)
	}

	for k, v := range res {
		if k == "_id" {
			switch userID := v.(type) {
			case primitive.ObjectID:
				id = userID
			}
		}

	}
	return 1, id, err

}
