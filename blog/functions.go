package blog

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/connection"
)

func fetchBlogs(userID primitive.ObjectID) (blogs []Blog, err error) {
	filter := bson.D{
		// {
		// 	Key: "_id",
		// 	Value: bson.D{
		// 		{
		// 			Key:   "$in",
		// 			Value: ids,
		// 		},
		// 	},
		// },
		{
			Key:   "deleted.ok",
			Value: false,
		},
		{
			Key: "users",
			Value: bson.D{
				{
					Key: "$elemMatch",
					Value: bson.D{
						{
							Key:   "user_id",
							Value: userID,
						},
						{
							Key:   "deleted.ok",
							Value: false,
						},
					},
				},
			},
		},
	}
	sort := bson.D{
		{
			Key:   "created.time",
			Value: -1,
		},
	}
	opts := options.Find()
	opts.SetSort(sort)
	cursor, err := connection.MI.DB.Collection(BlogsDBName).Find(context.Background(), filter, opts)
	if err != nil {
		return
	}
	err = cursor.All(context.Background(), &blogs)
	if err != nil {
		return
	}
	return
}

func deleteBlog(userID primitive.ObjectID, blogID primitive.ObjectID) (err error) {
	t := time.Now().UTC()
	filter := bson.D{
		{
			Key:   "_id",
			Value: blogID,
		},
		{
			Key:   "created.user_id",
			Value: userID,
		},
	}
	deleted := Deleted{
		Ok:     true,
		UserID: userID,
		Time:   t,
	}
	updated := Updated{
		UserID: userID,
		Time:   t,
	}
	set := bson.D{
		{
			Key:   "deleted",
			Value: deleted,
		},
		{
			Key:   "updated",
			Value: updated,
		},
	}
	update := bson.D{
		{
			Key:   "$set",
			Value: set,
		},
	}
	_, err = connection.MI.DB.Collection(BlogsDBName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}
