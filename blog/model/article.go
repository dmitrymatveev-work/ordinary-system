package model

import (
	"gopkg.in/mgo.v2/bson"
)

// Article is a blog article
type Article struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
}
