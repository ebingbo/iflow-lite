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

var DefaultUserDao = NewUserDao()

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

func (*UserDao) UserUpdate(ctx context.Context, m *model.User) error {
	if err := client.MysqlDB.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user update error: %+v", err)
		return err
	}
	return nil
}

func (*UserDao) UserList(ctx context.Context) ([]*model.User, error) {
	var items []*model.User
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*UserDao) UserListByIDs(ctx context.Context, ids []uint64) ([]*model.User, error) {
	var items []*model.User
	if err := client.MysqlDB.WithContext(ctx).Model(&model.User{}).Where("id IN ?", ids).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user list by ids error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*UserDao) UserQuery(ctx context.Context, input *input.UserQueryInput) (items []*model.User, total int64, err error) {
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

func (*UserDao) UserListForAssignment(ctx context.Context, keyword string, size int) ([]*model.User, error) {
	if size <= 0 {
		size = 50
	}
	if size > 200 {
		size = 200
	}

	db := client.MysqlDB.WithContext(ctx).Model(&model.User{}).Where("status = 1")
	if keyword != "" {
		db = db.Where("name LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	var items []*model.User
	if err := db.Order("id DESC").Limit(size).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user list for assignment error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*UserDao) UserIDListByRoleIDWithTransaction(ctx context.Context, tx *gorm.DB, roleID uint64) ([]uint64, error) {
	var userIDs []uint64
	if err := tx.WithContext(ctx).
		Table("user_role").
		Select("DISTINCT user_id").
		Where("role_id = ?", roleID).
		Pluck("user_id", &userIDs).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user id list by role id with transaction error: %+v", err)
		return nil, err
	}
	return userIDs, nil
}
