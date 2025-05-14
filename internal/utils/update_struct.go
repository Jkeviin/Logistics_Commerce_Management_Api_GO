package utils

import "reflect"

// UpdateStruct Updates the fields of target with the values from source if they are not empty
// target: The target struct to update
// source: The source struct to get the values from, the struct send all the fields must be pointers
func UpdateStruct[T any, TPointer any](target *T, source *TPointer) {
	targetVal := reflect.ValueOf(target).Elem()
	sourceVal := reflect.ValueOf(source).Elem()

	for i := 0; i < sourceVal.NumField(); i++ {
		sourceField := sourceVal.Field(i)
		sourceFieldName := sourceVal.Type().Field(i).Name

		// Check if the field exists in target
		targetField := targetVal.FieldByName(sourceFieldName)
		if !targetField.IsValid() {
			continue // If the field does not exist, skip to the next one
		}

		// If sourceField is a pointer and not nil, get its value
		if sourceField.Kind() == reflect.Ptr && !sourceField.IsNil() {
			sourceField = sourceField.Elem()
		}

		// Check if the type matches or is assignable
		if sourceField.IsValid() && targetField.CanSet() && sourceField.Type().AssignableTo(targetField.Type()) {
			targetField.Set(sourceField)
		}
	}
}
