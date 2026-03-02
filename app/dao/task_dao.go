package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultTaskDao = NewTaskDao()

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

func (*TaskDao) TaskList(ctx context.Context) ([]*model.Task, error) {
	var items []*model.Task
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*TaskDao) TaskQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Task, int64, error) {
	var items = make([]*model.Task, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Task{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("task query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
