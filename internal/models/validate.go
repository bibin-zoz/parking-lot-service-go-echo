package models

import (
	"reflect"
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	// Register custom validators
	validate.RegisterValidation("nonZeroPositive", validateNonZeroPositive)
	validate.RegisterValidation("nameLength", validateNameLength)
	validate.RegisterValidation("parkingSpots", validateParkingSpots)
	validate.RegisterValidation("vehicleNumber", validateVehicleNumber) // Register the vehicle number validator
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateNonZeroPositive(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val := field.Uint()
		return val > 0 && val < 10000
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

// Custom validator for vehicle number
func validateVehicleNumber(fl validator.FieldLevel) bool {
	vehicleNumber := fl.Field().String()

	// Check length
	if len(vehicleNumber) < 7 || len(vehicleNumber) > 10 {
		return false
	}

	// Count numbers and alphabets
	var numCount, alphaCount int
	for _, char := range vehicleNumber {
		if unicode.IsDigit(char) {
			numCount++
		} else if unicode.IsLetter(char) {
			alphaCount++
		}
	}

	// Ensure at least 4 numbers and 2 alphabets
	if numCount < 4 || alphaCount < 2 {
		return false
	}

	// Regular expression for format validation
	regex := `^[A-Z0-9]{7,10}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(vehicleNumber)
}
