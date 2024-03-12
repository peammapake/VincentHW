package custom

import (
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/peammapake/vincent-inventory-api/pkg/validators/tag"
)

func RegisterCustomValidator(v *validator.Validate) {
	v.RegisterValidation(tag.X_UUID, uuid4)
	v.RegisterValidation(tag.X_UUID4_OR_EMPTY, uuid4OrEmpty)
	v.RegisterValidation(tag.X_URL_OR_EMPTY, urlOrEmpty)
}

/*
urlOrEmpty validates if the field is either an empty string or a valid URL.

	usage: `validate:"url_or_empty"`
	valid: "", "http://www.google.com", "https://www.google.com"
	invalid: "http://", "http://www.google.com/", "http://www.google.com/?q=go+validator"
*/
func urlOrEmpty(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	if fieldValue == "" {
		return true // Allow empty string
	}
	_, err := url.ParseRequestURI(fieldValue)
	return err == nil // Check if it's a valid URL
}

/*
uuid4 validates if the field is a valid UUIDv4.

	usage: `validate:"uuid4"`
	valid: "e4eaaaf2-d142-11e1-b3e4-080027620cdd"
	invalid: "e4eaaaf2-d142-11e1-b3e4-080027620cd", "http://www.google.com/"
*/
func uuid4(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	_, err := uuid.Parse(fieldValue)
	return err == nil // Check if it's a valid URL
}

/*
uuid4OrEmpty validates if the field is either an empty string or a valid UUIDv4.

	usage: `validate:"uuid4_or_empty"`
	valid: "", "e4eaaaf2-d142-11e1-b3e4-080027620cdd"
	invalid: "e4eaaaf2-d142-11e1-b3e4-080027620cd", "http://www.google.com/"
*/
func uuid4OrEmpty(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	if fieldValue == "" {
		return true // Allow empty string
	}
	_, err := uuid.Parse(fieldValue)
	return err == nil // Check if it's a valid UUID
}
