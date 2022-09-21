package dao

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	col *mongo.Collection
}

func NewMongodb(db *mongo.Database) *Mongodb {
	return &Mongodb{col: db.Collection("test")}

}

func (m *Mongodb) FindID(c context.Context, name string) (string, error) {
	res, err := m.col.Find(c, bson.M{
		"Name": name,
	})
	if err != nil {
		logrus.Error("Mongodb FindRows Error:", err)
		return "", fmt.Errorf("Mongodb FindName err: %v ", err)
	}
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		Name   string             `bson:"Name"`
		Age    string             `bson:"Age"`
		Salary string             `bson:"Salary"`
		Sex    string             `bson:"Sex"`
	}
	err = res.Decode(&row)
	logrus.Info(row)
	return row.Name, nil
}

func (m *Mongodb) FindAndUpdate(c context.Context, oldNum string, newNum string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		"Salary": oldNum,
	}, bson.M{
		"$set": bson.M{
			"Salary": newNum,
		},
	}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("Mongodb FindOneAndUpdate err: %v ", err)
	}
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		Name   string             `bson:"Name"`
		Age    string             `bson:"Age"`
		Salary string             `bson:"Salary"`
		Sex    string             `bson:"Sex"`
	}
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("Mongodb Decode Result err: %v ", err)
	}
	return row.ID.Hex(), nil

}
