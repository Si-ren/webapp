package connectivity_test

import (
	"cmdb/provider/aliyun/connectivity"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	var secretID, secretKey string
	if secretID = os.Getenv("ALLOUD_ACCESS_KEY"); secretID == "" {
		t.Fatal("empty ALLOUD_ACCESS_KEY")
	}

	if secretKey = os.Getenv("ALCLOUD_ACCESS_SECRET"); secretKey == "" {
		t.Fatal("empty ALCLOUD_ACCESS_SECRET")
	}

	client := connectivity.NewAliCloudClient(secretID, secretKey, "cn-hangzhou")
	client.EcsClient()
}
