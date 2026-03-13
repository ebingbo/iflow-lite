package service

import (
	"context"
	"time"

	"iflow-lite/core/code"
	"iflow-lite/core/constant"
	"iflow-lite/core/util"
	"iflow-lite/dao"
	"iflow-lite/engine"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultTaskService = NewTaskService()

type TaskService struct {
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (this *TaskService) TaskGet(ctx context.Context, in *input.TaskGetInput) (interface{}, error) {
	return dao.DefaultTaskDao.TaskGet(ctx, in.ID)
}

func (this *TaskService) TaskAdd(ctx context.Context, in *input.TaskAddInput) (interface{}, error) {
	process, err := dao.DefaultProcessDao.ProcessGet(ctx, in.ProcessID)
	if err != nil {
		return nil, err
	}
	node, err := dao.DefaultNodeDao.NodeGet(ctx, in.NodeID)
	if err != nil {
		return nil, err
	}
	if process == nil {
		return nil, code.New(400, "process not found")
	}
	if node == nil {
		return nil, code.New(400, "node not found")
	}
	m := model.NewTaskBuilder().
		ProcessID(in.ProcessID).
		ProcessCode(process.Code).
		ProcessName(process.Name).
		NodeID(in.NodeID).
		NodeCode(node.Code).
		NodeName(node.Name).
		ExecutionID(in.ExecutionID).
		AssigneeID(in.AssigneeID).
		Remark(in.Remark).
		Status(constant.TaskStatusRunning).
		StartedAt(time.Now()).
		Build()
	if _, err := dao.DefaultTaskDao.TaskAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *TaskService) TaskQuery(ctx context.Context, in *input.TaskQueryInput) (interface{}, error) {
	result := new(output.TaskQueryOutput)
	cond := map[string]interface{}{}
	if in.ProcessID > 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.NodeID > 0 {
		cond["node_id"] = in.NodeID
	}
	if in.NodeCode != "" {
		cond["node_code"] = in.NodeCode
	}
	if in.ExecutionID > 0 {
		cond["execution_id"] = in.ExecutionID
	}
	if in.AssigneeID > 0 {
		cond["assignee_id"] = in.AssigneeID
	}
	if in.Status != "" {
		cond["status"] = in.Status
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultTaskDao.TaskQuery(ctx, cond, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, 0, len(items))
	seen := make(map[uint64]struct{}, len(items))
	for _, item := range items {
		if item.AssigneeID == 0 {
			continue
		}
		if _, ok := seen[item.AssigneeID]; ok {
			continue
		}
		seen[item.AssigneeID] = struct{}{}
		ids = append(ids, item.AssigneeID)
	}

	userMap := map[uint64]*model.User{}
	if len(ids) > 0 {
		users, userErr := dao.DefaultUserDao.UserListByIDs(ctx, ids)
		if userErr != nil {
			return nil, userErr
		}
		userMap = make(map[uint64]*model.User, len(users))
		for _, user := range users {
			userMap[user.ID] = user
		}
	}

	taskItems := make([]*dto.Task, 0, len(items))
	for _, item := range items {
		taskItem := &dto.Task{Task: item}
		if item.AssigneeID > 0 {
			if user, ok := userMap[item.AssigneeID]; ok {
				taskItem.AssigneeName = user.Name
			}
		}
		taskItems = append(taskItems, taskItem)
	}

	result.Total = total
	result.Items = taskItems
	return result, nil
}

func (this *TaskService) TaskClaimableQuery(ctx context.Context, in *input.TaskClaimableQueryInput) (interface{}, error) {
	result := new(output.TaskQueryOutput)
	uid := util.UIDWithContext(ctx)
	if uid == 0 {
		return nil, code.New(401, "unauthorized")
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultTaskDao.TaskClaimableQueryByUser(ctx, uid, in.Status, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, 0, len(items))
	seen := make(map[uint64]struct{}, len(items))
	for _, item := range items {
		if item.AssigneeID == 0 {
			continue
		}
		if _, ok := seen[item.AssigneeID]; ok {
			continue
		}
		seen[item.AssigneeID] = struct{}{}
		ids = append(ids, item.AssigneeID)
	}

	userMap := map[uint64]*model.User{}
	if len(ids) > 0 {
		users, userErr := dao.DefaultUserDao.UserListByIDs(ctx, ids)
		if userErr != nil {
			return nil, userErr
		}
		userMap = make(map[uint64]*model.User, len(users))
		for _, user := range users {
			userMap[user.ID] = user
		}
	}

	taskItems := make([]*dto.Task, 0, len(items))
	for _, item := range items {
		taskItem := &dto.Task{Task: item}
		if item.AssigneeID > 0 {
			if user, ok := userMap[item.AssigneeID]; ok {
				taskItem.AssigneeName = user.Name
			}
		}
		taskItems = append(taskItems, taskItem)
	}

	result.Total = total
	result.Items = taskItems
	return result, nil
}

func (this *TaskService) TaskComplete(ctx context.Context, in *input.TaskCompleteInput) error {
	return engine.DefaultExecutionEngine.TaskComplete(ctx, in.ID, in.AssigneeID, in.Remark)
}

func (this *TaskService) TaskClaim(ctx context.Context, in *input.TaskClaimInput) error {
	uid := util.UIDWithContext(ctx)
	if uid == 0 {
		return code.New(401, "unauthorized")
	}
	return engine.DefaultExecutionEngine.TaskClaim(ctx, in.ID, uid)
}

func (this *TaskService) TaskCandidateList(ctx context.Context, in *input.TaskCandidateListInput) (interface{}, error) {
	uid := util.UIDWithContext(ctx)
	if uid == 0 {
		return nil, code.New(401, "unauthorized")
	}
	task, err := dao.DefaultTaskDao.TaskGet(ctx, in.TaskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, code.New(404, "task not found")
	}
	if task.AssigneeID != uid {
		exists, existsErr := dao.DefaultTaskCandidateDao.TaskCandidateExists(ctx, task.ID, uid)
		if existsErr != nil {
			return nil, existsErr
		}
		if !exists {
			return nil, code.New(403, "forbidden")
		}
	}

	items, err := dao.DefaultTaskCandidateDao.TaskCandidateListByTaskID(ctx, in.TaskID)
	if err != nil {
		return nil, err
	}
	ids := make([]uint64, 0, len(items))
	for _, item := range items {
		ids = append(ids, item.UserID)
	}
	users, err := dao.DefaultUserDao.UserListByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	userMap := make(map[uint64]*model.User, len(users))
	for _, user := range users {
		userMap[user.ID] = user
	}
	result := make([]*dto.TaskCandidate, 0, len(items))
	for _, item := range items {
		taskCandidate := &dto.TaskCandidate{TaskCandidate: item}
		if user, ok := userMap[item.UserID]; ok {
			taskCandidate.UserName = user.Name
		}
		result = append(result, taskCandidate)
	}
	return result, nil
}

func (this *TaskService) TaskSkip(ctx context.Context, in *input.TaskSkipInput) error {
	return engine.DefaultExecutionEngine.TaskSkip(ctx, in.ID, in.AssigneeID)
}
func (this *TaskService) TaskDelegate(ctx context.Context, in *input.TaskDelegateInput) error {
	return engine.DefaultExecutionEngine.TaskDelegate(ctx, in.ID, in.FromAssigneeID, in.ToAssigneeID)
}
