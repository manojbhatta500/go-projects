package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/manojbhatta500/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connectionstring = "mongodb://localhost:27017"

var DbName = "netflix"

var CollName = "watchlist"

var Collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(Connectionstring)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database(DbName).Collection(CollName)
	fmt.Println("collection instance is ready ")
}

// mongo helper

func insertOneMovie(film model.Netflix) {
	id, err := Collection.InsertOne(context.TODO(), film)
	if err != nil {
		fmt.Println("we got err and err is ", err.Error())
	}
	fmt.Println("data successfully inserted  id : ", id.InsertedID)
}


