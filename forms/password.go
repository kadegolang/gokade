package forms

//用户密码修改表单
type PassWordModifyFrom struct {
	OldPassWord string `form:"oldpassword"`
	PassWord    string `form:"password"`
	PassWord2   string `form:"password2"`
}
