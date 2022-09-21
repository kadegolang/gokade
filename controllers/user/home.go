package user

import (
	"gokade1/base/auth"
)

type HomeController struct {
	auth.AuthorizationController
}

// //html active 跳转
// func (c *HomeController) Prepare() {
// 	c.AuthorizationController.Prepare()
// 	c.Data["nav"] = "home"
// }

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
