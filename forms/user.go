package forms

//用户修改表单
type UseModifyFrom struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
	Nickname string `form:"nickname"`
	Tel        string `form:"tel"`
	Email      string `form:"email"`
	Addr       string `form:"addr"`
	Department string `form:"department"`
	Status     int    `form:"status"`
	StaffID    string `orm:"column(staff_id);size(32)" form:"staffid"`
}

//添加用户表单
type UseCreateFrom struct {
	ID         int    `form:"id"`
	StaffID    string `orm:"column(staff_id);size(32)" form:"staffid"`
	Name       string `form:"name"`
	Password   string `form:"password"`
	Tel        string `form:"tel"`
	Email      string `form:"email"`
	Addr       string `form:"addr"`
	Department string `form:"department"`
	Status     int    `form:"status"`
	Nickname   string  `orm:"size(64)" form:"nickname"`
}
