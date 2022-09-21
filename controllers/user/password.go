package user

import (
	"gokade1/base/auth"
	"gokade1/base/errors"
	"gokade1/forms"
	"gokade1/services"
	"regexp"

	"github.com/astaxie/beego/validation"
)

//用户密码修改控制器
type PasswordController struct {
	auth.AuthorizationController
}

//用户修改密码
func (c *PasswordController) Modify() {
	form := &forms.PassWordModifyFrom{}
	text := ""
	errors := errors.New()
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if ok := c.LoginUser.ValidPassword(form.OldPassWord); !ok {
				errors.Add("default", "旧密码错误")
			} else {
				guifan := "^[0-9a-zA-Z_.\\$\\!#%^&\\*\\(\\)\\+]{6,20}$"
				validation := &validation.Validation{}                                                                    //验证数据
				validation.Match(form.PassWord, regexp.MustCompile(guifan), "default.default.default").Message("密码格式不正确") //验证正则表达式
				//验证密码范围数字，大小写英文字母，特殊字符（_.$!#%^&*()+)s
				if validation.HasErrors() {
					for key, error := range validation.ErrorsMap {
						for _, err := range error {
							errors.Add(key, err.Message)
						}
					}
				} else if form.PassWord != form.PassWord2 {
					errors.Add("default", "密码错误，两次输入密码不一致")

				} else if form.OldPassWord == form.PassWord2 {
					errors.Add("default", "新旧密码不能一致")
				} else {
					services.UserService.ModifyPassWord(c.LoginUser.ID, form.PassWord)
					text = "密码修改成功"
				}
			}
		}
	}
	c.TplName = "password/modify.html"
	c.Data["errors"] = errors
	c.Data["text"] = text
	// c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	// fmt.Println(template.HTML(c.XSRFFormHTML()))
}

