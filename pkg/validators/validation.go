package validators

import (
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	custom_tag "github.com/peammapake/vincent-inventory-api/pkg/validators/custom"
)

var validate *validator.Validate
var m sync.Mutex

func InitValidator() {
	if validate == nil {
		m.Lock()
		defer m.Unlock()

		validate = validator.New()

		// Register custom validation
		custom_tag.RegisterCustomValidator(validate)

		// Turns to use json tag instead of struct field name
		// When using validationError.Field() it will return the json tag instead of struct field name
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})
	}
}

func GetValidator() *validator.Validate {
	if validate == nil {
		InitValidator()
	}
	return validate
}

func ValidateStruct(s interface{}) validator.ValidationErrors {
	v := GetValidator()
	if err := v.Struct(s); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			return validationErrs
		}
	}

	return nil
}
