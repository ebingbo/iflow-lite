package output

import (
	"iflow-lite/type/dto"
	"iflow-lite/type/model"
)

type (
	ProcessQueryOutput struct {
		Total int64          `json:"total"`
		Items []*dto.Process `json:"items"`
	}
	ProcessGetOutput struct {
		Process     *dto.Process        `json:"process"`
		Nodes       []*dto.Node         `json:"nodes"`
		Transitions []*model.Transition `json:"transitions"`
		Assignments []*model.Assignment `json:"assignments"`
	}
)
