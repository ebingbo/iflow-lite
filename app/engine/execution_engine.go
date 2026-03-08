package engine

import (
	"context"
	"strconv"
	"time"

	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/code"
	"iflow-lite/core/constant"
	"iflow-lite/dao"
	"iflow-lite/type/model"

	"gorm.io/gorm"
)

var DefaultExecutionEngine = NewExecutionEngine()

type ExecutionEngine struct {
}

type taskCandidatePrincipal struct {
	UserID     uint64
	SourceType string
	SourceID   uint64
}

func NewExecutionEngine() *ExecutionEngine {
	return &ExecutionEngine{}
}

func (e *ExecutionEngine) ExecutionStart(ctx context.Context, processCode string, businessKey string, businessType string, createdBy string) (*model.Execution, error) {
	var execution *model.Execution
	if err := client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 执行事务操作
		process, err := dao.DefaultProcessDao.ProcessTakeWithTransaction(ctx, tx, processCode)
		if err != nil {
			return err
		}
		if process == nil {
			return code.New(400, "process not found")
		}
		execution = model.NewExecutionBuilder().
			ProcessID(process.ID).
			ProcessCode(processCode).
			ProcessName(process.Name).
			BusinessKey(businessKey).
			BusinessType(businessType).
			CreatedBy(createdBy).
			StartedAt(time.Now()).
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
		return e.taskCreate(ctx, tx, execution, node)
	}); err != nil {
		return nil, err
	}
	return execution, nil
}

func (e *ExecutionEngine) TaskComplete(ctx context.Context, taskID uint64, assigneeID uint64, remark string) error {
	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		task, err := dao.DefaultTaskDao.TaskGetWithTransaction(ctx, tx, taskID)
		if err != nil {
			return err
		}
		if task == nil {
			return code.New(400, "task not found")
		}
		if task.Status == constant.TaskStatusCompleted {
			return nil
		}
		if task.AssigneeID != assigneeID {
			return code.New(400, "not allow complete task")
		}
		previewStatus := task.Status

		now := time.Now()
		task.Status = constant.TaskStatusCompleted
		task.EndedAt = &now
		task.Remark = remark
		if err := dao.DefaultTaskDao.TaskUpdateWithTransaction(ctx, tx, task); err != nil {
			return err
		}

		log := model.NewLogBuilder().
			ProcessID(task.ProcessID).
			ProcessCode(task.ProcessCode).
			ExecutionID(task.ExecutionID).
			NodeID(task.NodeID).
			NodeCode(task.NodeCode).
			TaskID(task.ID).
			Action(constant.LogActionCompleteTask).
			AssigneeID(strconv.FormatUint(assigneeID, 10)).
			Remark("完成任务").
			Build()
		if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
			return err
		}

		execution, err := dao.DefaultExecutionDao.ExecutionGetWithTransaction(ctx, tx, task.ExecutionID)
		if err != nil {
			return err
		}
		if execution == nil {
			return code.New(400, "execution not found")
		}

		if previewStatus == constant.TaskStatusSkipped {
			return e.executionComplete(ctx, tx, execution)
		}
		node, err := dao.DefaultNodeDao.NodeGetWithTransaction(ctx, tx, task.NodeID)
		if err != nil {
			return err
		}
		if node == nil {
			return code.New(400, "node not found")
		}
		if err := e.nextTaskCreate(ctx, tx, execution, node); err != nil {
			return err
		}
		return e.executionProgressUpdate(ctx, tx, execution)
	})
}

func (e *ExecutionEngine) TaskClaim(ctx context.Context, taskID uint64, userID uint64) error {
	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		task, err := dao.DefaultTaskDao.TaskGetWithTransaction(ctx, tx, taskID)
		if err != nil {
			return err
		}
		if task == nil {
			return code.New(400, "task not found")
		}
		if task.Status != constant.TaskStatusPending || task.AssigneeID != 0 {
			return code.New(400, "task already claimed")
		}
		exists, err := dao.DefaultTaskCandidateDao.TaskCandidateExistsWithTransaction(ctx, tx, taskID, userID)
		if err != nil {
			return err
		}
		if !exists {
			return code.New(400, "not candidate user")
		}

		rowsAffected, err := dao.DefaultTaskDao.TaskClaimWithTransaction(ctx, tx, taskID, userID)
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return code.New(400, "task already claimed")
		}

		log := model.NewLogBuilder().
			ProcessID(task.ProcessID).
			ProcessCode(task.ProcessCode).
			ExecutionID(task.ExecutionID).
			NodeID(task.NodeID).
			NodeCode(task.NodeCode).
			TaskID(task.ID).
			Action(constant.LogActionClaimTask).
			AssigneeID(strconv.FormatUint(userID, 10)).
			Remark("认领任务").
			Build()
		if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
			return err
		}
		return nil
	})
}

