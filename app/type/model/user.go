package model

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;AUTO_INCREMENT;column:id;" json:"id" example:"1"`                        // ID
	Name      string    `gorm:"column:name;default:''" json:"name" example:"bill"`                                 // 	用户名
	Email     string    `gorm:"unique;column:email;default:''" json:"email" example:"bill@126.com"`                // 邮箱
	Password  string    `gorm:"column:password;default:''" json:"password" example:"123456"`                       // 用户密码
	Status    *uint8    `gorm:"column:status;default:1" json:"status" example:"1"`                                 // 状态: 1启用, 0禁用
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at" example:"2020-10-11T10:10:10"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at" example:"2020-10-11T10:10:10"` // 更新时间
}

func (*User) TableName() string {
	return "user"
}
