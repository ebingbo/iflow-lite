package dto

import "iflow-lite/type/model"

type (
	Task struct {
		*model.Task
		AssigneeName string `json:"assignee_name"`
	}
)
