package services

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type targetService struct {
}

var TargetService *targetService = new(targetService)

func (s *targetService) Query(query string) (targets []*models.Target, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Target{})

	if query != "" {
		cond := orm.NewCondition()
		cond = cond.Or("id__icontains", query)
		cond = cond.Or("name__icontains", query)
		cond = cond.Or("remark__icontains", query)
		querySet = querySet.SetCond(cond)

	}
	rows, err := querySet.All(&targets)
	fmt.Println("QueryUser :", rows, err)
	return targets, err
}
