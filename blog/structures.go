package blog

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var BlogsDBName string = "blogs"

type Blog struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type     string             `json:"type" bson:"type"`
	Body     string             `json:"body" bson:"body"`
	Users    []User             `json:"users" bson:"users"`
	Created  Created            `json:"created" bson:"created"`
	Updated  Updated            `json:"updated" bson:"updated"`
	Deleted  Deleted            `json:"deleted" bson:"deleted"`
	ParentID primitive.ObjectID `json:"parent_id" bson:"parent_id"`
	Likes    []Like             `json:"likes" bson:"likes"`
	Comments []Comment          `json:"comments" bson:"comments"`
	// OriginID primitive.ObjectID `json:"origin_id" bson:"origin_id"`
}

type Created struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Time   time.Time          `json:"time" bson:"time"`
}

type Updated struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Time   time.Time          `json:"time" bson:"time"`
}
type Deleted struct {
	Ok     bool               `json:"ok" bson:"ok"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Time   time.Time          `json:"time" bson:"time"`
}

type Like struct {
	Ok     bool               `json:"ok" bson:"ok"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Time   time.Time          `json:"time" bson:"time"`
}

type Comment struct {
	Ok     bool               `json:"ok" bson:"ok"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Time   time.Time          `json:"time" bson:"time"`
}

type User struct {
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	FullName  string             `json:"full_name" bson:"full_name"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Deleted   Deleted            `json:"deleted" bson:"deleted"`
}
