package user

import (
	"fmt"
	"gokade1/base/auth"
	"gokade1/base/errors"
	"gokade1/forms"
	"gokade1/models"
	"gokade1/services"
	"net/http"

	"github.com/astaxie/beego"
)

var Loginuser *models.User //cc 的值就是LoginUser

// AuthController 认证控制器
type AuthController struct {
	auth.BaseController
}

// Login 认证登录
func (c *AuthController) Login() {
	Sessionkey := beego.AppConfig.DefaultString("auth::Sessionkey", "user")
	sessionUser := c.GetSession(Sessionkey)
	if sessionUser != nil {
		HomeAction := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
		c.Redirect(beego.URLFor(HomeAction), http.StatusFound)
		return
	}

	form := &forms.LoginForm{}
	errs := errors.New() //定义错误,后面引用错误
	// errs := errors.New()
	// Get请求直接加载页面
	// Post请求进行数据验证
	if c.Ctx.Input.IsPost() {
		// 获取用户提交数据
		if err := c.ParseForm(form); err == nil {
			user := services.UserService.GetByName(form.Name)
			// fmt.Println(user)
			if user == nil {
				errs.Add("default", "用户名或密码错误")
				beego.Error(fmt.Sprintf("用户认证失败：%s", form.Name)) //打印认证失败用户名
				// 用户不存在
			} else if user.ValidPassword(form.Password) {
				beego.Informational(fmt.Sprintf("用户认证成功：%s", form.Name)) //打印认证成功用户名
				// 用户密码正确
				// 记录用户状态(session 记录服务器端)

				Loginuser = services.UserService.GetByName(string(form.Name)) //为后面的权限做准备，LoginUser只能在userconyroller里面用，这个是全局可以用的
				// fmt.Println("CCCCCCCCC:", Loginuser)

				Sessionkey := beego.AppConfig.DefaultString("auth::Sessionkey", "user")
				c.SetSession(Sessionkey, user.ID)
				// HomeAction := beego.AppConfig.DefaultString("auth::HomeAction", "HomeController.Index")
				// c.Redirect("/user/query", http.StatusFound) //认证成功302跳转
				c.Redirect(beego.URLFor("HomeController.Index"), http.StatusFound)
			} else {
				// 用户密码不正确
				errs.Add("default", "用户名或密码错误")
				beego.Error(fmt.Sprintf("用户密码认证失败：%s", form.Name)) //打印认证失败用户名
			}
		} else {
			fmt.Println(err)
			errs.Add("default", "用户名或密码错误")
		}
	}

	c.Data["form"] = form //返回值
	c.Data["errors"] = errs
	// 定义加载页面
	// c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken() //生成一个xsrf随机token防着sxrf攻击 html 需要写一个input标签
	c.TplName = "auth/login.html"
}

// Login 退出登录
func (c *AuthController) Loginout() {
	c.DestroySession() //销毁sessionid
	LoginoutAction := beego.AppConfig.DefaultString("auth::LoginoutAction", "AuthController.Login")
	c.Redirect(beego.URLFor(LoginoutAction), http.StatusFound)
}
