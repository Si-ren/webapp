package forms

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/validation"
	"regexp"
)

type ModifyPasswordForm struct {
	User           *models.User `form:"-"`
	OldPassword    string       `form:"oldpassowrd"`
	NewPassword    string       `form:"newpassword"`
	VerifyPassword string       `form:"verifypassword"`
}

func (f *ModifyPasswordForm) Valid(validation *validation.Validation) {
	fmt.Println("ModifyPasswordForm Valid Start")
	if f.User.ValidPassword(f.OldPassword) {
		fmt.Println("Valid true")
		passwordRegexp := "^[0-9a-zA-Z_.\\$\\!#%^%\\(\\)\\+\\-\\=]{1,20}$"
		validation.Match(f.NewPassword, regexp.MustCompile(passwordRegexp), "PasswordErr.PasswordErr.PasswordErr").Message("密码格式不正确")
		if validation.HasErrors() {
			return
		} else if f.NewPassword != f.VerifyPassword {
			validation.AddError("PasswordErr.PasswordErr", "密码不一致")
		} else if f.OldPassword == f.NewPassword {
			validation.AddError("PasswordErr.PasswordErr", "新旧密码一致")
		}
	} else {
		validation.AddError("PasswordErr", "旧密码错误")
	}
}
