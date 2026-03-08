package input

type (
	TaskGetInput struct {
		ID uint64 `form:"id"`
	}

	TaskAddInput struct {
		ProcessID   uint64 `json:"process_id"`
		ExecutionID uint64 `json:"execution_id"`
		NodeID      uint64 `json:"node_id"`
		AssigneeID  uint64 `json:"assignee_id"`
		Description string `json:"description"`
		Remark      string `json:"remark"`
	}

	TaskQueryInput struct {
		Page        int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size        int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		ExecutionID uint64 `json:"execution_id"`
		NodeID      uint64 `json:"node_id"`
		NodeCode    string `json:"node_code"`
		AssigneeID  uint64 `json:"assignee_id"`
		Status      string `json:"status"`
	}
	TaskClaimableQueryInput struct {
		Page   int    `json:"page" form:"page" binding:"required" example:"1"`
		Size   int    `json:"size" form:"size" binding:"required" example:"10"`
		Status string `json:"status"`
	}
	TaskCandidateListInput struct {
		TaskID uint64 `json:"task_id" binding:"required"`
	}

	TaskCompleteInput struct {
		ID         uint64 `json:"id"`
		AssigneeID uint64 `json:"assignee_id"`
		Remark     string `json:"remark"`
	}
	TaskClaimInput struct {
		ID     uint64 `json:"id"`
		UserID uint64 `json:"user_id"`
	}
	TaskSkipInput struct {
		ID         uint64 `json:"id"`
		AssigneeID uint64 `json:"assignee_id"`
	}
	TaskDelegateInput struct {
		ID             uint64 `json:"id"`
		FromAssigneeID uint64 `json:"from_assignee_id"`
		ToAssigneeID   uint64 `json:"to_assignee_id"`
	}
)
