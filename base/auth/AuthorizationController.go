package auth

import (
	"strings"

	"gokade1/models"
	"gokade1/services"

	"github.com/astaxie/beego"
)

// AuthorizationController 所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	BaseController
	LoginUser *models.User
}

//取控制器前缀，为active做准备
func (c *AuthorizationController) getNav() string {
	controllername, _ := c.GetControllerAndAction()
	return strings.ToLower(strings.TrimSuffix(controllername, "Controller"))
}

// Prepare 用户认证检查
func (c *AuthorizationController) Prepare() {
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user") //逗号后面是默认值
	sessionduanyan := c.GetSession(sessionKey)

	c.Data["LoginUser"] = nil //初始化一下值
	c.Data["nav"] = c.getNav()

	if sessionduanyan != nil {
		if k, ok := sessionduanyan.(int); ok {
			if user := services.UserService.GetByID(k); user != nil {
				c.Data["LoginUser"] = user
				c.LoginUser = user
				// fmt.Println("cccc:", c.LoginUser)
				return
			}
		}
	}
	// action := beego.AppConfig.DefaultString("auth::LoginAction",
	// 	"AuthController.Login") //逗号后面是默认值
	c.Redirect(beego.URLFor("AuthController.Login"), 302)
}

// sessionUser := c.GetSession("user")
// if sessionUser == nil {
// 	//无session信息（未登陆）
// 	//session断言=》int
// 	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
// 	return
// }
