package service

import (
	"context"

	"iflow-lite/core/util"
	"iflow-lite/dao"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
)

var (
	processDao = dao.NewProcessDao()
)

type ProcessService struct {
}

func NewProcessService() *ProcessService {
	return &ProcessService{}
}

func (this *ProcessService) ProcessGet(ctx context.Context, in *input.ProcessGetInput) (interface{}, error) {
	return processDao.ProcessGet(ctx, in.ID)
}

func (this *ProcessService) ProcessAdd(ctx context.Context, in *input.ProcessAddInput) (interface{}, error) {
	m := &model.Process{
		Name:        in.Name,
		Code:        in.Code,
		Description: in.Description,
		CreatedBy:   util.UIDWithContext(ctx),
		UpdatedBy:   util.UIDWithContext(ctx),
	}
	if _, err := processDao.ProcessAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
