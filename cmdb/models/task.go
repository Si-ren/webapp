package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name         string     `gorm:"size:16;not null; default:''"`
	Progress     int        `gorm:"not null; default:0"`
	User         string     `gorm:"not null;default:''"`
	Desc         string     `gorm:"column:description;size:512;not null;default:''"`
	Status       int        `gorm:"not null;default:0"`
	CompleteTime *time.Time `gorm:"type:datetime;column:complete_time"`
}

func GetTasks() []Task {
	tasks := make([]Task, 0)
	DB.Table("tasks").Find(tasks)
	fmt.Println("DB get tasks: ", tasks)
	return tasks
}
