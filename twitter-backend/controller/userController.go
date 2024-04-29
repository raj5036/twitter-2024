package userController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raj5036/twitter-2024/api"
	"github.com/raj5036/twitter-2024/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

const CONNECTION_STRING string = "mongodb://localhost:27017"
const DATABASE string = "twitter-project-go"
const COLLECTION string = "users"

// MOST IMPORTANT
var Users *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(CONNECTION_STRING)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	handleErrors(err)
	fmt.Println("MongoDB connection success")

	// Most Important: Assign collection to this Collection Instance
	Users = client.Database(DATABASE).Collection(COLLECTION)

	//collection instance
	fmt.Println("Collection instance is ready")
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Check if same Email or PhoneNumber already exists
	fmt.Printf("user through request %s\n", user)
	if user.Email != "" { // If user logs in using Email
		emailFilter := bson.M{"email": user.Email}
		userWithEmail := getOneUser(emailFilter)

		if userWithEmail.Name != "" {
			fmt.Println("User with same email already exists")
			api.ResponseError(w, "User with same email already exists", 409)
			return
		}
	}

	if user.PhoneNumber != "" { // If user logs in using Phone Number
		phoneFilter := bson.M{"phoneNumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name != "" {
			fmt.Println("User with same phone number already exists")
			api.ResponseError(w, "User with same phoneNumber already exists", 409)
			return
		}
	}

	hashedPassword, err := hashPassword(user.Password)
	handleErrors(err)
	user.Password = hashedPassword

	InsertOneUser(user)
	api.ResponseOK(w, user, 200)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email != "" {
		emailFilter := bson.M{"email": user.Email}
		userWithEmail := getOneUser(emailFilter)

		if userWithEmail.Name == "" {
			fmt.Println("No such Email exists in system")
			api.ResponseError(w, "No such Email exists in system", 404)
			return
		}
		passwordMatch := comparePassword(userWithEmail.Password, user.Password)
		if !passwordMatch {
			fmt.Println("Incorrect password")
			api.ResponseError(w, "Incorrect email / password combination", 400)
			return
		}

		api.ResponseOK(w, user, 200)
	} else if user.PhoneNumber != "" {
		phoneFilter := bson.M{"phoneNumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name == "" {
			fmt.Println("No such PhoneNumber exists in system")
			api.ResponseError(w, "No such PhoneNumber exists in system", 404)
			return
		}

		passwordMatch := comparePassword(userWithPhone.Password, user.Password)
		if !passwordMatch {
			fmt.Println("Incorrect password")
			api.ResponseError(w, "Incorrect phoneNumber / password combination", 400)
			return
		}

		api.ResponseOK(w, user, 200)
	}
}

func getOneUser(filter primitive.M) model.User {
	var user model.User
	err := Users.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No documents found with given result")
		return user
	} else if err != nil {
		handleErrors(err)
	}

	return user
}

func InsertOneUser(user model.User) {
	inserted, err := Users.InsertOne(context.Background(), user)
	handleErrors(err)

	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

func UpdateOneUser() {}

func DeleteOneUser() {}

func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deleteResult, err := Users.DeleteMany(context.Background(), bson.M{})
	handleErrors(err)
	fmt.Println("Deleted all Users, count = ", deleteResult.DeletedCount)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(deleteResult)
}

func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func comparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	fmt.Println("err", err)
	return err == nil
}
