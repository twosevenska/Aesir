package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type BaseEntity struct {
	ID        bson.Binary `bson:"_id,omitempty" json:"id"`
	Name      string      `bson:"name" json:"name"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at" json:"updated_at"`
}
