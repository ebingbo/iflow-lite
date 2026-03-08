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

var DefaultTaskDao = NewTaskDao()

type TaskDao struct{}

func NewTaskDao() *TaskDao {
	return &TaskDao{}
}
func (*TaskDao) TaskGet(ctx context.Context, id uint64) (*model.Task, error) {
	m := &model.Task{ID: id}
	if err := client.MysqlDB.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("task get error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TaskDao) TaskGetWithTransaction(ctx context.Context, tx *gorm.DB, id uint64) (*model.Task, error) {
	m := &model.Task{ID: id}
	if err := tx.WithContext(ctx).First(m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.ServiceLogger.WithContext(ctx).Errorf("task get error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TaskDao) TaskAdd(ctx context.Context, m *model.Task) (*model.Task, error) {
	if err := client.MysqlDB.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task add error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TaskDao) TaskAddWithTransaction(ctx context.Context, tx *gorm.DB, m *model.Task) (*model.Task, error) {
	if err := tx.WithContext(ctx).Create(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task add with transaction error: %+v", err)
		return nil, err
	}
	return m, nil
}

func (*TaskDao) TaskUpdateWithTransaction(ctx context.Context, tx *gorm.DB, m *model.Task) error {
	if err := tx.WithContext(ctx).Save(m).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task update with transaction error: %+v", err)
		return err
	}
	return nil
}

func (*TaskDao) TaskClaimWithTransaction(ctx context.Context, tx *gorm.DB, taskID, assigneeID uint64) (int64, error) {
	nowExpr := tx.WithContext(ctx).NowFunc()
	result := tx.WithContext(ctx).
		Model(&model.Task{}).
		Where("id = ? AND status = ? AND assignee_id = 0", taskID, constant.TaskStatusPending).
		Updates(map[string]interface{}{
			"assignee_id": assigneeID,
			"status":      constant.TaskStatusRunning,
			"claimed_at":  nowExpr,
			"started_at":  nowExpr,
		})
	if result.Error != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task claim with transaction error: %+v", result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
func (*TaskDao) TaskList(ctx context.Context) ([]*model.Task, error) {
	var items []*model.Task
	if err := client.MysqlDB.WithContext(ctx).Find(&items).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task list error: %+v", err)
		return nil, err
	}
	return items, nil
}

func (*TaskDao) TaskQuery(ctx context.Context, cond map[string]interface{}, page, size int) ([]*model.Task, int64, error) {
	var items = make([]*model.Task, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).Model(&model.Task{}).Where(cond)
	if err := db.Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("task query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}

func (*TaskDao) TaskClaimableQueryByUser(ctx context.Context, userID uint64, status string, page, size int) ([]*model.Task, int64, error) {
	var items = make([]*model.Task, 0)
	var total int64
	db := client.MysqlDB.WithContext(ctx).
		Table("task t").
		Select("t.*").
		Joins("JOIN task_candidate tc ON tc.task_id = t.id").
		Where("tc.user_id = ?", userID)
	if status != "" {
		db = db.Where("t.status = ?", status)
	}
	if err := db.Distinct("t.id").Count(&total).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task claimable count error: %+v", err)
		return nil, 0, err
	}
	if total > 0 {
		if err := db.
			Distinct("t.id").
			Order("t.id DESC").
			Offset((page - 1) * size).
			Limit(size).
			Find(&items).Error; err != nil {
			logger.ServiceLogger.WithContext(ctx).Errorf("task claimable query error: %+v", err)
			return nil, 0, err
		}
	}
	return items, total, nil
}

func (*TaskDao) TaskCompleted(ctx context.Context, tx *gorm.DB, executionID uint64, nodeIDs []uint64) (bool, error) {
	var count int64
	if err := tx.Model(&model.Task{}).Where("execution_id = ? AND node_id IN ? AND status NOT IN ?", executionID, nodeIDs, []string{constant.TaskStatusCompleted, constant.TaskStatusSkipped}).Count(&count).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task completed error: %+v", err)
		return false, err
	}
	return count == 0, nil
}

func (*TaskDao) TaskCountWithTransaction(ctx context.Context, tx *gorm.DB, cond map[string]interface{}) (int64, error) {
	var count int64
	if err := tx.Model(&model.Task{}).Where(cond).Count(&count).Error; err != nil {
		logger.ServiceLogger.WithContext(ctx).Errorf("task count with transaction error: %+v", err)
		return 0, err
	}
	return count, nil
}
