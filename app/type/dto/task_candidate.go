package dto

import "iflow-lite/type/model"

type TaskCandidate struct {
	*model.TaskCandidate
	UserName string `json:"user_name"`
}
