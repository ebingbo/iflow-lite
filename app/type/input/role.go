package input

type (
	RoleListInput struct {
		Keyword string `json:"keyword" form:"keyword"`
		Size    int    `json:"size" form:"size"`
	}
)
