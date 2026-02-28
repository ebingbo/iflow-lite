package model

import "time"

type Transition struct {
	ID         uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	ProcessID  uint64    `gorm:"column:process_id;default:0" json:"process_id" example:"1"`
	FromNodeID uint64    `gorm:"column:from_node_id" json:"from_node_id" example:"1"`
	ToNodeID   uint64    `gorm:"column:to_node_id" json:"to_node_id" example:"1"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*Transition) TableName() string {
	return "transition"
}
