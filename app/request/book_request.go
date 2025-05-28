package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateBookRequest struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	PublishedYear int    `json:"published_year" validate:"required,min=1000,max=2100"`
}

var validate = validator.New()

func (r *CreateBookRequest) Validate() error {
	err := validate.Struct(r)
	if err == nil {
		return nil
	}

	if ve, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fe := range ve {
			field := fe.Field()
			switch field {
			case "Title":
				errors["title"] = "Title is required"
			case "Author":
				errors["author"] = "Author is required"
			case "PublishedYear":
				errors["published_year"] = "Published year must be between 1000 and 2100"
			default:
				errors[fe.Field()] = fmt.Sprintf("Validation failed on %s", fe.Tag())
			}
		}
		return &ValidationError{Errors: errors}
	}

	return err
}

// Custom error type
type ValidationError struct {
	Errors map[string]string
}

func (v *ValidationError) Error() string {
	return "validation error"
}