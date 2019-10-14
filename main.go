package main

import (
	"fmt"
	"log"
	"net/http"
	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type Trainer struct {
	Name string
	Age int
	City string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWorld)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")

	collection : client.Database("test").Collection("trainers")


	log.Fatal(http.ListenAndServe(":8080", router))
}

