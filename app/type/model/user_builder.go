package model

import "time"

type UserBuilder struct {
	item *User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		item: &User{},
	}
}

func (builder *UserBuilder) ID(id uint64) *UserBuilder {
	builder.item.ID = id
	return builder
}

func (builder *UserBuilder) Name(name string) *UserBuilder {
	builder.item.Name = name
	return builder
}

func (builder *UserBuilder) Email(email string) *UserBuilder {
	builder.item.Email = email
	return builder
}

func (builder *UserBuilder) Password(password string) *UserBuilder {
	builder.item.Password = password
	return builder
}

func (builder *UserBuilder) Status(status uint8) *UserBuilder {
	builder.item.Status = &status
	return builder
}

func (builder *UserBuilder) CreatedAt(t time.Time) *UserBuilder {
	builder.item.CreatedAt = t
	return builder
}
func (builder *UserBuilder) Build() *User {
	return builder.item
}

func (builder *UserBuilder) UpdatedAt(t time.Time) *UserBuilder {
	builder.item.UpdatedAt = t
	return builder
}
