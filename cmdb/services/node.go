package services

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type nodeService struct {
}

var NodeService *nodeService = new(nodeService)

func (s *nodeService) Query(query string) (nodes []*models.Node, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Node{})
	cond := orm.NewCondition().And("delete_at__isnull", true)
	if query != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("id__icontains", query)
		qcond = qcond.Or("uuid__icontains", query)
		qcond = qcond.Or("hostname__icontains", query)
		qcond = qcond.Or("addr__icontains", query)
		cond = cond.AndCond(qcond)
	}
	rows, err := querySet.SetCond(cond).All(&nodes)
	fmt.Println("QueryUser :", rows, err)
	return nodes, err
}

func (s *nodeService) GetByID(ID int) *models.Node {
	mysql := orm.NewOrm()
	node := &models.Node{ID: ID}
	if err := mysql.Read(node); err == nil {
		return node
	}
	return nil
}

func (s *nodeService) DeleteByID(ID int) {
	node := s.GetByID(ID)
	if node != nil {
		now := time.Now()
		node.DeleteAt = &now
		mysql := orm.NewOrm()
		mysql.Update(node, "DeleteAt")
	}
}
