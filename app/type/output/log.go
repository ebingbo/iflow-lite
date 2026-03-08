package output

import "iflow-lite/type/dto"

type (
	LogQueryOutput struct {
		Total int64      `json:"total"`
		Items []*dto.Log `json:"items"`
	}
)
