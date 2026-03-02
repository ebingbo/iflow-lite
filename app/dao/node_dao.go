package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultNodeDao = NewNodeDao()

type NodeDao struct{}

func NewNodeDao() *NodeDao {
	return &NodeDao{}
}
func (*NodeDao) NodeGet(ctx context.Context, id uint64) (*model.Node, error) {
	m := &model.Node{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("node get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*NodeDao) NodeAdd(ctx context.Context, m *model.Node) (*model.Node, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("node add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*NodeDao) NodeList(ctx context.Context, cond map[string]interface{}) ([]*model.Node, error) {
	var items []*model.Node
	if err := client.MysqlDB.WithContext(ctx).Model(&model.Node{}).Where(cond).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("node list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*NodeDao) NodeQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Node, int64, error) {
	var items = make([]*model.Node, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Node{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("node count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("node query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}
