package dao

import (
	"context"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultUserRoleDao = NewUserRoleDao()

type UserRoleDao struct{}

func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{}
}

func (*UserRoleDao) UserRoleListByUserID(ctx context.Context, userID uint64) ([]uint64, error) {
	var ids []uint64
	if err := client.MysqlDB.WithContext(ctx).
		Model(&model.UserRole{}).
		Where("user_id = ?", userID).
		Pluck("role_id", &ids).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("user role list error: %+v", err)
		return nil, err
	}
	return ids, nil
}

func (*UserRoleDao) UserRoleReplace(ctx context.Context, userID uint64, roleIDs []uint64) error {
	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("user role delete error: %+v", err)
			return err
		}
		if len(roleIDs) == 0 {
			return nil
		}
		items := make([]*model.UserRole, 0, len(roleIDs))
		seen := make(map[uint64]struct{}, len(roleIDs))
		for _, roleID := range roleIDs {
			if roleID == 0 {
				continue
			}
			if _, ok := seen[roleID]; ok {
				continue
			}
			seen[roleID] = struct{}{}
			items = append(items, &model.UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}
		if len(items) == 0 {
			return nil
		}
		if err := tx.Create(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("user role add error: %+v", err)
			return err
		}
		return nil
	})
}
