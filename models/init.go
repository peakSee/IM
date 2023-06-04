package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

// InitMongo 初始化数据库
func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "mtjllb521",
	}).ApplyURI("mongodb://159.75.140.60:27017"))
	if err != nil {
		log.Println("Connection MongoDb Error:", err)
		return nil
	}
	return client.Database("im")
}
