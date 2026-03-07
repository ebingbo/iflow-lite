package input

type (
	NodeGetInput struct {
		ID uint64 `form:"id"`
	}
	NodeAddInput struct {
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		Tag         string `json:"tag"`
		Code        string `json:"code"`
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
	}
	NodeUpdateInput struct {
		ID          uint64 `json:"id"`
		Tag         string `json:"tag"`
		Description string `json:"description"`
	}
	NodeListInput struct {
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		Code        string `json:"code"`
		Type        string `json:"type"`
	}
	NodeQueryInput struct {
		Page        int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size        int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		Code        string `json:"code"`
		Type        string `json:"type"`
	}
	NodeDeleteInput struct {
		ID uint64 `uri:"id"`
	}
)
