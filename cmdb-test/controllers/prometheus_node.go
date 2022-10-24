package controllers

import (
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type NodeController struct {
	prometheusController
}

func (c *NodeController) Prepare() {
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

func (c *NodeController) Delete() {
	if id, err := c.GetInt("id"); err == nil {
		services.NodeService.DeleteByID(id)
	}
	c.Redirect(beego.URLFor("NodeController.Query"), http.StatusFound)

}
