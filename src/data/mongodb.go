package data

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error occurred while loading .env file")
	}

	connectionUrl := os.Getenv("MONGODB_CONNECTION_URL")

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionUrl).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal("error occurred while connecting to database")
	}

	return client
}

var mongoClient *mongo.Client = DbInstance()

func openCollection(dbClient *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = dbClient.Database("SmartFmDb").Collection(collectionName)
	return collection
}

var UsersCollection *mongo.Collection = openCollection(mongoClient, "users")
var OtpsCollection *mongo.Collection = openCollection(mongoClient, "otps")
