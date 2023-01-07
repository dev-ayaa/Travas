package repo

import (
	"context"
	"fmt"
	"github.com/travas-io/travas/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (td *TravasDB) SavePayment(system model.PaymentSystem) (int, primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "ID", Value: system.ID}}

	var res bson.M
	err := PaymentData(td.DB, "tourists").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			system.ID = primitive.NewObjectID()
			doc := bson.D{
				{Key: "_id", Value: system.ID},
				{Key: "is_successful", Value: system.IsSuccessful},
				{Key: "date_payed", Value: system.DateExpired},
				{Key: "amount_payed", Value: system.AmountPayed},
				{Key: "date_payed", Value: system.DatePayed},
			}
			_, insertErr := TouristsData(td.DB, "tourists").InsertOne(ctx, doc)
			if insertErr != nil {
				td.App.ErrorLogger.Fatalf("cannot add payment to the database : %v ", insertErr)
			}
			return 0, system.ID, nil
		}
		td.App.ErrorLogger.Fatal(err)
	}

	id := (res["_id"]).(primitive.ObjectID)

	return 1, id, err
}

func (td *TravasDB) SetPayment(systemID primitive.ObjectID) (payment *model.PaymentSystem, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: systemID}}
	err = PaymentData(td.DB, "tours").FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			td.App.ErrorLogger.Println("no document found for this query")
		}
		td.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}
	return payment, err

}

func (td *TravasDB) DeletePayment(systemID primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var result bson.M

	filter := bson.D{{Key: "_id", Value: systemID}}
	_, err := PaymentData(td.DB, "tours").DeleteOne(ctx, filter)
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

func (td *TravasDB) FindAllPayment() (tours []model.PaymentSystem, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := PaymentData(td.DB, "payments").Find(ctx, bson.D{{}})
	if err != nil {
		return tours, fmt.Errorf("cannot find document in the database %v ", err)
	}

	if err = cursor.All(ctx, &tours); err != nil {
		return nil, err
	}

	return tours, err

}
