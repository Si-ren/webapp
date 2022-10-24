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
		//job下的node下的hostname
		qcond = qcond.Or("job__node__hostname__icontains", query)
		cond = cond.AndCond(qcond)
	}
	// 默认情况下直接调用 RelatedSel 将进行最大 5 层的关系查询
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
	//fmt.Println("targetService Modify form.Job ", form.Job)
	target.Job = JobService.GetByID(form.Job)
	orm.NewOrm().Update(target)
	return nil
}

func (s *targetService) Create(form *forms.TargetCreateForm) (bool, error) {
	mysql := orm.NewOrm()
	_, err := mysql.Insert(form)
	if err != nil {
		return false, err
	}
	return true, err
}
