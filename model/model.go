package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tourist struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     string             `bson:"first_name" Usage:"required,alpha,omitempty"`
	LastName      string             `bson:"last_name" Usage:"required,alpha,omitempty"`
	Email         string             `bson:"email" Usage:"required,alphanumeric"`
	Password      string             `bson:"password" Usage:"required"`
	CheckPassword string             `bson:"check_password" Usage:"required"`
	Phone         string             `bson:"phone" Usage:"required"`
	BookedTours   []Tour             `bson:"booked_tour"`
	RequestTours  []Tour             `bson:"request_tour"`
	TaggedTourist []TaggedTourist    `bson:"tagged_tourist"`
	GeoLocation   string             `bson:"geo_location"`
	Token         string             `bson:"token" Usage:"jwt"`
	NewToken      string             `bson:"new_token" Usage:"jwt"`
	CreatedAt     time.Time          `bson:"created_at" Usage:"datetime"`
	UpdatedAt     time.Time          `bson:"updated_at" Usage:"datetime"`
}

type TaggedTourist struct {
	FirstName string `bson:"first_name" Usage:"required,alpha,omitempty"`
	LastName  string `bson:"last_name" Usage:"required,alpha,omitempty"`
	Email     string `bson:"email" Usage:"required,alphanumeric"`
	Password  string `bson:"password" Usage:"required"`
}

type Operator struct {
	ID            primitive.ObjectID `bson:"_id"`
	FirstName     string             `bson:"first_name" Usage:"required,alpha,omitempty"`
	LastName      string             `bson:"last_name" Usage:"required,alpha,omitempty"`
	Email         string             `bson:"email" Usage:"required,alphanumeric"`
	Password      string             `bson:"password" Usage:"required"`
	CheckPassword string             `bson:"check_password" Usage:"required"`
	Phone         string             `bson:"phone" Usage:"required"`
	ToursList     []Tour             `bson:"tours_list"`
	GeoLocation   string             `bson:"geo_location"`
	Token         string             `bson:"token" Usage:"jwt"`
	NewToken      string             `bson:"new_token" Usage:"jwt"`
	CreatedAt     time.Time          `bson:"created_at" Usage:"datetime"`
	UpdatedAt     time.Time          `bson:"updated_at" Usage:"datetime"`
}

type Tour struct {
	ID              primitive.ObjectID `bson:"_id"`
	OperatorID      string             `bson:"operator_id"`
	TourTitle       string             `bson:"tour_title"`
	MeetingPoint    string             `bson:"meeting_point"`
	StartTime       string             `bson:"start_time"`
	LanguageOffered string             `bson:"language_offered"`
	NumberOfTourist string             `bson:"number_of_tourist"`
	Description     string             `bson:"description"`
	TourGuide       string             `bson:"tour_guide"`
	TourOperator    string             `bson:"tour_operator"`
	OperatorContact string             `bson:"operator_contact"`
	Category        string             `bson:"category"`
	Activity        string             `bson:"activity"`
	Date            string             `bson:"date"`
	Price           int64              `bson:"price"`
}

type UserInfo struct {
	ID       primitive.ObjectID
	Email    string
	Password string
}
