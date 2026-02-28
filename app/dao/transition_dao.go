package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type TransitionDao struct{}

func NewTransitionDao() *TransitionDao {
	return &TransitionDao{}
}
func (*TransitionDao) TransitionGet(ctx context.Context, id uint64) (*model.Transition, error) {
	m := &model.Transition{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("transition get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*TransitionDao) TransitionAdd(ctx context.Context, m *model.Transition) (*model.Transition, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TransitionDao) TransitionList(ctx context.Context) ([]model.Transition, error) {
	var items []model.Transition
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition list error: %+v", err)
		return nil, err
	}
	return items, nil
}
