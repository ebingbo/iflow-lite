package output

import "iflow-lite/type/model"

type (
	AssignmentQueryOutput struct {
		Total int64               `json:"total"`
		Items []*model.Assignment `json:"items"`
	}
)
