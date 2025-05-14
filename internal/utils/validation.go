package utils

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"reflect"
	"regexp"
	"strings"
)

const (
	// TagName is the name of the struct tag used to specify validation rules.
	notNilTag  = "notNil"
	float19e2d = "float19_2"
)

var (
	regex = regexp.MustCompile(`^-?\d{1,19}(\.\d{1,2})?$`)
)

func init() {
	// Register custom validation rule for not nil fields.
	// The 'notNil' tag is registered to always return true because the actual nil-check
	// is performed later by the validateNotNilFields function. This design choice ensures
	// that the validation logic is centralized and avoids duplication.
	govalidator.CustomTypeTagMap.Set(notNilTag, func(i interface{}, context interface{}) bool {
		return true
	})

	// Register a custom validation rule for floats with up to 19 integer digits and 2 decimal places
	govalidator.CustomTypeTagMap.Set(float19e2d, func(i interface{}, context interface{}) bool {
		var number float64

		switch v := i.(type) {
		case float32:
			number = float64(v)
		case *float32:
			if v == nil {
				return false
			}
			number = float64(*v)
		case float64:
			number = v
		case *float64:
			if v == nil {
				return false
			}
			number = *v
		default:
			return false
		}

		return regex.MatchString(fmt.Sprintf("%.2f", number))
	})
}

func ValidateStructGoValidator(structToValidate any) error {
	if err := validateNotNilFields(structToValidate); err != nil {
		return err
	}
	if _, err := govalidator.ValidateStruct(structToValidate); err != nil {
		// Split errors and get the first error report
		firstError := fmt.Sprintf("%v", err)
		if errorsSlice := strings.Split(firstError, ";"); len(errorsSlice) > 0 {
			firstError = errorsSlice[0]
		}
		return errors.New(firstError)
	}
	return nil
}

// validateNotNilFields validates that fields with the "notNil" tag are not nil.
func validateNotNilFields(structToValidate any) error {
	val := reflect.ValueOf(structToValidate)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return errors.New("input must be a struct or a pointer to a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get("valid")
		if strings.HasPrefix(tag, notNilTag) {
			if val.Field(i).IsNil() {
				message := strings.SplitN(strings.TrimPrefix(tag, notNilTag+"~"), ",", 2)[0]
				if message == "" {
					message = fmt.Sprintf("field %s must not be nil", field.Name)
				}
				return errors.New(message)
			}
		}
	}
	return nil
}
