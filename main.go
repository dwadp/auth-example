package main

import (
	"context"
	"log"

	"github.com/dwadp/auth-example/app/config"
	"github.com/dwadp/auth-example/app/db"
	"github.com/dwadp/auth-example/app/registry"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	ctx := context.Background()

	app := gin.Default()

	mongoDB := db.NewMongoDB(ctx)
	redisDB := db.NewRedis()

	if _, err := redisDB.Ping(ctx).Result(); err != nil {
		log.Fatalf("Error connecting to Redis Server: %s\n", err.Error())
	}

	//
	registry.New(app, mongoDB, redisDB, ctx)

	log.Fatal(app.Run(":5000"))
}
