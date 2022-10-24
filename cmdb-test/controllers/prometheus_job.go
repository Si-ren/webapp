package controllers

import (
	"cmdb/forms"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
)

type JobController struct {
	prometheusController
}

func (c *JobController) Prepare() {
	c.prometheusController.Prepare()
}

func (c *JobController) Query() {
	//flash读取 flash存储的值, 读取完会把c.data里面的值删掉
	fmt.Println("JobController Query")
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

func (c *JobController) Delete() {
	if id, err := c.GetInt("id"); err == nil {
		services.JobService.DeleteByID(id)
	}
	c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
}

func (c *JobController) Create() {
	form := &forms.JobCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.JobService.Create(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Layout = "layout/layout.html"
	c.Data["nodes"], _ = services.NodeService.Query("")
	c.TplName = "prometheus/job/create.html"
}

func (c *JobController) Modify() {
	//Get显示查询
	//Post更新
	form := &forms.JobModifyForm{}
	if c.Ctx.Input.IsGet() {
		if id, err := c.GetInt("id"); err == nil {
			job := services.JobService.GetByID(id)
			form.ID = job.ID
			form.Key = job.Key
			form.Remark = job.Remark
			form.Node = job.Node.ID
		}

	}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.JobService.Modify(form)
			c.Redirect(beego.URLFor("JobController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["nodes"], _ = services.NodeService.Query("")
	c.Layout = "layout/layout.html"
	c.TplName = "prometheus/job/modify.html"
}
