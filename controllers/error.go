package controllers

import "gokade1/base/auth"

//错误处理控制器
type ErrorController struct {
	auth.BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}


func (c *ErrorController) ErrorNoPermissions() {
	c.TplName = "error/NoPermissions.html"
}


