package output

import (
	"iflow-lite/type/model"
)

type (
	TaskQueryOutput struct {
		Total int64         `json:"total"`
		Items []*model.Task `json:"items"`
	}
)
