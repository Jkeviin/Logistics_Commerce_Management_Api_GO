package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

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

// ParseIDFromRequest extracts the 'id' parameter from the URL and converts it to int64
func ParseIDFromRequest(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return ParseIDInt64(idStr)
}

func ParseDateTime(expirationDate []uint8) (*time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02 15:04:05", string(expirationDate))
	if err != nil {
		return nil, fmt.Errorf("error parsing expiration_date: %w", err)
	}
	return &parsedDate, nil
}

func ParseDate(str string) (*Date, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)
	if err != nil {
		return nil, err
	}
	return &Date{Time: t}, nil
}
