package input

type (
	LogGetInput struct {
		ID uint64 `form:"id"`
	}

	LogAddInput struct {
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		ExecutionID uint64 `json:"execution_id"`
		NodeID      uint64 `json:"node_id"`
		NodeCode    string `json:"node_code"`
		TaskID      uint64 `json:"task_id"`
		Action      string `json:"action"`
		AssigneeID  string `json:"assignee_id"`
		Remark      string `json:"remark"`
	}

	LogQueryInput struct {
		Page        int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size        int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		ExecutionID uint64 `json:"execution_id"`
		NodeID      uint64 `json:"node_id"`
		NodeCode    string `json:"node_code"`
		TaskID      uint64 `json:"task_id"`
		AssigneeID  string `json:"assignee_id"`
	}
)
