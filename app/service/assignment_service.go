package service

import (
	"context"
	"errors"

	"iflow-lite/dao"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultAssignmentService = NewAssignmentService()

type AssignmentService struct {
}

func NewAssignmentService() *AssignmentService {
	return &AssignmentService{}
}

func (this *AssignmentService) AssignmentGet(ctx context.Context, in *input.AssignmentGetInput) (interface{}, error) {
	return dao.DefaultAssignmentDao.AssignmentGet(ctx, in.ID)
}
func (this *AssignmentService) AssignmentDelete(ctx context.Context, in *input.AssignmentDeleteInput) (interface{}, error) {
	if err := dao.DefaultAssignmentDao.AssignmentDelete(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}
func (this *AssignmentService) AssignmentUpdate(ctx context.Context, in *input.AssignmentUpdateInput) (interface{}, error) {
	m, err := dao.DefaultAssignmentDao.AssignmentGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	m.PrincipalType = in.PrincipalType
	m.PrincipalID = in.PrincipalID
	m.Priority = in.Priority
	m.Strategy = in.Strategy
	if err := dao.DefaultAssignmentDao.AssignmentUpdate(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
func (this *AssignmentService) AssignmentAdd(ctx context.Context, in *input.AssignmentAddInput) (interface{}, error) {
	process, err := dao.DefaultProcessDao.ProcessGet(ctx, in.ProcessID)
	if err != nil {
		return nil, err
	}
	node, err := dao.DefaultNodeDao.NodeGet(ctx, in.NodeID)
	if err != nil {
		return nil, err
	}
	if process == nil {
		return nil, errors.New("process not found")
	}
	if node == nil {
		return nil, errors.New("node not found")
	}
	m := model.NewAssignmentBuilder().
		ProcessID(in.ProcessID).
		ProcessCode(process.Code).
		NodeID(in.NodeID).
		NodeCode(node.Code).
		PrincipalType(in.PrincipalType).
		PrincipalID(in.PrincipalID).
		Priority(in.Priority).
		Strategy(in.Strategy).
		Build()
	if _, err := dao.DefaultAssignmentDao.AssignmentAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *AssignmentService) AssignmentQuery(ctx context.Context, in *input.AssignmentQueryInput) (interface{}, error) {
	result := new(output.AssignmentQueryOutput)
	cond := map[string]interface{}{}
	if in.ProcessID != 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.NodeID != 0 {
		cond["node_id"] = in.NodeID
	}
	if in.NodeCode != "" {
		cond["node_code"] = in.NodeCode
	}
	if in.PrincipalType != "" {
		cond["principal_type"] = in.PrincipalType
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultAssignmentDao.AssignmentQuery(ctx, cond, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	result.Items = items
	result.Total = total
	return result, nil
}

func (this *AssignmentService) AssignmentList(ctx context.Context, in *input.AssignmentListInput) (interface{}, error) {
	cond := map[string]interface{}{}
	if in.ProcessID != 0 {
		cond["process_id"] = in.ProcessID
	}

	if in.NodeID != 0 {
		cond["node_id"] = in.NodeID
	}
	return dao.DefaultAssignmentDao.AssignmentList(ctx, cond)
}
