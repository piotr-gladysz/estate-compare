package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Separated package for smaller import

type Offer struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	SiteId  string             `json:"siteId" bson:"siteId"`
	Site    string             `json:"site" bson:"site"`
	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Created primitive.DateTime `json:"created" bson:"created"`

	Name           string   `json:"name" bson:"name"`
	Url            string   `json:"url" bson:"url"`
	Area           float32  `json:"area" bson:"area"`
	Rooms          int      `json:"rooms" bson:"rooms"`
	Floor          int      `json:"floor" bson:"floor"`
	BuildingFloors int      `json:"buildingFloors" bson:"buildingFloors"`
	Year           int      `json:"year" bson:"year"`
	Heating        string   `json:"heating" bson:"heating"`
	Market         string   `json:"market" bson:"market"`
	Window         string   `json:"window" bson:"window"`
	Elevator       bool     `json:"elevator" bson:"elevator"`
	Balcony        bool     `json:"balcony" bson:"balcony"`
	Media          []string `json:"media" bson:"media"`

	History []*OfferHistory `json:"history" bson:"history"`
}

type OfferHistory struct {
	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Price   int                `json:"price" bson:"price"`
}

type Condition struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Created primitive.DateTime `json:"created" bson:"created"`
	Name    string             `json:"name" bson:"name"`
	Wasm    []byte             `json:"wasm" bson:"wasm"`
}

type Notification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Created primitive.DateTime `json:"created" bson:"created"`

	ConditionId primitive.ObjectID `json:"conditionId" bson:"conditionId"`
	Config      map[string]any     `json:"config" bson:"config"`
}

type SentNotification struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Updated primitive.DateTime `json:"updated" bson:"updated"`
	Created primitive.DateTime `json:"created" bson:"created"`

	OfferId        primitive.ObjectID `json:"offerId" bson:"offerId"`
	NotificationId primitive.ObjectID `json:"notificationId" bson:"notificationId"`

	Message string `json:"message" bson:"message"`
}

type WatchUrl struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Created primitive.DateTime `json:"created" bson:"created"`
	Updated primitive.DateTime `json:"updated" bson:"updated"`

	Url      string `json:"url" bson:"url"`
	IsList   bool   `json:"isList" bson:"isList"`
	Disabled bool   `json:"disabled" bson:"disabled"`
}
