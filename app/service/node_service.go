package service

import (
	"context"
	"errors"

	"iflow-lite/core/util"
	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultNodeService = NewNodeService()

type NodeService struct {
}

func NewNodeService() *NodeService {
	return &NodeService{}
}

func (this *NodeService) NodeGet(ctx context.Context, in *input.NodeGetInput) (interface{}, error) {
	return dao.DefaultNodeDao.NodeGet(ctx, in.ID)
}

func (this *NodeService) NodeAdd(ctx context.Context, in *input.NodeAddInput) (interface{}, error) {
	process, err := dao.DefaultProcessDao.ProcessGet(ctx, in.ProcessID)
	if err != nil {
		return nil, err
	}
	if process == nil {
		return nil, errors.New("process not found")
	}
	m := &model.Node{
		ProcessID:   process.ID,
		ProcessCode: process.Code,
		Tag:         in.Tag,
		Name:        in.Name,
		Code:        in.Code,
		Type:        in.Type,
		Description: in.Description,
		CreatedBy:   util.UIDWithContext(ctx),
		UpdatedBy:   util.UIDWithContext(ctx),
	}
	if _, err := dao.DefaultNodeDao.NodeAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}
func (this *NodeService) NodeList(ctx context.Context, in *input.NodeListInput) (interface{}, error) {
	list := make([]*dto.Node, 0)
	cond := map[string]interface{}{}
	if in.ProcessID > 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.Code != "" {
		cond["code"] = in.Code
	}
	if in.Type != "" {
		cond["type"] = in.Type
	}

	items, err := dao.DefaultNodeDao.NodeList(ctx, cond)
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
			node := &dto.Node{Node: item}
			if user, ok := userMap[item.CreatedBy]; ok {
				node.CreatedByName = user.Name
			}
			if user, ok := userMap[item.UpdatedBy]; ok {
				node.UpdatedByName = user.Name
			}
			list = append(list, node)
		}
	}
	return list, nil
}
func (this *NodeService) NodeQuery(ctx context.Context, in *input.NodeQueryInput) (interface{}, error) {
	result := &output.NodeQueryOutput{Items: make([]*dto.Node, 0), Total: 0}
	cond := map[string]interface{}{}
	if in.ProcessID > 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.Code != "" {
		cond["code"] = in.Code
	}
	if in.Type != "" {
		cond["type"] = in.Type
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultNodeDao.NodeQuery(ctx, cond, in.Page, in.Size)
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
			node := &dto.Node{Node: item}
			if user, ok := userMap[item.CreatedBy]; ok {
				node.CreatedByName = user.Name
			}
			if user, ok := userMap[item.UpdatedBy]; ok {
				node.UpdatedByName = user.Name
			}
			result.Items = append(result.Items, node)
		}
	}
	result.Total = total
	return result, nil
}
