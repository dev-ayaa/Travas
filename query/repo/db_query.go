package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/travas-io/travas/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// database queries is done in this file

func (td *TravasDB) InsertUser(user model.Tourist) (int, primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "email", Value: user.Email}}

	var res bson.M
	err := TouristsData(td.DB, "tourists").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			user.ID = primitive.NewObjectID()
			_, insertErr := TouristsData(td.DB, "tourists").InsertOne(ctx, user)
			if insertErr != nil {
				td.App.ErrorLogger.Fatalf("cannot add user to the database : %v ", insertErr)
			}
			return 0, user.ID, nil
		}
		td.App.ErrorLogger.Fatal(err)
	}

	id := (res["_id"]).(primitive.ObjectID)

	return 1, id, err
}

func (td *TravasDB) CheckForUser(userID primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result bson.M

	filter := bson.D{{Key: "_id", Value: userID}}
	err := TouristsData(td.DB, "tourists").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
			return false, err
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}

	return true, nil
}

func (td *TravasDB) UpdateInfo(userID primitive.ObjectID, tk map[string]string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: tk["t1"]}, {Key: "new_token", Value: tk["t2"]}}}}

	_, err := TouristsData(td.DB, "tourists").UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (td *TravasDB) LoadTourPackage(res []model.Tour) ([]model.Tour, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := ToursData(td.DB, "tours").Find(ctx, bson.D{{}})
	if err != nil {
		return res, fmt.Errorf("cannot find document in the database %v ", err)
	}

	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (td *TravasDB) AddTourPackage(userID primitive.ObjectID, tour model.Tour) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "tour_list", Value: tour}}}}

	_, err := TouristsData(td.DB, "tourists").UpdateOne(ctx, filter, update)
	if err != nil {
		return false, fmt.Errorf("cannot update document : %v ", err)
	}
	return true, nil
}

func (td *TravasDB) UpdateTourPlans(userID, tourID primitive.ObjectID, tag []model.TaggedTourist) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "tagged_tourist", Value: tag}}}}

	// Add all the tagged tourist details in the main tourist database
	_, err := TouristsData(td.DB, "tourists").UpdateOne(ctx, filter, update)
	if err != nil {
		return false, fmt.Errorf("cannot update document : %v ", err)
	}

	// Update and add the number of tourist in the tour database for use in the tour operator
	filter_tour := bson.D{{Key: "_id", Value: tourID}}
	update_tour := bson.D{{Key: "$set", Value: bson.D{{Key: "number_of_tourist", Value: len(tag)}}}}
	_, err = ToursData(td.DB, "tours").UpdateOne(ctx, filter_tour, update_tour)
	if err != nil {
		return false, fmt.Errorf("cannot update document : %v ", err)
	}
	return true, nil
}

// Tours Operators Query Start Here
func (td *TravasDB) InsertTour(tour model.Tour) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var id primitive.ObjectID

	filter := bson.D{{Key: "tour_title", Value: tour.TourTitle}}
	var res bson.M
	err := ToursData(td.DB, "tours").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			tour.ID = primitive.NewObjectID()
			doc := bson.D{
				{Key: "_id", Value: tour.ID},
				{Key: "operator_id", Value: tour.OperatorID},
				{Key: "tour_title", Value: tour.TourTitle},
				{Key: "meeting_point", Value: tour.MeetingPoint},
				{Key: "start_time", Value: tour.StartTime},
				{Key: "language_offered", Value: tour.LanguageOffered},
				{Key: "number_of_tourist", Value: tour.NumberOfTourist},
				{Key: "description", Value: tour.Description},
				{Key: "tour_guide", Value: tour.TourGuide},
				{Key: "tour_operator", Value: tour.TourOperator},
				{Key: "operator_contact", Value: tour.OperatorContact},
				{Key: "date", Value: tour.Date},
			}
			_, insertErr := ToursData(td.DB, "tours").InsertOne(ctx, doc)
			if insertErr != nil {
				td.App.ErrorLogger.Fatalf("cannot add user to the database : %v ", insertErr)
			}
			return tour.ID, nil
		}
		td.App.ErrorLogger.Fatal(err)
	}

	for k, v := range res {
		if k == "_id" {
			switch tourID := v.(type) {
			case primitive.ObjectID:
				id = tourID
			}
		}
	}
	return id, err
}

func (td *TravasDB) DeleteTour(tourID primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result bson.M

	filter := bson.D{{Key: "_id", Value: tourID}}
	_, err := ToursData(td.DB, "tours").DeleteOne(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
			return false, err
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}
	td.App.InfoLogger.Printf("found document %v", result)
	return true, err
}

func (td *TravasDB) UpdateTour(tourID primitive.ObjectID, tour *model.Tour) *model.Tour {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: tour.ID}}
	err := ToursData(td.DB, "tours").FindOneAndReplace(ctx, filter, tour, &options.FindOneAndReplaceOptions{})
	if err != nil {
		return tour
	}
	return nil
}
func (td *TravasDB) GetTour(tourID primitive.ObjectID) (tour model.Tour, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: tourID}}
	err = ToursData(td.DB, "tours").FindOne(ctx, filter).Decode(&tour)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}
	return tour, err
}
func (td *TravasDB) FindAllTours() (tours []model.Tour, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := ToursData(td.DB, "tours").Find(ctx, bson.D{{}})
	if err != nil {
		return tours, fmt.Errorf("cannot find document in the database %v ", err)
	}

	if err = cursor.All(ctx, &tours); err != nil {
		return nil, err
	}

	return tours, err

}
