package service

import (
	"context"
	"strings"

	"iflow-lite/core/code"
	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
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

func (this *RoleService) RoleAdd(ctx context.Context, in *input.RoleAddInput) (interface{}, error) {
	name := strings.TrimSpace(in.Name)
	codeValue := strings.TrimSpace(in.Code)
	if name == "" || codeValue == "" {
		return nil, code.New(400, "name and code are required")
	}
	role := &model.Role{
		Name: name,
		Code: codeValue,
	}
	if err := dao.DefaultRoleDao.RoleAdd(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (this *RoleService) RoleUpdate(ctx context.Context, in *input.RoleUpdateInput) (interface{}, error) {
	role, err := dao.DefaultRoleDao.RoleGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, code.New(404, "role not found")
	}
	name := strings.TrimSpace(in.Name)
	codeValue := strings.TrimSpace(in.Code)
	if name == "" || codeValue == "" {
		return nil, code.New(400, "name and code are required")
	}
	role.Name = name
	role.Code = codeValue
	if err := dao.DefaultRoleDao.RoleUpdate(ctx, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (this *RoleService) RoleDelete(ctx context.Context, in *input.RoleDeleteInput) (interface{}, error) {
	if err := dao.DefaultRoleDao.RoleDelete(ctx, in.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *RoleService) RoleQuery(ctx context.Context, in *input.RoleQueryInput) (interface{}, error) {
	result := new(output.RoleQueryOutput)
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultRoleDao.RoleQuery(ctx, in.Keyword, in.Page, in.Size)
	if err != nil {
		return nil, err
	}
	result.Total = total
	result.Items = items
	return result, nil
}
