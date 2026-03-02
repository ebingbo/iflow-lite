package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultLogDao = NewLogDao()

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

func (*LogDao) LogList(ctx context.Context) ([]*model.Log, error) {
	var items []*model.Log
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("log list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*LogDao) LogQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Log, int64, error) {
	var items = make([]*model.Log, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Log{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("log count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("log query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
