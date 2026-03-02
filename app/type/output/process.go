package output

import (
	"iflow-lite/type/dto"
)

type (
	ProcessQueryOutput struct {
		Total int64          `json:"total"`
		Items []*dto.Process `json:"items"`
	}
)
