package output

import "iflow-lite/type/dto"

type (
	UserQueryOutput struct {
		Total int64       `json:"total"`
		Items []*dto.User `json:"items"`
	}
)
