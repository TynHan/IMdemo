package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   uint
	TargetId uint
	Type     string
	Media    int
	Content  string
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}
