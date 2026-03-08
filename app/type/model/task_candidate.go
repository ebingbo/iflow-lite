package model

import "time"

type TaskCandidate struct {
	ID         uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`
	TaskID     uint64    `gorm:"column:task_id;default:0" json:"task_id" example:"1"`
	UserID     uint64    `gorm:"column:user_id;default:0" json:"user_id" example:"1"`
	SourceType string    `gorm:"column:source_type;default:'user'" json:"source_type" example:"user"`
	SourceID   uint64    `gorm:"column:source_id;default:0" json:"source_id" example:"1"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"`
}

func (*TaskCandidate) TableName() string {
	return "task_candidate"
}
