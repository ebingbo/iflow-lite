package model

import "time"

type Role struct {
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`
	Name      string    `gorm:"column:name;default:''" json:"name" example:"系统管理员"`
	Code      string    `gorm:"column:code;default:''" json:"code" example:"admin"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"`
}

func (*Role) TableName() string {
	return "role"
}
