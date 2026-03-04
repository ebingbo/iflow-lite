package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultExecutionDao = NewExecutionDao()

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

func (*ExecutionDao) ExecutionGetWithTransaction(ctx context.Context, tx *gorm.DB, id uint64) (*model.Execution, error) {
	m := &model.Execution{ID: id}
	if err := tx.WithContext(ctx).First(m).Error; err != nil {
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

func (*ExecutionDao) ExecutionAddWithTransaction(ctx context.Context, tx *gorm.DB, m *model.Execution) (*model.Execution, error) {
	if err := tx.Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ExecutionDao) ExecutionUpdate(ctx context.Context, m *model.Execution) error {
	if err := client.MysqlDB.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution update error: %+v", err)
		return err
	}
	return nil
}

func (*ExecutionDao) ExecutionUpdateWithTransaction(ctx context.Context, tx *gorm.DB, m *model.Execution) error {
	if err := tx.Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution update error: %+v", err)
		return err
	}
	return nil
}

func (*ExecutionDao) ExecutionList(ctx context.Context) ([]*model.Execution, error) {
	var items []*model.Execution
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*ExecutionDao) ExecutionQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Execution, int64, error) {
	var items = make([]*model.Execution, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Execution{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("execution count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("execution query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
