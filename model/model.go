package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tourist struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName     string             `json:"first_name" Usage:"required,alpha" json:"name,omitempty"`
	LastName      string             `json:"last_name" Usage:"required,alpha" json:"name,omitempty"`
	Email         string             `json:"email" Usage:"required,alphanumeric"`
	Password      string             `json:"password" Usage:"required"`
	CheckPassword string             `json:"check_password" Usage:"required"`
	Phone         string             `json:"phone" Usage:"required"`
	ReservedTours []Tour             `json:"tours_list"`
	GeoLocation   string             `json:"geo_location"`
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
	TourOperator    string             `json:"tour_operator"`
	OperatorContact string             `json:"operator_contact"`
	Date            string             `json:"date"`
}

type IntraData struct {
	ID       primitive.ObjectID
	Email    string
	Password string
}
