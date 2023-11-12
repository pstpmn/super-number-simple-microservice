package pkg

import (
	"context"
	"log"
	"super-number-simple-microservice/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	IMongo interface {
		Conn(pctx context.Context, cfg *configs.Mongo) *mongo.Client
	}
	m struct{}
)

// Conn implements IMongo.
func (m) Conn(pctx context.Context, cfg *configs.Mongo) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url))
	if err != nil {
		log.Fatalf("Error: Conntect to database error: %s", err.Error())
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error: Pinging to database error: %s", err.Error())
	}
	return client
}

func NewMongo() IMongo {
	return m{}
}
