package services

import (
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type jobService struct {
}

var JobService *jobService = new(jobService)

func (s *jobService) Query(query string) (jobs []*models.Job, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Job{})
	cond := orm.NewCondition().And("delete_at__isnull", true)
	if query != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("key__icontains", query)
		qcond = qcond.Or("id__icontains", query)
		qcond = qcond.Or("remark__icontains", query)
		cond = cond.AndCond(qcond)
	}
	rows, err := querySet.SetCond(cond).All(&jobs)
	//RelatedSel()只能自动联查fk
	//querySet.RelatedSel().SetCond(cond).All(&jobs)
	fmt.Println("QueryUser :", rows, err)
	for _, job := range jobs {
		mysql.LoadRelated(job, "Node")
		//fmt.Printf("%#v\n", job.Node)
	}
	return jobs, err
}

func (s *jobService) GetByID(ID int) *models.Job {
	mysql := orm.NewOrm()
	job := &models.Job{ID: ID}
	if err := mysql.Read(job); err == nil {
		fmt.Printf("%#v\n", job)
		fmt.Printf("%#v\n", job.Node) //&models.Node{ID:5, UUID:"", Hostname:"", Addr:"", CreateAt:<nil>, UpdateAt:<nil>, DeleteAt:<nil>, Jobs:[]*models.Job(nil)}
		mysql.LoadRelated(job, "Node")
		fmt.Printf("%#v\n", job.Node) //&models.Node{ID:5, UUID:"abc", Hostname:"11111", Addr:"11111", CreateAt:time.Date(2022, time.May, 19, 11, 24, 22, 0, time.Location("Local")), UpdateAt:time.Date(2022, time.May, 19, 11, 24, 24, 0, time.Location("Local")), DeleteAt:<nil>, Jobs:[]*models.Job(nil)}

		return job
	}
	return nil
}

func (s *jobService) DeleteByID(ID int) {
	job := s.GetByID(ID)
	if job != nil {
		now := time.Now()
		job.DeleteAt = &now
		mysql := orm.NewOrm()
		mysql.Update(job, "DeleteAt")
	}
}

func (s *jobService) Create(form *forms.JobCreateForm) error {
	//需要先验证
	job := &models.Job{
		Key:    form.Key,
		Remark: form.Remark,
		Node:   NodeService.GetByID(form.Node),
	}
	if _, err := orm.NewOrm().Insert(job); err != nil {
		return err
	}
	return nil
}

func (s *jobService) Modify(form *forms.JobModifyForm) error {
	job := s.GetByID(form.ID)
	if job == nil {
		return fmt.Errorf("job is exits. Create job failed")
	}
	job.Key = form.Key
	job.Remark = form.Remark
	job.Node = NodeService.GetByID(form.ID)
	orm.NewOrm().Update(job)
	return nil
}

func (s *jobService) QueryByUUID(uuid string) []*models.Job {
	var jobs []*models.Job
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Job{})
	querySet.RelatedSel().Filter("delete_at__isnull", true).Filter("node_id__uuid", uuid).All(&jobs)
	//反向关联,反向关联jobs中的targets,即target表中的job
	for _, job := range jobs {
		mysql.LoadRelated(job, "Target")
	}
	return jobs
}
