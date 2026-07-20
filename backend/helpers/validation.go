package helpers

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(s interface{}) map[string]string {
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
			case "required":
					errors[e.Field()] = e.Field() + " wajib diisi."
			case "min":
					errors[e.Field()] = e.Field() + " minimal " + e.Param() + " karakter."
			case "max":
					errors[e.Field()] = e.Field() + " maksimal " + e.Param() + " karakter."
			default:
					errors[e.Field()] = e.Field() + " tidak valid."
		}
	}

	return errors
}