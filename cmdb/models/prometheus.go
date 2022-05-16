package models

import "time"

type Node struct {
	ID       int        `orm:"column(id)"`
	UUID     string     `orm:"column(uuid);varchar(64)"`
	Hostname string     `orm:"varchar(64)"`
	Addr     string     `orm:"varchar(256)"`
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
	Jobs     []*Job     `orm:"reverse(many)"`
}

type Job struct {
	ID       int        `orm:"column(id)"`
	Key      string     `orm:"varchar(64)"`
	Remark   string     `orm:"varchar(256)"`
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
	Node     *Node      `orm:"rel(fk)"`
	Target   []*Target  `orm:"reverse(many)"`
}

type Target struct {
	ID       int        `orm:"column(id)"`
	Name     string     `orm:"varchar(64)"`
	Remark   string     `orm:"varchar(256)"`
	CreateAt *time.Time `orm:"auto_now_add"`
	UpdateAt *time.Time `orm:"auto_now"`
	DeleteAt *time.Time `orm:"null"`
	Job      *Job       `orm:"rel(fk)"`
}
