package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tourist struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required,alpha" json:"name,omitempty"`
	Email     string             `json:"email" Usage:"required,alphanumeric"`
	Password  string             `json:"password" Usage:"required"`
	Address   string             `json:"address"`
	Tours     []ReservedTour     `json:"tours"`
	Token     string             `json:"token" Usage:"jwt"`
	NewToken  string             `json:"new_token" Usage:"jwt"`
	CreatedAt time.Time          `json:"created_at" Usage:"datetime"`
	UpdatedAt time.Time          `json:"updated_at" Usage:"datetime"`
}

type ReservedTour struct {
	//Selected tour will be added here
	ID       primitive.ObjectID `json:"_id"`
	TourList map[string]string  `json:"tour_list"`
}
