package cmds


import (
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"    //注释这个分前后

	_ "gokade1/routers"

)

func Webcc()  {
	beego.Run()
}