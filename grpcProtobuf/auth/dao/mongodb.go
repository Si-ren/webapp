package dao

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"grpcProtobuf/auth/model"
)

type Mongodb struct {
	col *mongo.Collection
}

func NewMongodb(db *mongo.Database) *Mongodb {
	return &Mongodb{col: db.Collection("test")}
}

func (m *Mongodb) FindID(c context.Context, name string) (string, error) {
	res := m.col.FindOne(c, bson.M{
		"Name": name,
	}, &options.FindOneOptions{AllowPartialResults: nil})
	logrus.Info(res)
	//if err != nil {
	//	logrus.Error("Mongodb FindRows Error:", err)
	//	return "", fmt.Errorf("Mongodb FindName err: %v ", err)
	//}
	stu := &model.Student{}
	err := res.Decode(&stu)
	if err != nil {
		logrus.Error("Mongodb FindID Decode err :", err)
	}
	//logrus.Info(stu)
	return stu.ID.Hex(), nil
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
	stu := &model.Student{}
	err := res.Decode(&stu)
	if err != nil {
		return "", fmt.Errorf("Mongodb Decode Result err: %v ", err)
	}
	return stu.ID.Hex(), nil

}
