package k8s

import (
	"gokade1/base/auth"
	"gokade1/controllers/user"
	"gokade1/forms"
	"gokade1/services/k8s"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	//重命名
)

type k8sController struct {
	auth.AuthorizationController
}

func (c *k8sController) Prepare() {
	c.Data["nav"] = "k8s"
	if err := user.Loginuser; err == nil {
		c.DestroySession() //销毁sessionid
		c.Redirect(beego.URLFor("AuthController.Login"), 302)
	}
	c.Data["loginuser"] = user.Loginuser //head标签用来标志当前登陆用户

	// fmt.Println("bbbbbbbb:", Loginuser) //测试能不能拿到CC，不能直接拿c.loginuser
	// fmt.Println("aaaaaaaa:", Loginuser.Name)  //登陆用户
}

type DeploymentController struct {
	k8sController
}

func (c *DeploymentController) Query() {
	c.Data["deployments"] = k8s.DeploymentService.Query()
	// q := c.GetString("q")
	// c.Data["q"] = q
	c.TplName = "k8s/query/deployment.html"
}

func (c *DeploymentController) QueryPod() {
	c.Data["pod"] = k8s.DeploymentService.QueryPod()
	// q := c.GetString("q")
	// c.Data["q"] = q
	c.TplName = "k8s/query/pod.html"
}

func (c *DeploymentController) Delete() {

	if user.Loginuser.Gender != 1 {
		c.Abort("NoPermissions")
		return
	}

	name := c.GetString("name")
	namespace := c.GetString("namespace", "default")
	// 数据检查&权限
	k8s.DeploymentService.Delete(name, namespace)
	c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
}

func (c *DeploymentController) DeletePod() {
	if user.Loginuser.Gender != 1 {
		c.Abort("NoPermissions")
		return
	}

	name := c.GetString("name")
	namespace := c.GetString("namespace", "default")
	// 数据检查&权限
	k8s.DeploymentService.DeletePod(name, namespace)
	c.Redirect(beego.URLFor("DeploymentController.QueryPod"), http.StatusFound)
}

func (c *DeploymentController) Create() {

	if user.Loginuser.Gender != 1 {
		c.Abort("NoPermissions")
		return
	}

	form := &forms.DeploymentCreateForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			k8s.DeploymentService.Create(form)
			c.Redirect(beego.URLFor("DeploymentController.Query"), http.StatusFound)
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["namespaces"] = k8s.NamespaceService.Query()
	c.TplName = "k8s/create.html"
}

func (c *DeploymentController) Modify() {

	form := &forms.DeploymentModifyForm{}
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			k8s.DeploymentService.Modify(form.Namespace, form.Name, form.Image, form.Exposes(), form.Replicas)
		}
		c.Redirect(beego.URLFor("DeploymentController.Query"), 302)
	}
	valid := validation.Validation{}
	c.Data["errors"] = valid.ErrorMap()
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "k8s/modify/modify.html"
}
