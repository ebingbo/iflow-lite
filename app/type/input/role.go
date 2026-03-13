package input

type (
	RoleListInput struct {
		Keyword string `json:"keyword" form:"keyword"`
		Size    int    `json:"size" form:"size"`
	}
	RoleAddInput struct {
		Name string `json:"name" binding:"required"`
		Code string `json:"code" binding:"required"`
	}
	RoleUpdateInput struct {
		ID   uint64 `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
		Code string `json:"code" binding:"required"`
	}
	RoleDeleteInput struct {
		ID uint64 `uri:"id" binding:"required"`
	}
	RoleQueryInput struct {
		Page    int    `json:"page" form:"page" binding:"required" example:"1"`
		Size    int    `json:"size" form:"size" binding:"required" example:"10"`
		Keyword string `json:"keyword" form:"keyword"`
	}
)
