package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"grpcProtobuf/common/id"
)

type Mongodb struct {
	col *mongo.Collection
}

type BlobRecord struct {
	IDField   primitive.ObjectID `bson:"_id"`
	AccountID string             `bson:"accountid"`
	Path      string             `bson:"path"`
}

func NewMongodb(db *mongo.Database) *Mongodb {
	return &Mongodb{col: db.Collection("test")}
}

func (m *Mongodb) CreateBolb(c context.Context, aid id.AccountID) (*BlobRecord, error) {
	br := &BlobRecord{AccountID: aid.String()}
	br.IDField = primitive.NewObjectID()
	br.Path = fmt.Sprintf("%s/%s", aid.String(), br.IDField.Hex())
	_, err := m.col.InsertOne(c, br)
	if err != nil {
		return nil, err
	}
	return br, nil

}

func (m *Mongodb) GetBolbUrl(c context.Context, aid id.AccountID) (string, error) {
	str, err := m.col.FindOne(c)
}
