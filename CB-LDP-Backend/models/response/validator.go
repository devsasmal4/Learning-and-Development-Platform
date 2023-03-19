package response

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidResponse interface {
	ValidateStruct() error
}

func (question QuestionJsonResponse) ValidateStruct() error {
	if validationErr := validate.Struct(&question); validationErr != nil {
		return validationErr
	}
	return nil
}
