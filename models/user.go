package models

import (
	"gokade1/tools"
	"time"

	"github.com/astaxie/beego/orm"
)

// User 用户对象
type User struct { //这个User要和数据库中的表名一样才能读到数据库中的内容
	ID         int        `orm:"column(id)" form:"id"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)" form:"name"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)" form:"password"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add;" json:"created_at"`
	UpdatedAt  *time.Time `orm:"auto_now;" json:"updated_at"`
	DeletedAt  *time.Time `orm:"null;" json:"-"`
}

// ValidPassword 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return tools.BcryptCheck(password, u.Password) //Bcrypt加密
	// fmt.Println(password, u.Password)
	// return u.Password == tools.Md5Text(password) //md5加密
}

func init() {
	orm.RegisterModel(new(User))
}
