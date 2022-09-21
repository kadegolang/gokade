package services

import (
	"gokade1/forms"
	"gokade1/models"
	"gokade1/tools"

	"github.com/astaxie/beego/orm"
)

type userService struct{}

// GetUserByID 通过id获取用户
func (s *userService) GetByID(k int) *models.User {
	user := &models.User{ID: k} //初始化id
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

// GetUserByName 通过用户名获取用户
func (s *userService) GetByName(name string) *models.User {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// QueryUser 查询用户
func (s *userService) Query(q string) []*models.User {
	var users []*models.User
	queryset := orm.NewOrm().QueryTable(&models.User{}) // 获取 QuerySeter 对象，user 为表名
	if q != "" {                                        //查询条件主页查询的q变量
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("staff_id__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("department__icontains", q)
		queryset = queryset.SetCond(cond)
	}
	queryset.All(&users) //查询user所有参数
	return users
}

//修改用户信息
func (s *userService) Modify(form *forms.UseModifyFrom) {
	if user := s.GetByID(form.ID); user != nil {
		user.Nickname = form.Nickname
		user.Tel = form.Tel
		user.Addr = form.Addr
		user.Department = form.Department
		user.Email = form.Email
		user.Status = form.Status
		ormer := orm.NewOrm()
		ormer.Update(user, "Nickname","Tel","Addr","Department","Email","Status")
	}

}

//删除用户信息
func (s *userService) Delete(k int) {
	ormer := orm.NewOrm()
	ormer.Delete(&models.User{ID: k})
}

//修改用户密码信息
func (s *userService) ModifyPassWord(id int, password string) {
	if user := s.GetByID(id); user != nil {
		user.Password = tools.Bcrypt(password)
		ormer := orm.NewOrm()
		ormer.Update(user, "Password")
	}
}

//添加用户信息
func AddUser(id, status int, staffid, name, password, tel, email, addr, department, nickname string) {
	user := &models.User{
		ID:         id,
		Status:     status,
		StaffID:    staffid,
		Name:       name,
		Password:   password,
		Tel:        tel,
		Email:      email,
		Addr:       addr,
		Department: department,
		Nickname:   nickname,
	}
	ormer := orm.NewOrm()
	ormer.Insert(user)
}

//用户操作服务
var UserService = new(userService) //通过new函数实例化一个结构体指针
