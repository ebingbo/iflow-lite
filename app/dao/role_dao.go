package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
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

func (*RoleDao) RoleGet(ctx context.Context, id uint64) (*model.Role, error) {
	m := &model.Role{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("role get error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*RoleDao) RoleAdd(ctx context.Context, m *model.Role) error {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("role add error: %+v", err)
		return err
	}
	return nil
}

func (*RoleDao) RoleUpdate(ctx context.Context, m *model.Role) error {
	if err := client.MysqlDB.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("role update error: %+v", err)
		return err
	}
	return nil
}

func (*RoleDao) RoleDelete(ctx context.Context, id uint64) error {
	if err := client.MysqlDB.WithContext(ctx).Delete(&model.Role{}, id).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("role delete error: %+v", err)
		return err
	}
	return nil
}

func (*RoleDao) RoleQuery(ctx context.Context, keyword string, page, size int) ([]*model.Role, int64, error) {
	var items = make([]*model.Role, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Role{})
	if keyword != "" {
		db = db.Where("name LIKE ? OR code LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("role count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Order("id DESC").Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("role query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
