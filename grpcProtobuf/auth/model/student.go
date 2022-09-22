package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"Name"`
	Age    string             `bson:"Age"`
	Salary string             `bson:"Salary"`
	Sex    string             `bson:"Sex"`
}
