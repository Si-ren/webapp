package controllers

import "fmt"

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
