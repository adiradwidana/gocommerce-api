package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Email     string        `bson:"email" json:"email"`
	Password  string        `bson:"password" json:"password"`
	Phone     string        `bson:"phone" json:"phone"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time     `bson:"deleted_at" json:"deleted_at"`
}

type Users []User
