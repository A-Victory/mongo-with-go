package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/A-Victory/mongo-with-go/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	check(err)

	filter := bson.M{"_id": objectID}

	u := models.User{}
	db := uc.Client.Database("go-web-dev").Collection("users")
	if err = db.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		panic(err)
	}

	uj, err := json.Marshal(u)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := []models.User{}

	db := uc.Client.Database("go-web-dev").Collection("users")
	docs, err := db.Find(context.TODO(), bson.M{})
	check(err)

	if err = docs.All(context.TODO(), &u); err != nil {
		panic(err)
	}

	uj, err := json.Marshal(u)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID()

	insert, err := uc.Client.Database("go-web-dev").Collection("users").InsertOne(context.TODO(), u)
	check(err)

	uj, err := json.Marshal(u)
	check(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Println(insert.InsertedID)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Writes code that deletes user
	/*id := p.ByName("id")

	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	*/

	//objectID, err := primitive.ObjectIDFromHex(id)
	//check(err)

	//filter := bson.M{"_id": objectID}

	db := uc.Client.Database("go-web-dev").Collection("users")
	delete, err := db.DeleteMany(context.TODO(), bson.D{})
	check(err)

	w.WriteHeader(200)
	fmt.Println(delete)
	fmt.Fprint(w, "User deleted successfully")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
