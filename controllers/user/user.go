package user

import (
	"fmt"
	"gokade1/base/auth"
	"gokade1/forms"
	"gokade1/tools"
	"net/http"

	"gokade1/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// UserController 用户管理控制器
type UserController struct {
	auth.AuthorizationController
}

// Query 查询用户
func (c *UserController) Query() {
	//读取消息，页面修改成功的消息flash
	flash := beego.ReadFromRequest(&c.Controller)
	fmt.Println(flash.Data)
	q := c.GetString("q")

	c.Data["users"] = services.UserService.Query(q)
	c.Data["q"] = q

	// c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/query.html"
}

//modify 修改用户
func (c *UserController) Modify() {
	//控制编辑删除操作
	if c.LoginUser.Gender != 1 {
		c.Abort("NoPermissions")
		return
	}

	//GET获取数据
	//POSt修改用户
	form := &forms.UseModifyFrom{}
	// form := &models.User{}

	if c.Ctx.Input.IsPost() { //c.Ctx.Input.IsPost判断是post提交，
		if err := c.ParseForm(form); err == nil { //ParseForm解析form数据
			services.UserService.Modify(form) //修改数据
			//存储消息
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			flash.Store(&c.Controller)

			c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
		}

	} else if pk, err := c.GetInt("k1"); err == nil { //http://localhost:8888/user/modify?k1=2 k1是页面请求参数  获取数据
		if user := services.UserService.GetByID(pk); user != nil {
			form.ID = user.ID //get请求，原本用户名显示
			form.Nickname = user.Nickname
			form.Tel = user.Tel
			form.Addr = user.Addr
			form.Department = user.Department
			form.Email = user.Email
			form.Status = user.Status
		}
	}
	// fmt.Println(form)
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken() //生成一个xsrf随机token防着sxrf攻击 html 需要写一个input标签 比template.HTML(c.XSRFFormHTML())稳定一点
	c.TplName = "user/modify.html"
}

//Delete 删除用户
func (c *UserController) Delete() {
	if k, err := c.GetInt("k2"); err == nil && c.LoginUser.ID != k {
		//控制编辑删除操作
		if c.LoginUser.Gender != 1 {
			c.Abort("NoPermissions")
			return
		}
		services.UserService.Delete(k)
		//存储消息
		flash := beego.NewFlash()
		flash.Set("notice", "删除用户成功")
		flash.Store(&c.Controller)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}

func (c *UserController) Create() {

	//控制编辑删除操作
	if c.LoginUser.Gender != 1 {
		c.Abort("NoPermissions")
		return
	}

	// var form forms.UseCreateFrom
	// form := &models.User{}
	form := &forms.UseCreateFrom{}
	vaild := validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if success, err := vaild.Valid(form); err == nil && success {
				services.AddUser(form.ID, form.Status, form.StaffID, form.Name, tools.Bcrypt(form.Password), form.Tel, form.Email, form.Addr, form.Department, form.Nickname)
				flash := beego.NewFlash()
				flash.Set("success", "新建成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("UserController.Query"), 302)
				return
			}
		}
	}
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["errors"] = vaild.ErrorMap() //后面研究作用
	c.Data["form"] = form
	c.TplName = "user/create.html"
}
