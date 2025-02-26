package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Url struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	LongUrl    *string            `json:"longUrl" bson:"longUrl" validate:"required"`
	ShortUrl   *string            `json:"shortUrl" bson:"shortUrl" validate:"required"`
	VisitCount int                `json:"visitCount" bson:"visitCount"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
}

func NewUrl(longUrl *string, shortUrl *string) Url {
	return Url{
		ID:         primitive.NewObjectID(),
		LongUrl:    longUrl,
		ShortUrl:   shortUrl,
		VisitCount: 0,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
}
