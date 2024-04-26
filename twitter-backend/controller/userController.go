package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/twitter-2024/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Check if same Email or PhoneNumber already exists
	fmt.Printf("user through request %s\n", user)
	if user.Email != "" {
		emailFilter := bson.M{"email": user.Email}
		userWithEmail := getOneUser(emailFilter)

		if userWithEmail.Name != "" {
			fmt.Println("Users with same email already exists")
			return
		}
	}

	if user.PhoneNumber != "" {
		phoneFilter := bson.M{"phonenumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name != "" {
			fmt.Println("Users with same phone number already exists")
			return
		}
	}

	InsertOneUser(user)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	w.WriteHeader(200)

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
}

func getOneUser(filter primitive.M) model.User {
	var user model.User
	err := users.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No documents found with given result")
		return user
	} else if err != nil {
		handleErrors(err)
	}

	return user
}

func InsertOneUser(user model.User) {
	inserted, err := users.InsertOne(context.Background(), user)
	handleErrors(err)

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func UpdateOneUser() {}

func DeleteOneUser() {}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
