package main

import "gokade1/cmds"

// "github.com/astaxie/beego"

// _ "github.com/go-sql-driver/mysql"    //注释这个分前后
// _ "gokade/routers"

func main() {
	cmds.Webcc()
	// web.CreateDeployment()
	// web.Podlist()
}

//GOARCH="amd64"  GOOS="darwin" 交叉编译，mac上写，要改编译环境 例如 set GOOS="linux" (linux )  GOOS="linux" (windows)  GOOS="darwin" (mac) 然后在go build

//go build 后要把conf和views目录 ,static 传进去
