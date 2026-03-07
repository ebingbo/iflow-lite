package input

type (
	TransitionGetInput struct {
		ID uint64 `form:"id"`
	}
	TransitionDeleteInput struct {
		ID uint64 `uri:"id"`
	}
	TransitionAddInput struct {
		ProcessID  uint64 `json:"process_id"`
		FromNodeID uint64 `json:"from_node_id"`
		ToNodeID   uint64 `json:"to_node_id"`
	}
	TransitionUpdateInput struct {
		ID         uint64 `json:"id"`
		FromNodeID uint64 `json:"from_node_id"`
		ToNodeID   uint64 `json:"to_node_id"`
	}
	TransitionListInput struct {
		ProcessID uint64 `json:"process_id"`
	}
)
