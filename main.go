package main

import (
	"context"
	"fmt"
	"log"

	"net/http"

	"github.com/A-Victory/mongo-with-go/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.DELETE("/user", uc.DeleteUser)
	r.POST("/user/:id", uc.CreateUser)
	r.GET("/user", uc.GetAllUsers)

	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getSession() *mongo.Client {
	cltOptns := options.Client().ApplyURI("mongodb://localhost:27017")
	s, err := mongo.Connect(context.TODO(), cltOptns)

	if err != nil {
		panic(err)
	}
	if err := s.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return s
}
