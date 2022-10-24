package v1

import (
	"cmdb/base/auth"
	"cmdb/base/response"
	"cmdb/forms"
	"cmdb/services"
	"encoding/json"
	"fmt"
)

type PrometheusController struct {
	auth.APIController
}

func (c *PrometheusController) Register() {
	form := &forms.NodeRegisterForm{}
	fmt.Println(c.Ctx.Input.RequestBody)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, form); err != nil {
		fmt.Println("PrometheusController  Register :", form, err)
		c.Data["json"] = response.BadRequest
	} else {
		fmt.Println(form)
		services.NodeService.Register(form)
		c.Data["json"] = response.OK
	}

}

func (c *PrometheusController) Config() {
	uuid := c.GetString("uuid")
	rt := services.JobService.QueryByUUID(uuid)
	c.Data["json"] = response.NewJsonResponse(200, "OK", rt)
}
