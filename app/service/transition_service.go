package service

import (
	"context"

	"iflow-lite/dao"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
)

var DefaultTransitionService = NewTransitionService()

type TransitionService struct {
}

func NewTransitionService() *TransitionService {
	return &TransitionService{}
}

func (this *TransitionService) TransitionGet(ctx context.Context, in *input.TransitionGetInput) (interface{}, error) {
	return dao.DefaultTransitionDao.TransitionGet(ctx, in.ID)
}
func (this *TransitionService) TransitionDelete(ctx context.Context, in *input.TransitionDeleteInput) (interface{}, error) {
	if err := dao.DefaultTransitionDao.TransitionDelete(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *TransitionService) TransitionUpdate(ctx context.Context, in *input.TransitionUpdateInput) (interface{}, error) {
	item, err := dao.DefaultTransitionDao.TransitionGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	item.FromNodeID = in.FromNodeID
	item.ToNodeID = in.ToNodeID
	if err := dao.DefaultTransitionDao.TransitionUpdate(ctx, item); err != nil {
		return nil, err
	}
	return nil, nil
}
func (this *TransitionService) TransitionAdd(ctx context.Context, in *input.TransitionAddInput) (interface{}, error) {
	m := model.NewTransitionBuilder().
		ProcessID(in.ProcessID).
		FromNodeID(in.FromNodeID).
		ToNodeID(in.ToNodeID).
		Build()
	if _, err := dao.DefaultTransitionDao.TransitionAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *TransitionService) TransitionList(ctx context.Context, in *input.TransitionListInput) (interface{}, error) {
	cond := map[string]interface{}{}
	if in.ProcessID != 0 {
		cond["process_id"] = in.ProcessID
	}

	items, err := dao.DefaultTransitionDao.TransitionList(ctx, cond)
	if err != nil {
		return nil, err
	}
	return items, nil
}
