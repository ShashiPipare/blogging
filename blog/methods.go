package blog

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/connection"
)

func (blog *Blog) create(userID primitive.ObjectID, t time.Time) (err error) {
	created := Created{
		UserID: userID,
		Time:   t,
	}
	blog.Created = created
	ID, err := connection.MI.DB.Collection(BlogsDBName).InsertOne(context.TODO(), blog)
	if err != nil {
		return
	}
	blog.ID, _ = ID.InsertedID.(primitive.ObjectID)
	return
}

func (blog *Blog) update(userID primitive.ObjectID, blogID primitive.ObjectID, body string, newUsers []User, t time.Time) (err error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: blogID,
		},
	}
	updated := Updated{
		UserID: userID,
		Time:   t,
	}
	set := bson.D{
		{
			Key:   "body",
			Value: body,
		},
		{
			Key:   "updated",
			Value: updated,
		},
	}
	push := bson.D{
		{
			Key: "users",
			Value: bson.D{
				{
					Key:   "$each",
					Value: newUsers,
				},
			},
		},
	}
	update := bson.D{
		{
			Key:   "$set",
			Value: set,
		},
	}
	if len(newUsers) > 0 {
		element := bson.E{
			Key:   "$push",
			Value: push,
		}
		update = append(update, element)
	}
	after := options.After
	opts := options.FindOneAndUpdate()
	opts.ReturnDocument = &after
	err = connection.MI.DB.Collection(BlogsDBName).FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&blog)
	if err != nil {
		return
	}
	return
}

func (blog *Blog) fetchBlog(userID primitive.ObjectID, blogID primitive.ObjectID) (err error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: blogID,
		},
		{
			Key:   "created.user_id",
			Value: userID,
		},
		{
			Key:   "deleted.ok",
			Value: false,
		},
	}
	err = connection.MI.DB.Collection(BlogsDBName).FindOne(context.TODO(), filter).Decode(blog)
	if err != nil {
		return
	}
	return
}
