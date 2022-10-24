package conf

import (
	"fmt"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {

	err := LoadConfigFromToml("../etc/demo.toml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(global.Log)
	if global.Log.PathDir != "logs" {
		t.Fatalf("Load Config From Toml Error")
	}
}
