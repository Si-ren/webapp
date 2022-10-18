package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func restful(engine *gin.Engine) {
	//127.0.0.1:8180/student/siri/wangjaicun
	//siri live in /wangjaicun
	engine.GET("/student/:name/*addr", func(context *gin.Context) {
		name := context.Param("name")
		addr := context.Param("addr")
		context.String(http.StatusOK, name+" live in "+addr)
	})
}

func get(engine *gin.Engine) {
	engine.GET("/student", func(context *gin.Context) {
		name := context.Query("name")
		addr := context.DefaultQuery("addr", "Shanghai")
		context.String(http.StatusOK, name+" live in "+addr)

	})
}

type Student struct {
	Name       string    `form:"name" json:"name" uri:"name" xml:"name" yaml:"name" binding:"required" `
	Addr       string    `form:"addr" json:"addr" uri:"addr" xml:"addr" yaml:"addr" binding:"required" `
	Enrollment time.Time `form:"enrollment" binding:"before_today" time_format:"2006_01_02"`
	Graduation time.Time `form:"graduation" binding:"gtfield=Enrollment" time_format:"2006_01_02"`
}

func formBind(engine *gin.Engine) {
	engine.POST("/stu/form", func(context *gin.Context) {
		var stu Student
		if err := context.ShouldBind(&stu); err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			context.String(http.StatusOK, stu.Name+" live in "+stu.Addr)

		}
	})
}

// 绑定json
func jsonBind(engine *gin.Engine) {
	engine.POST("/stu/json", func(context *gin.Context) {
		var stu Student
		if err := context.ShouldBindJSON(&stu); err != nil {
			fmt.Println(err)
			context.String(http.StatusBadRequest, "parse paramter failed")
		} else {
			//返回字符串
			//context.String(http.StatusOK, stu.Name+" live in "+stu.Addr)
			//返回json，如果是返回别的类型，那么就是context.yaml等
			context.JSON(http.StatusOK, gin.H{"name": "lsl", "addr": "BeiHai"})
		}
	})
}

func main() {
	engine := gin.Default()
	//restful(engine)
	get(engine)
	formBind(engine)
	jsonBind(engine)
	engine.Run(":8180")
}
