package configs

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func NewValidator() {
	Validate = validator.New()
}
