package dao

import (
	"context"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultTaskCandidateDao = NewTaskCandidateDao()

type TaskCandidateDao struct{}

func NewTaskCandidateDao() *TaskCandidateDao {
	return &TaskCandidateDao{}
}

func (*TaskCandidateDao) TaskCandidateBatchAddWithTransaction(ctx context.Context, tx *gorm.DB, items []*model.TaskCandidate) error {
	if len(items) == 0 {
		return nil
	}
	if err := tx.WithContext(ctx).Create(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task candidate batch add with transaction error: %+v", err)
		return err
	}
	return nil
}

func (*TaskCandidateDao) TaskCandidateExistsWithTransaction(ctx context.Context, tx *gorm.DB, taskID, userID uint64) (bool, error) {
	var count int64
	if err := tx.WithContext(ctx).Model(&model.TaskCandidate{}).Where("task_id = ? AND user_id = ?", taskID, userID).Count(&count).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task candidate exists with transaction error: %+v", err)
		return false, err
	}
	return count > 0, nil
}

func (*TaskCandidateDao) TaskCandidateExists(ctx context.Context, taskID, userID uint64) (bool, error) {
	var count int64
	if err := client.MysqlDB.WithContext(ctx).
		Model(&model.TaskCandidate{}).
		Where("task_id = ? AND user_id = ?", taskID, userID).
		Count(&count).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task candidate exists error: %+v", err)
		return false, err
	}
	return count > 0, nil
}

func (*TaskCandidateDao) TaskCandidateListByTaskID(ctx context.Context, taskID uint64) ([]*model.TaskCandidate, error) {
	var items []*model.TaskCandidate
	if err := client.MysqlDB.WithContext(ctx).
		Model(&model.TaskCandidate{}).
		Where("task_id = ?", taskID).
		Order("id ASC").
		Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task candidate list by task id error: %+v", err)
		return nil, err
	}
	return items, nil
}
