package dto

type (
	UserOption struct {
		ID    uint64 `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	RoleOption struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}
)
