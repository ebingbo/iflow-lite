package dto

import "iflow-lite/type/model"

type Log struct {
	*model.Log
	AssigneeName string `json:"assignee_name"`
}
