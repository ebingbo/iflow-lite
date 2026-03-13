package output

import "iflow-lite/type/model"

type (
	RoleQueryOutput struct {
		Total int64         `json:"total"`
		Items []*model.Role `json:"items"`
	}
)
