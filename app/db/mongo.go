package db

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context) *mongo.Database {
	prefix := viper.GetString("db.mongo.prefix")
	host := viper.GetString("db.mongo.host")
	port := viper.GetInt("db.mongo.port")
	username := viper.GetString("db.mongo.username")
	password := viper.GetString("db.mongo.password")
	suffix := viper.GetString("db.mongo.suffix")
	name := viper.GetString("db.mongo.name")

	connectionURI := fmt.Sprintf("%s://", prefix)

	if username != "" {
		connectionURI += username
	}

	if password != "" {
		connectionURI += fmt.Sprintf(":%s@", password)
	}

	if host != "" {
		connectionURI += host
	}

	if port != 0 {
		connectionURI += fmt.Sprintf(":%d", port)
	}

	connectionURI += fmt.Sprintf("/%s", name)

	if suffix != "" {
		connectionURI += fmt.Sprintf("?%s", suffix)
	}

	opt := options.Client().ApplyURI(connectionURI)

	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %s\n", err)
	}

	db := client.Database(name)

	return db
}
