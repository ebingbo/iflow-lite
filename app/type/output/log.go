package output

import "iflow-lite/type/model"

type (
	LogQueryOutput struct {
		Total int64        `json:"total"`
		Items []*model.Log `json:"items"`
	}
)
