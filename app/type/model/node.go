package model

import "time"

type Node struct {
	ID          uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	ProcessID   uint64    `gorm:"column:process_id;default:0" json:"process_id" example:"1"`
	ProcessCode string    `gorm:"column:process_code;default:''" json:"process_code" example:"1"`
	Tag         string    `gorm:"column:tag;default:''" json:"tag" example:"1"`
	Name        string    `gorm:"column:name;default:''" json:"name" example:"1"`
	Code        string    `gorm:"column:code;default:'';unique" json:"code" example:"1"`
	Type        string    `gorm:"column:type;default:'start'" json:"type" example:"start"`
	Description string    `gorm:"column:description;default:''" json:"description" example:"1"`
	CreateBy    uint64    `gorm:"column:create_by;default:0" json:"create_by" example:"1"`
	UpdateBy    uint64    `gorm:"column:update_by;default:0" json:"update_by" example:"1"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Node) TableName() string {
	return "node"
}
