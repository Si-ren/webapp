package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

// go test -timeout 30s grpcProtobuf/auth/dao -run ^TestResolveID$
func TestResolveID(t *testing.T) {
	c := context.Background()
	mgc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://122.51.16.26:27017"))
	if err != nil {
		t.Fatalf("cant connect mongodb: %v", err)
	}
	mongodb := NewMongodb(mgc.Database("grpcProtobuf"))
	id, err := mongodb.FindAndUpdate(c, "20000", "25000")
	if err != nil {
		t.Error("mongodb find and update err")

	} else {
		if id != "63289fff6053de32ba0b3a22" {
			t.Error("mongodb find and update err")
		}
	}
}
