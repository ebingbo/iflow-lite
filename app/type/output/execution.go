package output

import "iflow-lite/type/model"

type (
	ExecutionQueryOutput struct {
		Total int64              `json:"total"`
		Items []*model.Execution `json:"items"`
	}
)
