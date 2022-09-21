package routers

import (
	"gokade1/controllers"
	"gokade1/controllers/k8s"
	"gokade1/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogger("file", `{"filename" : "logs/cc.log"}`) //定义日志记录
	beego.SetLogFuncCall(true)                              //定义记录日志级别
	beego.SetLevel(beego.LevelDebug)                        //定义记录日志级别
	// beego.BeeLogger.DelLogger("console")                    //不记录控制台日志，但是日志里有

	beego.Router("/", &user.UserController{}, "*:Query")
	beego.ErrorController(&controllers.ErrorController{}) //自定义错误路由
	beego.AutoRouter(&user.AuthController{})              //登陆认证路由
	beego.AutoRouter(&user.HomeController{})              //跳转主页路由
	beego.AutoRouter(&user.UserController{})              //查询用户路由
	beego.AutoRouter(&user.PasswordController{})          //修改用户密码路由

	//k8s控制器
	beego.AutoRouter(&k8s.DeploymentController{})
	beego.AutoRouter(&k8s.SvcController{})
	beego.AutoRouter(&k8s.DaemonsetController{})
	beego.AutoRouter(&k8s.SecretsController{})

}
