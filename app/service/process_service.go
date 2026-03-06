package service

import (
	"context"

	"iflow-lite/core/util"
	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultProcessService = NewProcessService()

type ProcessService struct {
}

func NewProcessService() *ProcessService {
	return &ProcessService{}
}

func (this *ProcessService) ProcessGet(ctx context.Context, in *input.ProcessGetInput) (interface{}, error) {
	return dao.DefaultProcessDao.ProcessGet(ctx, in.ID)
}

func (this *ProcessService) ProcessDelete(ctx context.Context, in *input.ProcessDeleteInput) (interface{}, error) {
	if err := dao.DefaultProcessDao.ProcessDelete(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *ProcessService) ProcessDisable(ctx context.Context, in *input.ProcessDisableInput) (interface{}, error) {
	if err := dao.DefaultProcessDao.ProcessDisable(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *ProcessService) ProcessEnable(ctx context.Context, in *input.ProcessEnableInput) (interface{}, error) {
	if err := dao.DefaultProcessDao.ProcessEnable(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *ProcessService) ProcessTake(ctx context.Context, in *input.ProcessTakeInput) (interface{}, error) {
	return dao.DefaultProcessDao.ProcessTake(ctx, in.Code)
}
func (this *ProcessService) ProcessAdd(ctx context.Context, in *input.ProcessAddInput) (interface{}, error) {
	m := model.NewProcessBuilder().
		Name(in.Name).
		Code(in.Code).
		Description(in.Description).
		CreatedBy(util.UIDWithContext(ctx)).
		UpdatedBy(util.UIDWithContext(ctx)).
		Build()
	if _, err := dao.DefaultProcessDao.ProcessAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *ProcessService) ProcessUpdate(ctx context.Context, in *input.ProcessUpdateInput) (interface{}, error) {
	m, err := dao.DefaultProcessDao.ProcessGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	m.Name = in.Name
	m.Description = in.Description
	if err := dao.DefaultProcessDao.ProcessUpdate(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *ProcessService) ProcessQuery(ctx context.Context, in *input.ProcessQueryInput) (interface{}, error) {
	result := &output.ProcessQueryOutput{}
	cond := map[string]interface{}{}
	if in.Name != "" {
		cond["name"] = in.Name
	}
	if in.Code != "" {
		cond["code"] = in.Code
	}
	if in.Status != nil {
		cond["status"] = in.Status
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultProcessDao.ProcessQuery(ctx, cond, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		var ids []uint64
		for _, item := range items {
			ids = append(ids, item.CreatedBy, item.UpdatedBy)
		}
		users, err := dao.DefaultUserDao.UserListByIDs(ctx, ids)
		if err != nil {
			return nil, err
		}
		userMap := make(map[uint64]*model.User)
		for _, user := range users {
			userMap[user.ID] = user
		}
		for _, item := range items {
			process := &dto.Process{Process: item}
			if user, ok := userMap[item.CreatedBy]; ok {
				process.CreatedByName = user.Name
			}
			if user, ok := userMap[item.UpdatedBy]; ok {
				process.UpdatedByName = user.Name
			}
			result.Items = append(result.Items, process)
		}
	}
	result.Total = total
	return result, nil
}
