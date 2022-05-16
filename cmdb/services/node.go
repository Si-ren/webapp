package services

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type nodeService struct {
}

var NodeService *nodeService = new(nodeService)

func (s *nodeService) Query(query string) (nodes []*models.Node, err error) {
	mysql := orm.NewOrm()
	querySet := mysql.QueryTable(&models.Node{})

	if query != "" {
		cond := orm.NewCondition()
		cond = cond.Or("id__icontains", query)
		cond = cond.Or("uuid__icontains", query)
		cond = cond.Or("hostname__icontains", query)
		cond = cond.Or("addr__icontains", query)
		querySet = querySet.SetCond(cond)

	}
	rows, err := querySet.All(&nodes)
	fmt.Println("QueryUser :", rows, err)
	return nodes, err
}