func (e *ExecutionEngine) TaskSkip(ctx context.Context, taskID uint64, assigneeID uint64) error {
	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		task, err := dao.DefaultTaskDao.TaskGetWithTransaction(ctx, tx, taskID)
		if err != nil {
			return err
		}
		if task == nil {
			return code.New(400, "task not found")
		}
		if task.Status == constant.TaskStatusSkipped {
			return nil
		}
		if task.AssigneeID != assigneeID {
			return code.New(400, "not allow skip task")
		}

		task.Status = constant.TaskStatusSkipped
		if err := dao.DefaultTaskDao.TaskUpdateWithTransaction(ctx, tx, task); err != nil {
			return err
		}
		log := model.NewLogBuilder().
			ProcessID(task.ProcessID).
			ProcessCode(task.ProcessCode).
			ExecutionID(task.ExecutionID).
			NodeID(task.NodeID).
			NodeCode(task.NodeCode).
			TaskID(task.ID).
			Action(constant.LogActionSkipTask).
			AssigneeID(strconv.FormatUint(assigneeID, 10)).
			Remark("跳过任务").
			Build()
		if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
			return err
		}

		execution, err := dao.DefaultExecutionDao.ExecutionGetWithTransaction(ctx, tx, task.ExecutionID)
		if err != nil {
			return err
		}
		if execution == nil {
			return code.New(400, "execution not found")
		}
		node, err := dao.DefaultNodeDao.NodeGetWithTransaction(ctx, tx, task.NodeID)
		if err != nil {
			return err
		}
		if node == nil {
			return code.New(400, "node not found")
		}
		return e.nextTaskCreate(ctx, tx, execution, node)
	})
}

func (e *ExecutionEngine) TaskDelegate(ctx context.Context, taskID uint64, fromAssigneeID uint64, toAssigneeID uint64) error {
	return client.MysqlDB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		task, err := dao.DefaultTaskDao.TaskGetWithTransaction(ctx, tx, taskID)
		if err != nil {
			return err
		}
		if task == nil {
			return code.New(400, "task not found")
		}
		if task.AssigneeID != fromAssigneeID {
			return code.New(400, "not allow delegate task")
		}
		task.AssigneeID = toAssigneeID
		if err := dao.DefaultTaskDao.TaskUpdateWithTransaction(ctx, tx, task); err != nil {
			return err
		}
		log := model.NewLogBuilder().
			ProcessID(task.ProcessID).
			ProcessCode(task.ProcessCode).
			ExecutionID(task.ExecutionID).
			NodeID(task.NodeID).
			NodeCode(task.NodeCode).
			TaskID(task.ID).
			Action(constant.LogActionDelegateTask).
			AssigneeID(strconv.FormatUint(fromAssigneeID, 10)).
			Remark("委托任务").
			Build()
		if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
			return err
		}
		return nil
	})
}
func (e *ExecutionEngine) executionComplete(ctx context.Context, tx *gorm.DB, execution *model.Execution) error {
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

func (e *ExecutionEngine) taskCreate(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	if !e.previewTaskCompleted(ctx, tx, execution, node) {
		return nil
	}
	if node.Type == constant.NodeTypeStart {
		return e.taskCreateForStartNode(ctx, tx, execution, node)
	}
	if node.Type == constant.NodeTypeJoin {
		return e.taskCreateForJoinNode(ctx, tx, execution, node)
	}
	if node.Type == constant.NodeTypeEnd {
		return e.taskCreateForEndNode(ctx, tx, execution, node)
	}
	if node.Type == constant.NodeTypeUserTask {
		return e.taskCreateForUserTaskNode(ctx, tx, execution, node)
	}
	return nil
}

func (e *ExecutionEngine) taskCreateForStartNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
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
	return e.nextTaskCreate(ctx, tx, execution, node)
}

func (e *ExecutionEngine) taskCreateForJoinNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
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
	return e.nextTaskCreate(ctx, tx, execution, node)
}

func (e *ExecutionEngine) taskCreateForEndNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	return e.executionComplete(ctx, tx, execution)
}

