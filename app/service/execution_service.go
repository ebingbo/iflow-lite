package service

import (
	"context"

	"iflow-lite/core/constant"
	"iflow-lite/dao"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultExecutionService = NewExecutionService()

type ExecutionService struct {
}

func NewExecutionService() *ExecutionService {
	return &ExecutionService{}
}

func (this *ExecutionService) ExecutionGet(ctx context.Context, in *input.ExecutionGetInput) (interface{}, error) {
	return dao.DefaultExecutionDao.ExecutionGet(ctx, in.ID)
}

func (this *ExecutionService) ExecutionAdd(ctx context.Context, in *input.ExecutionAddInput) (interface{}, error) {
	process, err := dao.DefaultProcessDao.ProcessTake(ctx, in.ProcessCode)
	if err != nil {
		return nil, err
	}
	m := &model.Execution{
		ProcessID:    process.ID,
		ProcessCode:  process.Code,
		ProcessName:  process.Name,
		BusinessKey:  in.BusinessKey,
		BusinessType: in.BusinessType,
		CreatedBy:    in.CreatedBy,
		Status:       constant.ExecutionStatusRunning,
	}
	return dao.DefaultExecutionDao.ExecutionAdd(ctx, m)
}

func (this *ExecutionService) ExecutionQuery(ctx context.Context, in *input.ExecutionQueryInput) (interface{}, error) {
	result := new(output.ExecutionQueryOutput)
	cond := map[string]interface{}{}
	if in.ProcessID > 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.BusinessKey != "" {
		cond["business_key"] = in.BusinessKey
	}
	if in.BusinessType != "" {
		cond["business_type"] = in.BusinessType
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
	items, total, err := dao.DefaultExecutionDao.ExecutionQuery(ctx, cond, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	result.Items = items
	result.Total = total
	return result, nil
}
