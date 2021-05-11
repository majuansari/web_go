package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUser struct {
	UserName string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	City     string `json:"city" validate:"required"`
}

func (cu CreateUser) Validate() error {
	return validation.ValidateStruct(
		&cu,
		validation.Field(&cu.Name, validation.Required),
		validation.Field(&cu.UserName, validation.Required),
		validation.Field(&cu.Phone, validation.Required),
		validation.Field(&cu.City, validation.Required),
	)
}
