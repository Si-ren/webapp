package models

import (
	"cmdb/utils"
	"gorm.io/gorm"
)

func init() {
	if !DB.Migrator().HasTable(&User{}) {
		DB.AutoMigrate(&User{}, &Task{})
		passwd := utils.GeneratePassword("siri")
		DB.Create(&User{
			Model:      gorm.Model{},
			StaffID:    "",
			Name:       "siri",
			NickName:   "siri",
			Password:   passwd,
			Gender:     1,
			Tel:        "13391110017",
			Addr:       "王家村",
			Email:      "834555340@qq.com",
			Department: "Devops",
			Status:     0,
		})
	}
}
