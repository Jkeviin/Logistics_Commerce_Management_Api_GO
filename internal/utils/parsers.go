package utils

import "strconv"

// ParseIDInt64 parses the ID from the request and validates it
func ParseIDInt64(id string) (int64, error) {
	if id == "" {
		return 0, ErrMandatoryId
	}
	result, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, ErrInvalidId
	}
	return result, nil
}
