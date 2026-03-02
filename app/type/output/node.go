package output

import (
	"iflow-lite/type/dto"
)

type (
	NodeQueryOutput struct {
		Total int64       `json:"total"`
		Items []*dto.Node `json:"items"`
	}
)
