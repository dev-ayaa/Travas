package query

import (
	"github.com/travas-io/travas/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// todo -> all our queries method are to implement the interface

type TravasDBRepo interface {
	InsertUser(user model.Tourist) (int, primitive.ObjectID, error)
	CheckForUser(userID primitive.ObjectID) (bool, error)
	UpdateInfo(userID primitive.ObjectID, tk map[string]string) (bool, error)
	LoadTourPackage(res []model.Tour) ([]model.Tour, error) 
	AddTourPackage(userID primitive.ObjectID, tour model.Tour) (bool, error)
	UpdateTourPlans(userID, tourID primitive.ObjectID, tag []model.TaggedTourist) (bool, error)	
	
	//InsertTour(tour model.Tour) (primitive.ObjectID, error)
	//DeleteTour(tourID string) (bool, error)
	//UpdateTour(tourID string, tour model.Tour) (bool, error)
	//GetTour(tourID string) (tour model.Tour, err error)
	//FindAllTours() (tours []model.Tour, err error)
}
