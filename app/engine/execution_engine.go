package engine

import (
	"context"
	"time"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/code"
	"iflow-lite/core/constant"
	"iflow-lite/dao"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

type ExecutionEngine struct {
}

func NewExecutionEngine() *ExecutionEngine {
	return &ExecutionEngine{}
}

func (e *ExecutionEngine) ExecutionStart(ctx context.Context, processCode string, businessKey string, businessType string, createdBy string) error {

	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 执行事务操作
		process, err := dao.DefaultProcessDao.ProcessTakeWithTransaction(ctx, tx, processCode)
		if err != nil {
			return err
		}
		if process == nil {
			return code.New(400, "process not found")
		}
		execution := model.NewExecutionBuilder().
			ProcessID(process.ID).
			ProcessCode(processCode).
			ProcessName(process.Name).
			BusinessKey(businessKey).
			BusinessType(businessType).
			CreatedBy(createdBy).
			Status(constant.ExecutionStatusRunning).
			Build()
		if _, err := dao.DefaultExecutionDao.ExecutionAddWithTransaction(ctx, tx, execution); err != nil {
			return err
		}

		node, err := dao.DefaultNodeDao.FirstNodeTakeWithTransaction(ctx, tx, process.ID)
		if err != nil {
			return err
		}
		if node == nil {
			return code.New(400, "first node not found")
		}
		return nil
	})
}

func (e *ExecutionEngine) ExecutionComplete(ctx context.Context, tx *gorm.DB, execution *model.Execution) error {
	if execution.Status == constant.ExecutionStatusCompleted {
		return nil
	}
	if !e.allTaskCompleted(ctx, tx, execution) {
		return nil
	}
	// 更新流程状态
	now := time.Now()
	execution.Status = constant.ExecutionStatusCompleted
	execution.Progress = 100
	execution.EndedAt = &now
	if err := dao.DefaultExecutionDao.ExecutionUpdateWithTransaction(ctx, tx, execution); err != nil {
		return err
	}
	log := model.NewLogBuilder().
		ProcessID(execution.ProcessID).
		ProcessCode(execution.ProcessCode).
		ExecutionID(execution.ID).
		Action(constant.LogActionCompleteProcess).
		Remark("完成流程").
		Build()
	if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
		return err
	}
	// todo 发送流程完成通知
	return nil
}

func (e *ExecutionEngine) TaskCreate(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	if !e.previewTaskCompleted(ctx, tx, execution, node) {
		return nil
	}
	if node.Type == constant.NodeTypeStart {
		return e.TaskCreateForStartNode(ctx, tx, execution, node)
	}
	if node.Type == constant.NodeTypeJoin {
		return e.TaskCreateForJoinNode(ctx, tx, execution, node)
	}
	if node.Type == constant.NodeTypeEnd {
		return e.TaskCreateForEndNode(ctx, tx, execution, node)
	}
	return nil
}

func (e *ExecutionEngine) TaskCreateForStartNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	log := model.NewLogBuilder().
		ProcessID(execution.ProcessID).
		ProcessCode(execution.ProcessCode).
		ExecutionID(execution.ID).
		NodeID(node.ID).
		NodeCode(node.Code).
		Action(constant.LogActionStartProcess).
		AssigneeID(execution.CreatedBy).
		Remark("启动流程").
		Build()
	if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
		return err
	}
	return e.NextTaskCreate(ctx, tx, execution, node)
}

func (e *ExecutionEngine) TaskCreateForJoinNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	log := model.NewLogBuilder().
		ProcessID(execution.ProcessID).
		ProcessCode(execution.ProcessCode).
		ExecutionID(execution.ID).
		NodeID(node.ID).
		NodeCode(node.Code).
		Action(constant.LogActionJoinTask).
		Remark("完成上游任务").
		Build()
	if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
		return err
	}
	return e.NextTaskCreate(ctx, tx, execution, node)
}

func (e *ExecutionEngine) TaskCreateForEndNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	return e.ExecutionComplete(ctx, tx, execution)
}

func (e *ExecutionEngine) NextTaskCreate(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	nextNodes, err := e.NextNodeList(ctx, tx, node)
	if err != nil {
		return err
	}
	for _, nextNode := range nextNodes {
		if err := e.TaskCreate(ctx, tx, execution, nextNode); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionEngine) NextNodeList(ctx context.Context, tx *gorm.DB, node *model.Node) ([]*model.Node, error) {
	toNodeIDs, err := dao.DefaultTransitionDao.ToNodeIDListWithTransaction(ctx, tx, node.ProcessID, node.ID)
	if err != nil {
		return nil, err
	}
	cond := map[string]interface{}{"id": toNodeIDs}
	return dao.DefaultNodeDao.NodeListWithTransaction(ctx, tx, cond)
}

func (e *ExecutionEngine) previewTaskCompleted(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) bool {
	fromNodeIDs, err := dao.DefaultTransitionDao.FromNodeIDListWithTransaction(ctx, tx, node.ProcessID, node.ID)
	if err != nil {
		return false
	}
	if len(fromNodeIDs) == 0 {
		return true
	}
	ok, err := dao.DefaultTaskDao.TaskCompleted(ctx, tx, execution.ID, fromNodeIDs)
	if err != nil {
		return false
	}
	return ok
}

func (e *ExecutionEngine) allTaskCompleted(ctx context.Context, tx *gorm.DB, execution *model.Execution) bool {
	nodeCond := map[string]interface{}{"process_id": execution.ProcessID, "type": constant.NodeTypeUserTask}
	nodeCount, err := dao.DefaultNodeDao.NodeCountWithTransaction(ctx, tx, nodeCond)
	if err != nil {
		return false
	}
	taskCond := map[string]interface{}{"execution_id": execution.ID, "status": constant.TaskStatusCompleted}
	taskCount, err := dao.DefaultTaskDao.TaskCountWithTransaction(ctx, tx, taskCond)
	if err != nil {
		return false
	}
	return nodeCount == taskCount
}
