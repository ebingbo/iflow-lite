package input

type (
	AssignmentGetInput struct {
		ID uint64 `form:"id"`
	}
	AssignmentDeleteInput struct {
		ID uint64 `uri:"id"`
	}
	AssignmentAddInput struct {
		ProcessID uint64 `json:"process_id"`
		NodeID    uint64 `json:"node_id"`
		Type      string `json:"type"`
		Value     string `json:"value"`
		Priority  int    `json:"priority"`
		Strategy  string `json:"strategy"`
	}
	AssignmentQueryInput struct {
		Page        int    `json:"page" form:"page" binding:"required" example:"1"`  // 页码，从1开始
		Size        int    `json:"size" form:"size" binding:"required" example:"10"` // 每页数量
		ProcessID   uint64 `json:"process_id"`
		ProcessCode string `json:"process_code"`
		NodeID      uint64 `json:"node_id"`
		NodeCode    string `json:"node_code"`
		Type        string `json:"type"`
	}
	AssignmentUpdateInput struct {
		ID       uint64 `json:"id"`
		Type     string `json:"type"`
		Value    string `json:"value"`
		Priority int    `json:"priority"`
		Strategy string `json:"strategy"`
	}
	AssignmentListInput struct {
		ProcessID uint64 `json:"process_id" binding:"required"`
		NodeID    uint64 `json:"node_id"`
	}
)
