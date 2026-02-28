package model

import "time"

type Log struct {
	ID          uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"` // ID
	ProcessID   uint64    `gorm:"column:process_id;default:0" json:"process_id" example:"1"`
	ProcessCode string    `gorm:"column:process_code;default:''" json:"process_code" example:"1"`
	ExecutionID uint64    `gorm:"column:execution_id;default:0" json:"execution_id" example:"1"`
	NodeID      uint64    `gorm:"column:node_id;default:0" json:"node_id" example:"1"`
	NodeCode    string    `gorm:"column:node_code;default:''" json:"node_code" example:"1"`
	TaskID      uint64    `gorm:"column:task_id;default:0" json:"task_id" example:"1"`
	Action      string    `gorm:"column:action;default:''" json:"action" example:"1"`
	AssigneeID  string    `gorm:"column:assignee_id;default:''" json:"assignee_id" example:"1"`
	Remark      string    `gorm:"column:remark;default:''" json:"remark" example:"1"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
}

func (*Log) TableName() string {
	return "log"
}
