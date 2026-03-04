package input

type (
	ExecutionGetInput struct {
		ID uint64 `form:"id"`
	}

	ExecutionAddInput struct {
		ProcessCode  string `json:"process_code"`
		BusinessKey  string `json:"business_key"`
		BusinessType string `json:"business_type"`
		CreatedBy    string `json:"created_by"`
	}

	ExecutionQueryInput struct {
		Page         int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size         int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		ProcessID    uint64 `json:"process_id"`
		ProcessCode  string `json:"process_code"`
		BusinessKey  string `json:"business_key"`
		BusinessType string `json:"business_type"`
		Status       string `json:"status"`
	}

	ExecutionStartInput struct {
		ProcessCode  string `json:"process_code"`
		BusinessKey  string `json:"business_key"`
		BusinessType string `json:"business_type"`
		CreatedBy    string `json:"created_by"`
	}
)
