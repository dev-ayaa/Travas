package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Name          string             `json:"name" Usage:"required,alpha" json:"name,omitempty"`
	Email         string             `json:"email" Usage:"required,alphanumeric"`
	Password      string             `json:"password" Usage:"required"`
	Address       string             `json:"address"`
	ReservedTours []Tour             `json:"tours_list"`
	Token         string             `json:"token" Usage:"jwt"`
	NewToken      string             `json:"new_token" Usage:"jwt"`
	CreatedAt     time.Time          `json:"created_at" Usage:"datetime"`
	UpdatedAt     time.Time          `json:"updated_at" Usage:"datetime"`
}

type Tour struct {
	ID              primitive.ObjectID `json:"_id"`
	TourTitle       string             `json:"tour_title"`
	MeetingPoint    string             `json:"meeting_point"`
	StartTime       string             `json:"start_time"`
	LanguageOffered string             `json:"language_offered"`
	NumberOfTourist string             `json:"number_of_tourist"`
	Description     string             `json:"description"`
	TourGuide       string             `json:"tour_guide"`
	TourOperator    string             `json:"tour-operator"`
	OperatorContact string             `json:"operator_contact"`
	Date            string             `json:"date"`
}
