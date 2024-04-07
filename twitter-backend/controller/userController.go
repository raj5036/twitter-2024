package userController

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_STRING string = "mongodb://localhost:27017"
const DATABASE string = "twitter-project-go"
const COLLECTION string = "users"

// MOST IMPORTANT
var users *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING)
	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	handleErrors(err)
	fmt.Println("MongoDB connection success")

	// Most Important: Assign collection to this Collection Instance
	users = client.Database(DATABASE).Collection(COLLECTION)

	//collection instance
	fmt.Println("Collection instance is ready")
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func insertOneUser() {}

func updateOneUser() {}

func getOneUser() {}

func deleteOneUser() {}
