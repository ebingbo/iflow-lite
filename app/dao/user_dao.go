package dao

import (
	"context"
	"errors"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/input"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type UserDao struct{}

func NewUserDao() *UserDao {
	return &UserDao{}
}
func (*UserDao) UserGet(ctx context.Context, id uint64) (*model.User, error) {
	m := &model.User{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("user get error: %+v", err)
		return nil, err
	}
	return m, nil
}
func (*UserDao) UserAdd(ctx context.Context, m *model.User) error {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user add error: %+v", err)
		return err
	}
	return nil
}

func (*UserDao) UserList(ctx context.Context) ([]model.User, error) {
	var items []model.User
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*UserDao) UserQuery(ctx context.Context, input *input.UserQueryInput) (items []model.User, total int64, err error) {
	offset := (input.Page - 1) * input.Size
	tx := client.MysqlDB.WithContext(ctx).Model(&model.User{})
	// 动态条件
	if input.Name != "" {
		tx = tx.Where("name LIKE ?", input.Name+"%")
	}

	if input.Email != "" {
		tx = tx.Where("email = ?", input.Email)
	}

	// status 用指针判断
	if input.Status != nil {
		tx = tx.Where("status = ?", *input.Status)
	}

	// 总数
	if err = tx.Count(&total).Error; err != nil {
		return
	}

	err = tx.
		Order("id DESC").
		Offset(offset).
		Limit(input.Size).
		Find(&items).Error
	return
}
func (*UserDao) UserTake(ctx context.Context, email string) (*model.User, error) {
	var item model.User
	if err := client.MysqlDB.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Take(&item).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user take error: %+v", err)
		return nil, err
	}
	return &item, nil
}
