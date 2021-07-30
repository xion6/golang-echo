package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xion6/golang-echo/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewMongoConnection(cfg *config.Settings) Connection {

	uri := fmt.Sprintf("mongodb://%s/%s", cfg.DbHost, cfg.DbName)

	credentials := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPass,
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to database.")

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

func (c Connection) Disconnect() {
	c.Client.Disconnect(c.ctx)
}
