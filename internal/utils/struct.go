package utils

import (
	"reflect"
)

// HasFilters determines if a filter contains any non-null fields.
// The filter must be a pointer to a struct where all fields are pointers.
func HasFilters(filter any) bool {
	v := reflect.ValueOf(filter)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return false
	}

	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			return true
		}
	}
	return false
}

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

		// Check if the field is a pointer and not nil, get its value
		if sourceField.Kind() == reflect.Ptr && !sourceField.IsNil() {
			sourceField = sourceField.Elem()
		}

		// If targetField is a pointer, get its value or initialize it if it is nil
		if targetField.Kind() == reflect.Ptr {
			if targetField.IsNil() {
				targetField.Set(reflect.New(targetField.Type().Elem()))
			}
			targetField = targetField.Elem()
		}

		// Check if the types match or are assignable
		if sourceField.IsValid() && targetField.CanSet() && sourceField.Type().AssignableTo(targetField.Type()) {
			targetField.Set(sourceField)
		}
	}
}
