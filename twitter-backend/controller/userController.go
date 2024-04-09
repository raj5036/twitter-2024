package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/twitter-2024/model"
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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	InsertOneUser(user)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {}

func InsertOneUser(user model.User) {
	inserted, err := users.InsertOne(context.Background(), user)
	handleErrors(err)

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func UpdateOneUser() {}

func GetOneUser() {}

func DeleteOneUser() {}
