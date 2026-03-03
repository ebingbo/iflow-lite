package model

import "time"

type ExecutionBuilder struct {
	execution *Execution
}

func NewExecutionBuilder() *ExecutionBuilder {
	return &ExecutionBuilder{
		execution: &Execution{},
	}
}
func (builder *ExecutionBuilder) ID(id uint64) *ExecutionBuilder {
	builder.execution.ID = id
	return builder
}
func (builder *ExecutionBuilder) ProcessID(id uint64) *ExecutionBuilder {
	builder.execution.ProcessID = id
	return builder
}

func (builder *ExecutionBuilder) ProcessCode(code string) *ExecutionBuilder {
	builder.execution.ProcessCode = code
	return builder
}
func (builder *ExecutionBuilder) ProcessName(name string) *ExecutionBuilder {
	builder.execution.ProcessName = name
	return builder
}
func (builder *ExecutionBuilder) BusinessKey(key string) *ExecutionBuilder {
	builder.execution.BusinessKey = key
	return builder
}
func (builder *ExecutionBuilder) BusinessType(typeStr string) *ExecutionBuilder {
	builder.execution.BusinessType = typeStr
	return builder
}
func (builder *ExecutionBuilder) Status(status string) *ExecutionBuilder {
	builder.execution.Status = status
	return builder
}
func (builder *ExecutionBuilder) Progress(progress float64) *ExecutionBuilder {
	builder.execution.Progress = progress
	return builder
}
func (builder *ExecutionBuilder) CreatedBy(createdBy string) *ExecutionBuilder {
	builder.execution.CreatedBy = createdBy
	return builder
}
func (builder *ExecutionBuilder) StartedAt(startedAt time.Time) *ExecutionBuilder {
	builder.execution.StartedAt = &startedAt
	return builder
}
func (builder *ExecutionBuilder) EndedAt(endedAt time.Time) *ExecutionBuilder {
	builder.execution.EndedAt = &endedAt
	return builder
}

func (builder *ExecutionBuilder) CreatedAt(createdAt time.Time) *ExecutionBuilder {
	builder.execution.CreatedAt = createdAt
	return builder
}
func (builder *ExecutionBuilder) UpdatedAt(updatedAt time.Time) *ExecutionBuilder {
	builder.execution.UpdatedAt = updatedAt
	return builder
}
func (builder *ExecutionBuilder) Build() *Execution {
	return builder.execution
}
