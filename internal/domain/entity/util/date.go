package util_entity

import (
	"time"
)

func GetDateFromString(date string) (time.Time, error) {
	const shortForm = "2006-01-02"
	result, err := time.Parse(shortForm, date)
	return result, err
}

func GetStringFromDate(date time.Time) string {
	const shortForm = "2006-01-02"
	return date.Format(shortForm)
}

func GetDateTimeFromString(date string) (time.Time, error) {
	result, err := time.Parse(time.DateTime, date)
	return result, err
}

func GetStringFromDateTime(date time.Time) string {
	return date.Format(time.DateTime)
}
