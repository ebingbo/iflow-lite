package model

import "time"

type NodeBuilder struct {
	item *Node
}

func NewNodeBuilder() *NodeBuilder {
	return &NodeBuilder{
		item: &Node{},
	}
}

func (builder *NodeBuilder) ID(id uint64) *NodeBuilder {
	builder.item.ID = id
	return builder
}
func (builder *NodeBuilder) ProcessID(id uint64) *NodeBuilder {
	builder.item.ProcessID = id
	return builder
}
func (builder *NodeBuilder) ProcessCode(code string) *NodeBuilder {
	builder.item.ProcessCode = code
	return builder
}

func (builder *NodeBuilder) Tag(tag string) *NodeBuilder {
	builder.item.Tag = tag
	return builder
}

func (builder *NodeBuilder) Name(name string) *NodeBuilder {
	builder.item.Name = name
	return builder
}

func (builder *NodeBuilder) Build() *Node {
	return builder.item
}

func (builder *NodeBuilder) Code(code string) *NodeBuilder {
	builder.item.Code = code
	return builder
}

func (builder *NodeBuilder) Type(typeStr string) *NodeBuilder {
	builder.item.Type = typeStr
	return builder
}

func (builder *NodeBuilder) AssignMode(assignMode string) *NodeBuilder {
	builder.item.AssignMode = assignMode
	return builder
}

func (builder *NodeBuilder) X(x int) *NodeBuilder {
	builder.item.X = x
	return builder
}

func (builder *NodeBuilder) Y(y int) *NodeBuilder {
	builder.item.Y = y
	return builder
}

func (builder *NodeBuilder) Description(description string) *NodeBuilder {
	builder.item.Description = description
	return builder
}

func (builder *NodeBuilder) CreatedAt(t time.Time) *NodeBuilder {
	builder.item.CreatedAt = t
	return builder
}
func (builder *NodeBuilder) UpdatedAt(t time.Time) *NodeBuilder {
	builder.item.UpdatedAt = t
	return builder
}
func (builder *NodeBuilder) CreatedBy(createdBy uint64) *NodeBuilder {
	builder.item.CreatedBy = createdBy
	return builder
}

func (builder *NodeBuilder) UpdatedBy(updatedBy uint64) *NodeBuilder {
	builder.item.UpdatedBy = updatedBy
	return builder
}
