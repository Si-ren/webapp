package services

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type jobService struct {
}

var JobService *jobService = new(jobService)

func (s *jobService) Query(query string) (jobs []*models.Job, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Job{})

	if query != "" {
		cond := orm.NewCondition()
		cond = cond.Or("key__icontains", query)
		cond = cond.Or("id__icontains", query)
		cond = cond.Or("remark__icontains", query)
		querySet = querySet.SetCond(cond)

	}
	rows, err := querySet.All(&jobs)
	fmt.Println("QueryUser :", rows, err)
	return jobs, err
}
