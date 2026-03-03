package model

import "time"

type LogBuilder struct {
	log *Log
}

func NewLogBuilder() *LogBuilder {
	return &LogBuilder{
		log: &Log{},
	}
}

func (builder *LogBuilder) ID(id uint64) *LogBuilder {
	builder.log.ID = id
	return builder
}
func (builder *LogBuilder) ProcessID(id uint64) *LogBuilder {
	builder.log.ProcessID = id
	return builder
}
func (builder *LogBuilder) ProcessCode(code string) *LogBuilder {
	builder.log.ProcessCode = code
	return builder
}
func (builder *LogBuilder) ExecutionID(id uint64) *LogBuilder {
	builder.log.ExecutionID = id
	return builder
}
func (builder *LogBuilder) NodeID(id uint64) *LogBuilder {
	builder.log.NodeID = id
	return builder
}
func (builder *LogBuilder) NodeCode(code string) *LogBuilder {
	builder.log.NodeCode = code
	return builder
}
func (builder *LogBuilder) TaskID(id uint64) *LogBuilder {
	builder.log.TaskID = id
	return builder
}
func (builder *LogBuilder) Action(action string) *LogBuilder {
	builder.log.Action = action
	return builder
}
func (builder *LogBuilder) AssigneeID(id string) *LogBuilder {
	builder.log.AssigneeID = id
	return builder
}
func (builder *LogBuilder) Remark(remark string) *LogBuilder {
	builder.log.Remark = remark
	return builder
}
func (builder *LogBuilder) CreatedAt(t time.Time) *LogBuilder {
	builder.log.CreatedAt = t
	return builder
}
func (builder *LogBuilder) Build() *Log {
	return builder.log
}
