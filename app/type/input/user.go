package input

type (
	UserGetInput struct {
		ID uint64 `form:"id"`
	}
	UserAddInput struct {
		Email    string `json:"email" form:"email" binding:"required" example:"bill@126.com"`
		Name     string `json:"name" form:"name" binding:"required" example:"bill"`
		Password string `json:"password" form:"password" binding:"required" example:"123456"`
	}
	UserQueryInput struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size   int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		Email  string `form:"email" json:"email" example:"bill@126.com"`
		Name   string `form:"name" json:"name" example:"bill"`
		Status *uint8 `form:"status" json:"status" example:"1"`
	}
	UserStatusUpdateInput struct {
		ID     uint64 `json:"id" binding:"required"`
		Status uint8  `json:"status" binding:"required"`
	}
	UserRoleListInput struct {
		UserID uint64 `json:"user_id" binding:"required"`
	}
	UserRoleUpdateInput struct {
		UserID  uint64   `json:"user_id" binding:"required"`
		RoleIDs []uint64 `json:"role_ids"`
	}
	UserLoginInput struct {
		Email    string `json:"email" form:"email" binding:"required" example:"bill@126.com"`
		Password string `json:"password" form:"password" binding:"required" example:"123456"`
	}
	UserListInput struct {
		Keyword string `json:"keyword" form:"keyword"`
		Size    int    `json:"size" form:"size"`
	}
	UserProfileUpdateInput struct {
		Name string `json:"name" binding:"required"`
	}
	UserPasswordUpdateInput struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	UserPasswordForgotInput struct {
		Email string `json:"email" binding:"required"`
	}
)
