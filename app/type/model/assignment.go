package model

import "time"

type Assignment struct {
	ID            uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	ProcessID     uint64    `gorm:"column:process_id;default:0" json:"process_id" example:"1"`
	ProcessCode   string    `gorm:"column:process_code;default:''" json:"process_code" example:"1"`
	NodeID        uint64    `gorm:"column:node_id;default:0" json:"node_id" example:"1"`
	NodeCode      string    `gorm:"column:node_code;default:''" json:"node_code" example:"1"`
	PrincipalType string    `gorm:"column:principal_type;default:'user'" json:"principal_type" example:"user"`
	PrincipalID   uint64    `gorm:"column:principal_id;default:0" json:"principal_id" example:"1"`
	Priority      int       `gorm:"column:priority;default:0" json:"priority" example:"1"`
	Strategy      string    `gorm:"column:strategy;default:'sequential'" json:"strategy" example:"sequential"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Assignment) TableName() string {
	return "assignment"
}
