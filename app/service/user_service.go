package service

import (
	"context"
	"fmt"
	"strings"

	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/core/code"
	"iflow-lite/core/token"
	"iflow-lite/core/util"
	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"

	"golang.org/x/crypto/bcrypt"
)

var DefaultUserService = NewUserService()

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) UserGet(ctx context.Context, in *input.UserGetInput) (interface{}, error) {
	return dao.DefaultUserDao.UserGet(ctx, in.ID)
}

func (this *UserService) UserProfile(ctx context.Context) (interface{}, error) {
	uid := util.UIDWithContext(ctx)
	return dao.DefaultUserDao.UserGet(ctx, uid)
}

func (this *UserService) UserAdd(ctx context.Context, in *input.UserAddInput) (interface{}, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ServiceLogger.Errorf("bcrypt generate from password error: %+v", err)
		return nil, err
	}

	m := model.NewUserBuilder().
		Name(in.Name).
		Email(in.Email).
		Password(string(pwd)).
		Build()
	if err := dao.DefaultUserDao.UserAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *UserService) UserQuery(ctx context.Context, in *input.UserQueryInput) (interface{}, error) {
	result := new(output.UserQueryOutput)
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultUserDao.UserQuery(ctx, in)
	if err != nil {
		return nil, err
	}
	result.Total = total
	result.Items = make([]*dto.User, 0, len(items))
	for _, item := range items {
		result.Items = append(result.Items, &dto.User{
			ID:        item.ID,
			Name:      item.Name,
			Email:     item.Email,
			Status:    item.Status,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return result, nil
}

func (this *UserService) UserStatusUpdate(ctx context.Context, in *input.UserStatusUpdateInput) (interface{}, error) {
	user, err := dao.DefaultUserDao.UserGet(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, code.New(404, "user not found")
	}
	status := in.Status
	if status != 0 && status != 1 {
		return nil, code.New(400, "invalid status")
	}
	user.Status = &status
	if err := dao.DefaultUserDao.UserUpdate(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (this *UserService) UserRoleList(ctx context.Context, in *input.UserRoleListInput) (interface{}, error) {
	return dao.DefaultUserRoleDao.UserRoleListByUserID(ctx, in.UserID)
}

func (this *UserService) UserRoleUpdate(ctx context.Context, in *input.UserRoleUpdateInput) (interface{}, error) {
	user, err := dao.DefaultUserDao.UserGet(ctx, in.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, code.New(404, "user not found")
	}
	if err := dao.DefaultUserRoleDao.UserRoleReplace(ctx, in.UserID, in.RoleIDs); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *UserService) UserLogin(ctx context.Context, in *input.UserLoginInput) (interface{}, error) {
	user, err := dao.DefaultUserDao.UserTake(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, err
	}
	jwt, err := token.GenerateJWT(fmt.Sprintf("%d", user.ID))
	if err != nil {
		logger.ServiceLogger.Errorf("generate jwt error: %+v", err)
		return nil, err
	}
	return map[string]interface{}{"token": jwt, "user": user}, err
}

func (this *UserService) UserList(ctx context.Context, in *input.UserListInput) (interface{}, error) {
	items, err := dao.DefaultUserDao.UserListForAssignment(ctx, in.Keyword, in.Size)
	if err != nil {
		return nil, err
	}
	result := make([]*dto.UserOption, 0, len(items))
	for _, item := range items {
		result = append(result, &dto.UserOption{
			ID:    item.ID,
			Name:  item.Name,
			Email: item.Email,
		})
	}
	return result, nil
}

func (this *UserService) UserProfileUpdate(ctx context.Context, in *input.UserProfileUpdateInput) (interface{}, error) {
	uid := util.UIDWithContext(ctx)
	if uid == 0 {
		return nil, code.New(401, "unauthorized")
	}
	user, err := dao.DefaultUserDao.UserGet(ctx, uid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, code.New(404, "user not found")
	}
	name := strings.TrimSpace(in.Name)
	if name == "" {
		return nil, code.New(400, "name is required")
	}
	user.Name = name
	if err := dao.DefaultUserDao.UserUpdate(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (this *UserService) UserPasswordUpdate(ctx context.Context, in *input.UserPasswordUpdateInput) (interface{}, error) {
	uid := util.UIDWithContext(ctx)
	if uid == 0 {
		return nil, code.New(401, "unauthorized")
	}
	user, err := dao.DefaultUserDao.UserGet(ctx, uid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, code.New(404, "user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return nil, code.New(400, "old password not correct")
	}
	newPassword := strings.TrimSpace(in.NewPassword)
	if len(newPassword) < 8 {
		return nil, code.New(400, "new password length must be at least 8")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.ServiceLogger.Errorf("bcrypt generate from password error: %+v", err)
		return nil, err
	}
	user.Password = string(hash)
	if err := dao.DefaultUserDao.UserUpdate(ctx, user); err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *UserService) UserPasswordForgot(ctx context.Context, in *input.UserPasswordForgotInput) (interface{}, error) {
	// 预留：后续接入邮件服务与重置令牌逻辑。
	_ = strings.TrimSpace(in.Email)
	return nil, nil
}
