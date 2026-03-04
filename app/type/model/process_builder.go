package model

import "time"

type ProcessBuilder struct {
	item *Process
}

func NewProcessBuilder() *ProcessBuilder {
	return &ProcessBuilder{
		item: &Process{},
	}
}

func (builder *ProcessBuilder) ID(id uint64) *ProcessBuilder {
	builder.item.ID = id
	return builder
}

func (builder *ProcessBuilder) Name(name string) *ProcessBuilder {
	builder.item.Name = name
	return builder
}

func (builder *ProcessBuilder) Code(code string) *ProcessBuilder {
	builder.item.Code = code
	return builder
}

func (builder *ProcessBuilder) Description(description string) *ProcessBuilder {
	builder.item.Description = description
	return builder
}

func (builder *ProcessBuilder) Status(status uint8) *ProcessBuilder {
	builder.item.Status = &status
	return builder
}

func (builder *ProcessBuilder) CreatedAt(t time.Time) *ProcessBuilder {
	builder.item.CreatedAt = t
	return builder
}
func (builder *ProcessBuilder) UpdatedAt(t time.Time) *ProcessBuilder {
	builder.item.UpdatedAt = t
	return builder
}
func (builder *ProcessBuilder) CreatedBy(createdBy uint64) *ProcessBuilder {
	builder.item.CreatedBy = createdBy
	return builder
}
func (builder *ProcessBuilder) UpdatedBy(updatedBy uint64) *ProcessBuilder {
	builder.item.UpdatedBy = updatedBy
	return builder
}

func (builder *ProcessBuilder) Build() *Process {
	return builder.item
}
