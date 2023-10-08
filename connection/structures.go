package connection

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var MI MongoInstance

var Database struct {
	URI     string
	Name    string
	Timeout time.Duration
}

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
