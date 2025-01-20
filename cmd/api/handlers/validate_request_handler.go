package handlers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/RhoNit/budgetapi/cmd/api/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ValidateRequestBody(c echo.Context, payload interface{}) []*validation.ValidationError {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(payload)

	var errs []*validation.ValidationError

	validationErrs, ok := err.(validator.ValidationErrors)

	if ok {
		reflected := reflect.ValueOf(payload)

		for _, validationErr := range validationErrs {
			field, _ := reflected.Type().FieldByName(validationErr.StructField())

			key := field.Tag.Get("json")
			if key == "" {
				key = strings.ToLower(validationErr.StructField())
			}

			paramInt := validationErr.Param()

			condition := validationErr.Tag()

			keyToTitleCase := strings.Replace(key, "_", " ", -1)
			errMsg := keyToTitleCase + " field is " + condition

			switch condition {
			case "required":
				errMsg = keyToTitleCase + " is required"
			case "email":
				errMsg = keyToTitleCase + " must be a valid email address"
			case "min":
				if _, err := strconv.Atoi(paramInt); err == nil {
					errMsg = fmt.Sprintf("%s must be of %d characters long", keyToTitleCase, paramInt)
				}
			}

			currentValidationError := &validation.ValidationError{
				Key:       key,
				ErrorMsg:  errMsg,
				Condition: condition,
			}

			errs = append(errs, currentValidationError)
		}
	}

	return errs
}
