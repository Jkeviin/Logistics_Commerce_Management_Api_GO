package utils

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (cd *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1]

	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		return err
	}

	cd.Time = parsedTime
	return nil
}

func (cd Date) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", cd.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}
