package util

import "time"

const dateTimeLayout = "2006-01-02T15:04:05"

func StringToDateTime(t time.Time) (time.Time, error) {
	timeString := t.Format(dateTimeLayout)
	parsedTime, err := time.Parse(dateTimeLayout, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
