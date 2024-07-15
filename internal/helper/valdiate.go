package models

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	// Register custom validators
	validate.RegisterValidation("nonZeroPositive", validateNonZeroPositive)
	validate.RegisterValidation("nameLength", validateNameLength)
	validate.RegisterValidation("parkingSpots", validateParkingSpots)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateNonZeroPositive(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := field.Uint()
		return val > 0
	}
	return false
}

func validateNameLength(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		val := field.String()
		return len(val) >= 4 && len(val) <= 20
	}
	return false
}

func validateParkingSpots(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Int:
		val := field.Int()
		return val >= 1 && val <= 500
	}
	return false
}
