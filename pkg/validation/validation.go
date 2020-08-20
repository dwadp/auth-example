package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type (
	Error struct {
		Key     string `json:"key"`
		Message string `json:"message"`
	}
)

func Format(errs validator.ValidationErrors, s interface{}) ([]Error, error) {
	var result []Error

	sType := reflect.TypeOf(s)

	if sType.Kind() != reflect.Struct {
		return nil, errors.New("[validation error]: source must be the type of struct")
	}

	for _, err := range errs {
		field, ok := sType.FieldByName(err.Field())

		if !ok {
			continue
		}

		json := field.Tag.Get("json")
		label := json

		if strings.Contains(json, "_") {
			label = strings.Title(strings.ReplaceAll(json, "_", " "))
		}

		switch err.ActualTag() {
		case "required":
			result = append(result, Error{
				Key:     json,
				Message: fmt.Sprintf("The %s field is required", label),
			})
		case "email":
			result = append(result, Error{
				Key:     json,
				Message: fmt.Sprintf("The %s field contains an invalid email format", label),
			})
		case "min":
			param := err.Param()

			result = append(result, Error{
				Key:     json,
				Message: fmt.Sprintf("The %s field must contain at least %s characters", label, param),
			})
		case "max":
			param := err.Param()

			result = append(result, Error{
				Key:     json,
				Message: fmt.Sprintf("The %s field cannot contain more than %s characters", label, param),
			})
		default:
			result = append(result, Error{
				Key:     json,
				Message: fmt.Sprintf("Cannot find formatter for '%s' on %s field", err.ActualTag(), json),
			})
		}
	}

	return result, nil
}
