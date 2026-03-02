package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultAssignmentDao = NewAssignmentDao()

type AssignmentDao struct{}

func NewAssignmentDao() *AssignmentDao {
	return &AssignmentDao{}
}

func (*AssignmentDao) AssignmentGet(ctx context.Context, id uint64) (*model.Assignment, error) {
	m := &model.Assignment{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("assignment get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*AssignmentDao) AssignmentAdd(ctx context.Context, m *model.Assignment) (*model.Assignment, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("assignment add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*AssignmentDao) AssignmentList(ctx context.Context) ([]*model.Assignment, error) {
	var items []*model.Assignment
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("assignment list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*AssignmentDao) AssignmentQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Assignment, int64, error) {
	var items = make([]*model.Assignment, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Assignment{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("assignment count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("assignment query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
