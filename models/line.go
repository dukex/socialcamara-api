package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Line struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     []Link    `json:"links"`
	Status    Status    `json:"status"`
}

type Status struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Status    string        `json:"message"`
	LineId    string        `json:"line_id"  bson:"line_id"`
	Links     []Link        `json:"links"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
