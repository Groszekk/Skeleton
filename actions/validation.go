package actions

import (
	"fmt"

	"github.com/go-playground/validator"
)

func (user *UserR) ValidateUser() (bool, string) {

	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(user)

	if user.Password != user.RePasword {
		return false, "passwords invalid"
	} else {
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				fmt.Println(err)
				return false, err.Error()
			}
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.StructField())
				return false, err.StructField() + " was wrong"
			}
			return false, err.Error()
		}
	}

	return true, "registration success"
}
