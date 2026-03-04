package model

import "time"

type TransitionBuilder struct {
	item *Transition
}

func NewTransitionBuilder() *TransitionBuilder {
	return &TransitionBuilder{
		item: &Transition{},
	}
}

func (builder *TransitionBuilder) ID(id uint64) *TransitionBuilder {
	builder.item.ID = id
	return builder
}

func (builder *TransitionBuilder) ProcessID(processID uint64) *TransitionBuilder {
	builder.item.ProcessID = processID
	return builder
}

func (builder *TransitionBuilder) FromNodeID(fromNodeID uint64) *TransitionBuilder {
	builder.item.FromNodeID = fromNodeID
	return builder
}

func (builder *TransitionBuilder) ToNodeID(toNodeID uint64) *TransitionBuilder {
	builder.item.ToNodeID = toNodeID
	return builder
}

func (builder *TransitionBuilder) CreatedAt(t time.Time) *TransitionBuilder {
	builder.item.CreatedAt = t
	return builder
}
func (builder *TransitionBuilder) Build() *Transition {
	return builder.item
}

func (builder *TransitionBuilder) UpdatedAt(t time.Time) *TransitionBuilder {
	builder.item.UpdatedAt = t
	return builder
}
