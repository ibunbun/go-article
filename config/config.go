package config

import (
	"log"

	"kumparan/model"

	"github.com/jinzhu/gorm"
)

func DbConnect() *gorm.DB {
	consStr := "root:root@tcp(kumparan-db)/kumparan?parseTime=true"
	db, err := gorm.Open("mysql", consStr)
	if err != nil {
		log.Fatal("Error when connect db" + consStr + " : " + err.Error())
		return nil
	}

	db.Debug().AutoMigrate(model.Article{})

	return db
}
