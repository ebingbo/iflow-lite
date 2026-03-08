package model

import "time"

type AssignmentBuilder struct {
	item *Assignment
}

func NewAssignmentBuilder() *AssignmentBuilder {
	return &AssignmentBuilder{
		item: &Assignment{},
	}
}

func (builder *AssignmentBuilder) ID(id uint64) *AssignmentBuilder {
	builder.item.ID = id
	return builder
}

func (builder *AssignmentBuilder) Build() *Assignment {
	return builder.item
}

func (builder *AssignmentBuilder) ProcessID(processID uint64) *AssignmentBuilder {
	builder.item.ProcessID = processID
	return builder
}

func (builder *AssignmentBuilder) ProcessCode(processCode string) *AssignmentBuilder {
	builder.item.ProcessCode = processCode
	return builder
}

func (builder *AssignmentBuilder) NodeID(nodeID uint64) *AssignmentBuilder {
	builder.item.NodeID = nodeID
	return builder
}

func (builder *AssignmentBuilder) NodeCode(nodeCode string) *AssignmentBuilder {
	builder.item.NodeCode = nodeCode
	return builder
}

func (builder *AssignmentBuilder) PrincipalType(principalType string) *AssignmentBuilder {
	builder.item.PrincipalType = principalType
	return builder
}

func (builder *AssignmentBuilder) PrincipalID(principalID uint64) *AssignmentBuilder {
	builder.item.PrincipalID = principalID
	return builder
}

func (builder *AssignmentBuilder) Priority(priority int) *AssignmentBuilder {
	builder.item.Priority = priority
	return builder
}

func (builder *AssignmentBuilder) Strategy(strategy string) *AssignmentBuilder {
	builder.item.Strategy = strategy
	return builder
}

func (builder *AssignmentBuilder) CreatedAt(createdAt time.Time) *AssignmentBuilder {
	builder.item.CreatedAt = createdAt
	return builder
}

func (builder *AssignmentBuilder) UpdatedAt(updatedAt time.Time) *AssignmentBuilder {
	builder.item.UpdatedAt = updatedAt
	return builder
}
