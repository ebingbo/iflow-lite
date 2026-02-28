package model

import "time"

type Process struct {
	ID          uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	Name        string    `gorm:"column:name;default:''" json:"name" example:"1"`
	Code        string    `gorm:"column:code;default:'';unique" json:"code" example:"1"`
	Description string    `gorm:"column:description;default:''" json:"description" example:"1"`
	Status      *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`
	CreatedBy   uint64    `gorm:"column:created_by;default:0" json:"created_by" example:"1"`
	UpdatedBy   uint64    `gorm:"column:updated_by;default:0" json:"updated_by" example:"1"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Process) TableName() string {
	return "process"
}
