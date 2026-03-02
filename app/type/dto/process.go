package dto

import "iflow-lite/type/model"

type (
	Process struct {
		*model.Process
		CreatedByName string `json:"created_by_name"`
		UpdatedByName string `json:"updated_by_name"`
	}
)
