package service

import (
	"context"
	"time"

	"iflow-lite/core/code"
	"iflow-lite/core/constant"
	"iflow-lite/dao"
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
	now := time.Now()
	m := &model.Task{
		ProcessID:   in.ProcessID,
		ProcessCode: process.Code,
		ProcessName: process.Name,
		NodeID:      in.NodeID,
		NodeCode:    node.Code,
		NodeName:    node.Name,
		ExecutionID: in.ExecutionID,
		AssigneeID:  in.AssigneeID,
		Description: in.Description,
		Remark:      in.Remark,
		Status:      constant.TaskStatusRunning,
		StartedAt:   &now,
	}
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
	if in.AssigneeID != "" {
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

	result.Total = total
	result.Items = items
	return result, nil
}
