package services

import (
	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
}

var UserService = new(userService)

func (u *userService) GetByName(name string) (*models.User, error) {
	fmt.Println(name)
	user := &models.User{}
	user.Name = name
	mysql := orm.NewOrm()
	if err := mysql.Read(user, "name"); err != nil {
		return nil, err
	} else {
		return user, err
	}
}

func (u *userService) Create(user *models.User) (bool, error) {
	mysql := orm.NewOrm()
	//user.go.Password = utils.Md5Text(user.go.Password)
	fmt.Println(user.Password)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return false, err
	}

	user.Password = string(password)
	fmt.Println(user.Password)

	_, err = mysql.Insert(user)
	if err != nil {
		return false, err
	}
	return true, err

}

func (u *userService) Query(query string) ([]*models.User, error) {
	mysql := orm.NewOrm()
	var users []*models.User
	querySet := mysql.QueryTable(&models.User{})

	if query != "" {
		cond := orm.NewCondition()
		cond = cond.Or("name__icontains", query)
		cond = cond.Or("nickname__icontains", query)
		cond = cond.Or("tel__icontains", query)
		cond = cond.Or("addr__icontains", query)
		cond = cond.Or("email__icontains", query)
		cond = cond.Or("status__icontains", query)
		querySet = querySet.SetCond(cond)

	}
	rows, err := querySet.All(&users)
	fmt.Println("QueryUser :", rows, err)
	return users, err
}

func (u *userService) GetByID(ID int) *models.User {
	mysql := orm.NewOrm()
	user := &models.User{ID: ID}
	if err := mysql.Read(user); err == nil {
		return user
	}
	return nil
}

func (u *userService) ModifyByForm(form *forms.UserModifyForm) {
	mysql := orm.NewOrm()
	fmt.Println(form)
	if user := u.GetByID(form.ID); user != nil {
		user.Name = form.Name
		user.Password = utils.GeneratePassword(form.Password)
		fmt.Println(user.Password)
		mysql.Update(user, "Name", "Password")
	}
}

func (u *userService) DeleteByID(ID int) {
	mysql := orm.NewOrm()
	mysql.Delete(&models.User{ID: ID})
}

func (u *userService) ModifyPassword(user *models.User, password string) {
	mysql := orm.NewOrm()
	user.Password = utils.GeneratePassword(password)
	mysql.Update(user, "Password")
}
