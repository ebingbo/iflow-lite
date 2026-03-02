package model

import "time"

type Execution struct {
	ID           uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	ProcessID    uint64    `gorm:"column:process_id;default:0" json:"process_id" example:"1"`
	ProcessCode  string    `gorm:"column:process_code;default:''" json:"process_code" example:"1"`
	ProcessName  string    `gorm:"column:process_name;default:''" json:"process_name" example:"xxx"`
	BusinessKey  string    `gorm:"column:business_key;default:''" json:"business_key" example:"1"`
	BusinessType string    `gorm:"column:business_type;default:''" json:"business_type" example:"order"`
	Status       string    `gorm:"column:status;default:'running'" json:"status" example:"1"`
	Progress     float64   `gorm:"column:progress;default:0" json:"progress" example:"1"`
	CreatedBy    string    `gorm:"column:created_by;default:''" json:"created_by" example:"1"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Execution) TableName() string {
	return "execution"
}
