package pkg

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoDb interface {
	Connect(uri string) (*mongo.Client, error)
	Ping(client mongo.Client) error
}
type m struct {
}

func NewMongo() IMongoDb {
	return &m{}
}

func (m m) Connect(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return client, err
	}

	return client, nil
}

// Ping implements IMongoDb.
func (*m) Ping(client mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	return err
}
