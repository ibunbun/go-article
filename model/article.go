package model

import (
	"time"
)

type Article struct {
	Id      int       `json:"id" gorm:"primary_key"`
	Author  string    `json:"author" gorm:"type:text"`
	Title   string    `json:"title" gorm:"type:text"`
	Body    string    `json:"body" gorm:"type:text"`
	Created time.Time `json:"created" gorm:"type:timestamp"`
}
