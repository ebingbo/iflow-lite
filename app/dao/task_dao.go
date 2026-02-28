package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type TaskDao struct{}

func NewTaskDao() *TaskDao {
	return &TaskDao{}
}
func (*TaskDao) TaskGet(ctx context.Context, id uint64) (*model.Task, error) {
	m := &model.Task{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("task get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*TaskDao) TaskAdd(ctx context.Context, m *model.Task) (*model.Task, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TaskDao) TaskList(ctx context.Context) ([]model.Task, error) {
	var items []model.Task
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task list error: %+v", err)
		return nil, err
	}
	return items, nil
}
