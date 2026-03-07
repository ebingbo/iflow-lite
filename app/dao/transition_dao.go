package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultTransitionDao = NewTransitionDao()

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

func (*TransitionDao) TransitionDelete(ctx context.Context, id uint64) error {
	m := &model.Transition{ID: id}
	if err := client.MysqlDB.WithContext(ctx).Delete(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition delete error: %+v", err)
		return err
	}
	return nil
}
func (*TransitionDao) TransitionAdd(ctx context.Context, m *model.Transition) (*model.Transition, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TransitionDao) TransitionUpdate(ctx context.Context, m *model.Transition) error {
	if err := client.MysqlDB.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition update error: %+v", err)
		return err
	}
	return nil
}

func (*TransitionDao) TransitionList(ctx context.Context, cond map[string]interface{}) ([]*model.Transition, error) {
	var items = make([]*model.Transition, 0)
	if err := client.MysqlDB.WithContext(ctx).Model(model.Transition{}).Where(cond).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("transition list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*TransitionDao) FromNodeIDListWithTransaction(ctx context.Context, tx *gorm.DB, processID uint64, toNodeID uint64) ([]uint64, error) {
	var fromNodeIDs []uint64
	if err := tx.Model(&model.Transition{}).Where("process_id = ? AND to_node_id = ?", processID, toNodeID).Pluck("from_node_id", &fromNodeIDs).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("from node id list error: %+v", err)
		return nil, err
	}
	return fromNodeIDs, nil
}

func (*TransitionDao) ToNodeIDListWithTransaction(ctx context.Context, tx *gorm.DB, processID uint64, fromNodeID uint64) ([]uint64, error) {
	var toNodeIDs []uint64
	if err := tx.Model(&model.Transition{}).Where("process_id = ? AND from_node_id = ?", processID, fromNodeID).Pluck("to_node_id", &toNodeIDs).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("to node id list error: %+v", err)
		return nil, err
	}
	return toNodeIDs, nil
}
