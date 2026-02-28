package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type ExecutionDao struct{}

func NewExecutionDao() *ExecutionDao {
	return &ExecutionDao{}
}
func (*ExecutionDao) ExecutionGet(ctx context.Context, id uint64) (*model.Execution, error) {
	m := &model.Execution{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("execution get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*ExecutionDao) ExecutionAdd(ctx context.Context, m *model.Execution) (*model.Execution, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ExecutionDao) ExecutionList(ctx context.Context) ([]model.Execution, error) {
	var items []model.Execution
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution list error: %+v", err)
		return nil, err
	}
	return items, nil
}
