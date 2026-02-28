package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

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

func (*NodeDao) NodeList(ctx context.Context) ([]model.Node, error) {
	var items []model.Node
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("node list error: %+v", err)
		return nil, err
	}
	return items, nil
}
