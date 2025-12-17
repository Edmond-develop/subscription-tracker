package utils

import "time"

func ParseDates(dateStart, dateEnd string) error {
	_, err := time.Parse("01-2006", dateStart)
	if err != nil {
		return err
	}
	_, err = time.Parse("01-2006", dateEnd)
	if err != nil {
		return err
	}
	return nil
}

func ParseDate(date string) (time.Time, error) {
	newDate, err := time.Parse("01-2006", date)
	if err != nil {
		return time.Time{}, err
	}
	return newDate, nil
}
