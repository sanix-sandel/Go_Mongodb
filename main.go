package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

//the mongo-driver package uses another package
//called bson to serialize Go structs into BSON format.

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB successfully ")

	collection := client.Database("appDB").Collection("movies")

	//Create  a Movie
	darkNight := Movie{
		Name:      "The Dark Knight",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross:  533331538,
		},
	}

	//Insert a document into MongoDB
	_, err = collection.InsertOne(context.TODO(), darkNight)

	if err != nil {
		log.Fatal(err)
	}

	//Retrieve
	queryResult := &Movie{}
	//bson.M is used for building map for filter query
	filter := bson.M{"boxOffice.budget": bson.M{"$gt": 15000000}} //$gt => greater than
	result := collection.FindOne(context.TODO(), filter)
	err = result.Decode(queryResult)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie:", queryResult)
}
