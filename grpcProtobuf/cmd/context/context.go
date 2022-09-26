package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	c, cancel := context.WithTimeout(c, 14*time.Second)
	defer cancel()
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Println("Main task started with param ：", c.Value(paramKey{}))
	go smallTask(context.Background(), "Task1")
	go smallTask(c, "Task2")
}

func smallTask(c context.Context, name string) {
	fmt.Println("Start task :", name, c.Value(paramKey{}))
	select {
	case <-time.After(14 * time.Second):
		fmt.Println(name, " Time after Done")
	case <-c.Done():
		fmt.Println(name, "Context Done")
	}
}
