package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/raj5036/twitter-2024/api"
	"github.com/raj5036/twitter-2024/model"
	"github.com/raj5036/twitter-2024/response"
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
var secretKey = []byte("secret-key")

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
			api.ResponseError(w, "User with same email already exists", http.StatusConflict)
			return
		}
	}

	if user.PhoneNumber != "" { // If user registers using Phone Number
		phoneFilter := bson.M{"phoneNumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name != "" {
			fmt.Println("User with same phone number already exists")
			api.ResponseError(w, "User with same phoneNumber already exists", http.StatusConflict)
			return
		}
	}

	hashedPassword, err := hashPassword(user.Password)
	handleErrors(err)
	user.Password = hashedPassword

	insertOneUser(user)
	api.ResponseOK(w, user, http.StatusCreated)
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
			api.ResponseError(w, "No such Email exists in system", http.StatusBadRequest)
			return
		}
		passwordMatch := comparePassword(userWithEmail.Password, user.Password)
		if !passwordMatch {
			fmt.Println("Incorrect password")
			api.ResponseError(w, "Incorrect email / password combination", http.StatusBadRequest)
			return
		}

		tokenString := createUserToken(userWithEmail)
		data := response.LoginResponse{Token: tokenString}

		api.ResponseOK(w, data, http.StatusOK)
	} else if user.PhoneNumber != "" {
		phoneFilter := bson.M{"phoneNumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name == "" {
			fmt.Println("No such PhoneNumber exists in system")
			api.ResponseError(w, "No such PhoneNumber exists in system", http.StatusBadRequest)
			return
		}

		passwordMatch := comparePassword(userWithPhone.Password, user.Password)
		if !passwordMatch {
			fmt.Println("Incorrect password")
			api.ResponseError(w, "Incorrect phoneNumber / password combination", http.StatusBadRequest)
			return
		}

		tokenString := createUserToken(userWithPhone)
		data := response.LoginResponse{Token: tokenString}

		api.ResponseOK(w, data, http.StatusOK)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email != "" {
		emailFilter := bson.M{"email": user.Email}
		userWithEmail := getOneUser(emailFilter)

		if userWithEmail.Name == "" {
			fmt.Println("No such Email exists.")
			api.ResponseError(w, "No such Email exists in system", http.StatusBadRequest)
			return
		}

		api.ResponseOK(w, userWithEmail, http.StatusOK)
	} else if user.PhoneNumber != "" {
		phoneFilter := bson.M{"phoneNumber": user.PhoneNumber}
		userWithPhone := getOneUser(phoneFilter)

		if userWithPhone.Name == "" {
			fmt.Println("No such phoneNumber exists.")
			api.ResponseError(w, "No such Phone Number exists", http.StatusBadRequest)
			return
		}

		api.ResponseOK(w, userWithPhone, http.StatusOK)
	} else {
		api.ResponseError(w, "Request must contain Email or PhoneNumber", http.StatusBadRequest)
	}

}

func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bulkDeleteFilter := bson.M{}
	deleteResult := deleteManyUser(bulkDeleteFilter)
	fmt.Println("Deleted all Users, count = ", deleteResult.DeletedCount)

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(deleteResult)
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

func insertOneUser(user model.User) {
	inserted, err := Users.InsertOne(context.Background(), user)
	handleErrors(err)

	fmt.Println("Inserted 1 user in db with id: ", inserted.InsertedID)
}

// func updateOneUser() {}

// func deleteOneUser() {}

func deleteManyUser(filter primitive.M) *mongo.DeleteResult {
	deleteResult, err := Users.DeleteMany(context.Background(), filter)
	handleErrors(err)
	return deleteResult
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
	return err == nil
}

func createUserToken(user model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	handleErrors(err)

	return tokenString
}
