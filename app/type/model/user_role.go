package model

import "time"

type UserRole struct {
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`
	UserID    uint64    `gorm:"column:user_id;default:0" json:"user_id" example:"1"`
	RoleID    uint64    `gorm:"column:role_id;default:0" json:"role_id" example:"1"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"`
}

func (*UserRole) TableName() string {
	return "user_role"
}
