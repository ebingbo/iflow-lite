package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/core/constant"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultProcessDao = NewProcessDao()

type ProcessDao struct{}

func NewProcessDao() *ProcessDao {
	return &ProcessDao{}
}

func (*ProcessDao) ProcessGet(ctx context.Context, id uint64) (*model.Process, error) {
	m := &model.Process{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("process get error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ProcessDao) ProcessDelete(ctx context.Context, id uint64) error {
	m := &model.Process{ID: id}
	if err := client.MysqlDB.WithContext(ctx).Delete(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process delete error: %+v", err)
		return err
	}
	return nil
}

func (*ProcessDao) ProcessDisable(ctx context.Context, id uint64) error {
	m := &model.Process{ID: id}
	if err := client.MysqlDB.WithContext(ctx).Model(m).Update("status", constant.ProcessStatusDisabled).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process disable error: %+v", err)
		return err
	}
	return nil
}

func (*ProcessDao) ProcessEnable(ctx context.Context, id uint64) error {
	m := &model.Process{ID: id}
	if err := client.MysqlDB.WithContext(ctx).Model(m).Update("status", constant.ProcessStatusEnabled).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process enable error: %+v", err)
		return err
	}
	return nil
}

func (*ProcessDao) ProcessTake(ctx context.Context, code string) (*model.Process, error) {
	m := &model.Process{}
	if err := client.MysqlDB.WithContext(ctx).Model(m).Where("code = ?", code).Take(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("process take error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ProcessDao) ProcessTakeWithTransaction(ctx context.Context, tx *gorm.DB, code string) (*model.Process, error) {
	m := &model.Process{}
	if err := tx.Model(m).Where("code = ?", code).Take(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("process take error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ProcessDao) ProcessAdd(ctx context.Context, m *model.Process) (*model.Process, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*ProcessDao) ProcessList(ctx context.Context) ([]*model.Process, error) {
	var items []*model.Process
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*ProcessDao) ProcessQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Process, int64, error) {
	var items []*model.Process
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Process{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("process query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}

func (*ProcessDao) ProcessUpdate(ctx context.Context, m *model.Process) error {
	if err := client.MysqlDB.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("process update error: %+v", err)
		return err
	}
	return nil
}
