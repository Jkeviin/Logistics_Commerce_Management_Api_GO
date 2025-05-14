package utils

import (
	"reflect" // Package for runtime reflection to inspect types and values
)

// Filterable is an interface that filters must implement
type Filterable interface {
	HasFilters() bool
	Matches(any) bool
}

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

// Matches compares an entity against filtering criteria.
// The filter must be a pointer to a struct where all fields are pointers.
func Matches(filter, entity any) bool {
	filterVal := reflect.ValueOf(filter).Elem()
	entityVal := reflect.ValueOf(entity)

	for i := 0; i < filterVal.NumField(); i++ {
		filterField := filterVal.Field(i)
		if !filterField.IsZero() {
			entityField := entityVal.Field(i)
			if filterField.Elem().Interface() != entityField.Interface() {
				return false
			}
		}
	}
	return true
}
