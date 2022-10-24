package controllers

import (
	"cmdb/base/error"
	"cmdb/forms"
	"cmdb/services"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"html/template"
	"net/http"
)

type PasswordController struct {
	Authentication
}

func (c *PasswordController) Modify() {
	modifyPasswordForm := &forms.ModifyPasswordForm{User: c.LoginUser}
	errors := error.New()
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(modifyPasswordForm); err == nil {
			fmt.Println("ParseForm(modifyPasswordForm) : ", modifyPasswordForm.OldPassword, modifyPasswordForm.NewPassword, modifyPasswordForm.VerifyPassword)
			//	验证密码
			//if !c.LoginUser.ValidPassword(modifyPasswordForm.OldPassword) {
			//	errors.AddError("PasswordErr", "旧密码错误")
			//} else {
			//	if _, err := regexp.MatchString("^[0-9a-zA-Z_.\\$\\!#%^%\\(\\)\\+\\-\\=]{6,20}$",
			//		modifyPasswordForm.NewPassword); err != nil {
			//		errors.AddError("PasswordErr", "密码只能用由大小写英文字母,数字,特殊字符组成")
			//	}
			//	if modifyPasswordForm.NewPassword != modifyPasswordForm.VerifyPassword {
			//		errors.AddError("PasswordErr", "密码不一致")
			//	} else if modifyPasswordForm.OldPassword == modifyPasswordForm.NewPassword {
			//		errors.AddError("PasswordErr", "新旧密码一致")
			//	} else {
			//		models.ModifyUserPassword(c.LoginUser, modifyPasswordForm.NewPassword)
			//		c.DestroySession()
			//		c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
			//	}
			//}

			valid := &validation.Validation{}
			//modifyPasswordForm.Valid(valid)
			if success, err := valid.Valid(modifyPasswordForm); err != nil {
				errors.AddError("PasswordErr", err.Error())
			} else if !success {
				errors.AddValidation(valid)
			} else {
				fmt.Println(valid)
				services.UserService.ModifyPassword(c.LoginUser, modifyPasswordForm.NewPassword)
				c.Data["text"] = "修改密码成功"
				c.DestroySession()
				c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
			}

		}
	}
	c.Data["error"] = errors
	c.TplName = "password/modify.html"
	//自己定义标签,传递key value ,key为"_xsrf" value=一串随机字符串
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	fmt.Println(template.HTML(c.XSRFFormHTML()))
}