func (e *ExecutionEngine) taskCreateForUserTaskNode(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	principals, err := e.resolveTaskCandidatePrincipals(ctx, tx, execution, node)
	if err != nil {
		return err
	}
	if len(principals) == 0 {
		return code.New(400, "no assignee")
	}

	if node.AssignMode == constant.NodeAssignModeCandidate {
		task := model.NewTaskBuilder().
			ProcessID(execution.ProcessID).
			ProcessCode(execution.ProcessCode).
			ProcessName(execution.ProcessName).
			ExecutionID(execution.ID).
			NodeID(node.ID).
			NodeCode(node.Code).
			NodeName(node.Name).
			AssigneeID(0).
			Status(constant.TaskStatusPending).
			Build()
		if _, err := dao.DefaultTaskDao.TaskAddWithTransaction(ctx, tx, task); err != nil {
			return err
		}

		candidates := make([]*model.TaskCandidate, 0, len(principals))
		for _, principal := range principals {
			candidates = append(candidates, &model.TaskCandidate{
				TaskID:     task.ID,
				UserID:     principal.UserID,
				SourceType: principal.SourceType,
				SourceID:   principal.SourceID,
			})
		}
		if err := dao.DefaultTaskCandidateDao.TaskCandidateBatchAddWithTransaction(ctx, tx, candidates); err != nil {
			return err
		}

		log := model.NewLogBuilder().
			ProcessID(execution.ProcessID).
			ProcessCode(execution.ProcessCode).
			ExecutionID(execution.ID).
			NodeID(node.ID).
			NodeCode(node.Code).
			TaskID(task.ID).
			Action(constant.LogActionCreateTask).
			Remark("创建候选任务").
			Build()
		if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
			return err
		}
		return nil
	}

	assignee := principals[0].UserID
	task := model.NewTaskBuilder().
		ProcessID(execution.ProcessID).
		ProcessCode(execution.ProcessCode).
		ProcessName(execution.ProcessName).
		ExecutionID(execution.ID).
		NodeID(node.ID).
		NodeCode(node.Code).
		NodeName(node.Name).
		AssigneeID(assignee).
		Status(constant.TaskStatusRunning).
		StartedAt(time.Now()).
		Build()
	if _, err := dao.DefaultTaskDao.TaskAddWithTransaction(ctx, tx, task); err != nil {
		return err
	}
	log := model.NewLogBuilder().
		ProcessID(execution.ProcessID).
		ProcessCode(execution.ProcessCode).
		ExecutionID(execution.ID).
		NodeID(node.ID).
		NodeCode(node.Code).
		TaskID(task.ID).
		Action(constant.LogActionCreateTask).
		AssigneeID(strconv.FormatUint(assignee, 10)).
		Remark("创建任务").
		Build()
	if _, err := dao.DefaultLogDao.LogAddWithTransaction(ctx, tx, log); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionEngine) resolveTaskCandidatePrincipals(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) ([]taskCandidatePrincipal, error) {
	assignmentList, err := dao.DefaultAssignmentDao.AssignmentListWithTransaction(ctx, tx, map[string]interface{}{"node_id": node.ID})
	if err != nil {
		return nil, err
	}
	candidates := make([]taskCandidatePrincipal, 0)
	seenUser := make(map[uint64]struct{})
	for _, assignment := range assignmentList {
		if assignment.PrincipalType == constant.AssignmentTypeUser && assignment.PrincipalID > 0 {
			if _, ok := seenUser[assignment.PrincipalID]; ok {
				continue
			}
			seenUser[assignment.PrincipalID] = struct{}{}
			candidates = append(candidates, taskCandidatePrincipal{
				UserID:     assignment.PrincipalID,
				SourceType: constant.AssignmentTypeUser,
				SourceID:   assignment.PrincipalID,
			})
			continue
		}
		if assignment.PrincipalType == constant.AssignmentTypeRole && assignment.PrincipalID > 0 {
			userIDs, roleErr := dao.DefaultUserDao.UserIDListByRoleIDWithTransaction(ctx, tx, assignment.PrincipalID)
			if roleErr != nil {
				return nil, roleErr
			}
			for _, userID := range userIDs {
				if userID == 0 {
					continue
				}
				if _, ok := seenUser[userID]; ok {
					continue
				}
				seenUser[userID] = struct{}{}
				candidates = append(candidates, taskCandidatePrincipal{
					UserID:     userID,
					SourceType: constant.AssignmentTypeRole,
					SourceID:   assignment.PrincipalID,
				})
			}
		}
	}
	if len(candidates) == 0 {
		createdBy, parseErr := strconv.ParseUint(execution.CreatedBy, 10, 64)
		if parseErr == nil && createdBy > 0 {
			candidates = append(candidates, taskCandidatePrincipal{
				UserID:     createdBy,
				SourceType: constant.AssignmentTypeUser,
				SourceID:   createdBy,
			})
		}
	}
	return candidates, nil
}
func (e *ExecutionEngine) nextTaskCreate(ctx context.Context, tx *gorm.DB, execution *model.Execution, node *model.Node) error {
	nextNodes, err := e.nextNodeList(ctx, tx, node)
	if err != nil {
		return err
	}
	for _, nextNode := range nextNodes {
		if err := e.taskCreate(ctx, tx, execution, nextNode); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionEngine) nextNodeList(ctx context.Context, tx *gorm.DB, node *model.Node) ([]*model.Node, error) {
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

func (e *ExecutionEngine) executionProgressUpdate(ctx context.Context, tx *gorm.DB, execution *model.Execution) error {
	nodeCond := map[string]interface{}{"process_id": execution.ProcessID, "type": constant.NodeTypeUserTask}
	nodeCount, err := dao.DefaultNodeDao.NodeCountWithTransaction(ctx, tx, nodeCond)
	if err != nil {
		return err
	}
	taskCond := map[string]interface{}{"execution_id": execution.ID, "status": constant.TaskStatusCompleted}
	taskCount, err := dao.DefaultTaskDao.TaskCountWithTransaction(ctx, tx, taskCond)
	if err != nil {
		return err
	}
	execution.Progress = float64(taskCount) / float64(nodeCount) * 100
	if err := dao.DefaultExecutionDao.ExecutionUpdateWithTransaction(ctx, tx, execution); err != nil {
		return err
	}
	return nil
}
