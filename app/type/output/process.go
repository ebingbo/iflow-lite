package output

import "iflow-lite/type/model"

type (
	ProcessQueryOutput struct {
		Total int64           `json:"total"`
		Items []model.Process `json:"items"`
	}
)
