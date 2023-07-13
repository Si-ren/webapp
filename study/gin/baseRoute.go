package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法1前")
		c.Next()
		fmt.Println("方法1后")

	}
}

func middleTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("方法2前")
		c.Next()
		fmt.Println("方法2后")

	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1").Use(middle(), middleTwo())
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("v1方法")
		c.JSON(200, gin.H{
			"sucess": true,
		})
	})
	r.Run(":8080")
}
