package controllers

type ErrorController struct {
	Authentication
}

//Error开头的方法

// Error404 -> 处理404
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}
