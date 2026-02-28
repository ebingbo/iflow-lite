package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type LogDao struct{}

func NewLogDao() *LogDao {
	return &LogDao{}
}
func (*LogDao) LogGet(ctx context.Context, id uint64) (*model.Log, error) {
	m := &model.Log{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("log get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*LogDao) LogAdd(ctx context.Context, m *model.Log) (*model.Log, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("log add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*LogDao) LogList(ctx context.Context) ([]model.Log, error) {
	var items []model.Log
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("log list error: %+v", err)
		return nil, err
	}
	return items, nil
}
