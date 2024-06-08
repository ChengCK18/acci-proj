package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongoDB() {
	var err error
	if err := godotenv.Load("../../.env"); err != nil { //relative to main.go loc
		log.Println("No .env file found~")
	}

	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		log.Fatalf("MONGODB_URI is not defined, cant connect :<")
	}

	// init mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Initiated!")

}

func DisconnectMongoClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
