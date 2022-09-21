package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB //定义成指针要引用

func init() {
	orm.Debug = true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/goweb?charset=utf8mb4&loc=PRC&parseTime=true",
		beego.AppConfig.DefaultString("mysql::User", "root"),
		beego.AppConfig.DefaultString("mysql::Password", "@Cys000522"),
		beego.AppConfig.DefaultString("mysql::Host", "101.34.44.247"),
		beego.AppConfig.DefaultInt("mysql::Port", 3306),
	)
	orm.RegisterDriver("mysql", orm.DRMySQL)       //注册数据库驱动
	orm.RegisterDataBase("default", "mysql", dsn)  //注册数据库
	//"root:@Cys000522@tcp(49.235.122.254:3306)/goweb?charset=utf8mb4&loc=PRC&parseTime=true"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
