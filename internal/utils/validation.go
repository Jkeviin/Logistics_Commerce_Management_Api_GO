package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

func ValidateStructGoValidator(structToValidate any) error {
	_, err := govalidator.ValidateStruct(structToValidate)
	if err != nil {
		// Split errors and get the first error report
		firstError := fmt.Sprintf("%v", err)
		if errorsSlice := strings.Split(firstError, ";"); len(errorsSlice) > 0 {
			firstError = errorsSlice[0]
		}
		return errors.New(firstError)
	}
	return nil
}
