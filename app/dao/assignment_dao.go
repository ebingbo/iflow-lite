package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

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

func (*AssignmentDao) AssignmentList(ctx context.Context) ([]model.Assignment, error) {
	var items []model.Assignment
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("assignment list error: %+v", err)
		return nil, err
	}
	return items, nil
}
