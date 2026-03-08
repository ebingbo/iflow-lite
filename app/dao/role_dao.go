package dao

import (
	"context"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"
)

var DefaultRoleDao = NewRoleDao()

type RoleDao struct{}

func NewRoleDao() *RoleDao {
	return &RoleDao{}
}

func (*RoleDao) RoleListForAssignment(ctx context.Context, keyword string, size int) ([]*model.Role, error) {
	if size <= 0 {
		size = 50
	}
	if size > 200 {
		size = 200
	}

	db := client.MysqlDB.WithContext(ctx).Model(&model.Role{})
	if keyword != "" {
		db = db.Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	var items []*model.Role
	if err := db.Order("id DESC").Limit(size).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("role list for assignment error: %+v", err)
		return nil, err
	}
	return items, nil
}
