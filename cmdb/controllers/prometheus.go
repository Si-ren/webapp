package controllers

import (
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
)

type prometheusController struct {
	Authentication
}

func (c *prometheusController) PrePare() {
	c.Authentication.Prepare()
	c.Data["nav"] = "prometheus"
	controller, action := c.GetControllerAndAction()
	fmt.Println("prometheusController  PrePare", controller, action)
	c.Data["subnav"] = controller
}

type NodeController struct {
	prometheusController
}

func (c *NodeController) PrePare() {
	c.prometheusController.Prepare()
}

func (c *NodeController) Query() {
	//flash读取 flash存储的值, 读取完会把c.data里面的值删掉
	fmt.Println("NodeController Query")
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)

	query := c.Ctx.Input.Query("query")
	nodes, err := services.NodeService.Query(query)
	if err != nil {
		fmt.Println(nodes, err)
	}
	c.Data["nodes"] = nodes
	c.TplName = "prometheus/node/query.html"
}

type TargetController struct {
	prometheusController
}

func (c *TargetController) PrePare() {
	c.prometheusController.Prepare()
}

func (c *TargetController) Query() {
	//flash读取 flash存储的值, 读取完会把c.data里面的值删掉
	fmt.Println("NodeController Query")
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)

	query := c.Ctx.Input.Query("query")
	targets, err := services.TargetService.Query(query)
	if err != nil {
		fmt.Println(targets, err)
	}
	c.Data["targets"] = targets
	c.TplName = "prometheus/target/query.html"
}

type JobController struct {
	prometheusController
}

func (c *JobController) PrePare() {
	c.prometheusController.Prepare()
}
func (c *JobController) Query() {
	//flash读取 flash存储的值, 读取完会把c.data里面的值删掉
	fmt.Println("NodeController Query")
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)

	query := c.Ctx.Input.Query("query")
	jobs, err := services.JobService.Query(query)
	if err != nil {
		fmt.Println(jobs, err)
	}
	c.Data["jobs"] = jobs
	c.TplName = "prometheus/job/query.html"
}
