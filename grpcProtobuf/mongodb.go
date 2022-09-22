package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {
	c := context.Background()

	mgc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://122.51.16.26:27017"))
	if err != nil {
		fmt.Println("Init Mongodb Client Err: ", err)
		os.Exit(1)
	}
	col := mgc.Database("grpcProtobuf").Collection("test")
	res, err := col.InsertOne(c, bson.M{
		"Name":   "qwe",
		"Age":    "11",
		"Salary": "1234",
		"Sex":    "0",
	})
	fmt.Println("Result is :", res.InsertedID)

	findRows(c, col)
}

func findRows(c context.Context, col *mongo.Collection) {
	res, err := col.Find(c, bson.M{
		"Name": "siri",
	})
	if err != nil {
		fmt.Println("Mongodb Find Err: ", err)
		os.Exit(1)
	}
	for res.Next(c) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			Name   string             `bson:"Name"`
			Age    string             `bson:"Age"`
			Salary string             `bson:"Salary"`
			Sex    string             `bson:"Sex"`
		}
		err = res.Decode(&row)
		if err != nil {
			fmt.Println("Decode Err: ", err)
			os.Exit(1)
		}
		logrus.Info(row)
	}
	//fmt.Printf("%+v\n", res)

}
