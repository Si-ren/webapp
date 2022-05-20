package services

import (
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type targetService struct {
}

var TargetService *targetService = new(targetService)

func (s *targetService) Query(query string) (targets []*models.Target, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Target{})
	cond := orm.NewCondition().And("delete_at__isnull", true)
	if query != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("id__icontains", query)
		qcond = qcond.Or("name__icontains", query)
		qcond = qcond.Or("remark__icontains", query)
		cond = cond.AndCond(qcond)
	}
	rows, err := querySet.RelatedSel().SetCond(cond).All(&targets)
	fmt.Println("QueryUser :", rows, err)
	return targets, err
}

func (s *targetService) GetByID(ID int) *models.Target {
	mysql := orm.NewOrm()
	target := &models.Target{ID: ID}
	if err := mysql.Read(target); err == nil {
		return target
	}
	return nil
}

func (s *targetService) DeleteByID(ID int) {
	target := s.GetByID(ID)
	if target != nil {
		now := time.Now()
		target.DeleteAt = &now
		mysql := orm.NewOrm()
		mysql.Update(target, "DeleteAt")
	}
}

func (s *targetService) Modify(form *forms.TargetModifyForm) error {
	target := s.GetByID(form.ID)
	if target == nil {
		return fmt.Errorf("job is exits. Create job failed")
	}
	target.Name = form.Name
	target.Remark = form.Remark
	target.Job = JobService.GetByID(form.ID)
	orm.NewOrm().Update(target)
	return nil
}
