package controllers

import (
	"cmdb/forms"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type TargetController struct {
	prometheusController
}

func (c *TargetController) Prepare() {
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

func (c *TargetController) Delete() {
	if id, err := c.GetInt("id"); err == nil {
		services.TargetService.DeleteByID(id)
	}
	c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)

}

func (c *TargetController) Create() {
	form := &forms.TargetCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.TargetService.Create(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Layout = "layout/layout.html"
	c.Data["nodes"], _ = services.NodeService.Query("")
	c.TplName = "prometheus/target/create.html"
}

func (c *TargetController) Modify() {
	//Get显示查询
	//Post更新
	form := &forms.TargetModifyForm{}
	if c.Ctx.Input.IsGet() {
		if id, err := c.GetInt("id"); err == nil {
			target := services.TargetService.GetByID(id)
			form.ID = target.ID
			form.Name = target.Name
			form.Remark = target.Remark
			form.Job = target.Job.ID
		}

	}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.TargetService.Modify(form)
			c.Redirect(beego.URLFor("TargetController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["jobs"], _ = services.JobService.Query("")
	c.Layout = "layout/layout.html"
	c.TplName = "prometheus/target/modify.html"
}
