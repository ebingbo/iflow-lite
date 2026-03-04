package model

import "time"

type TaskBuilder struct {
	task *Task
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{
		task: &Task{},
	}
}

func (builder *TaskBuilder) ID(id uint64) *TaskBuilder {
	builder.task.ID = id
	return builder
}
func (builder *TaskBuilder) ProcessID(id uint64) *TaskBuilder {
	builder.task.ProcessID = id
	return builder
}
func (builder *TaskBuilder) ProcessCode(code string) *TaskBuilder {
	builder.task.ProcessCode = code
	return builder
}
func (builder *TaskBuilder) ProcessName(name string) *TaskBuilder {
	builder.task.ProcessName = name
	return builder
}
func (builder *TaskBuilder) ExecutionID(id uint64) *TaskBuilder {
	builder.task.ExecutionID = id
	return builder
}
func (builder *TaskBuilder) NodeID(id uint64) *TaskBuilder {
	builder.task.NodeID = id
	return builder
}
func (builder *TaskBuilder) NodeCode(code string) *TaskBuilder {
	builder.task.NodeCode = code
	return builder
}
func (builder *TaskBuilder) NodeName(name string) *TaskBuilder {
	builder.task.NodeName = name
	return builder
}
func (builder *TaskBuilder) AssigneeID(id string) *TaskBuilder {
	builder.task.AssigneeID = id
	return builder
}
func (builder *TaskBuilder) Status(status string) *TaskBuilder {
	builder.task.Status = status
	return builder
}
func (builder *TaskBuilder) StartedAt(t time.Time) *TaskBuilder {
	builder.task.StartedAt = &t
	return builder
}
func (builder *TaskBuilder) EndedAt(t time.Time) *TaskBuilder {
	builder.task.EndedAt = &t
	return builder
}

func (builder *TaskBuilder) Remark(remark string) *TaskBuilder {
	builder.task.Remark = remark
	return builder
}
func (builder *TaskBuilder) CreatedAt(t time.Time) *TaskBuilder {
	builder.task.CreatedAt = t
	return builder
}
func (builder *TaskBuilder) UpdatedAt(t time.Time) *TaskBuilder {
	builder.task.UpdatedAt = t
	return builder
}

func (builder *TaskBuilder) Build() *Task {
	return builder.task
}
