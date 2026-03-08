package service

import (
	"context"

	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
)

var DefaultRoleService = NewRoleService()

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (this *RoleService) RoleList(ctx context.Context, in *input.RoleListInput) (interface{}, error) {
	items, err := dao.DefaultRoleDao.RoleListForAssignment(ctx, in.Keyword, in.Size)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.RoleOption, 0, len(items))
	for _, item := range items {
		result = append(result, &dto.RoleOption{
			ID:   item.ID,
			Name: item.Name,
			Code: item.Code,
		})
	}
	return result, nil
}
