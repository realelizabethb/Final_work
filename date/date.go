// Задание 2 работа с датами
package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func TrueDate(date time.Time, now time.Time) bool {
	return date.After(now)
}

func WorkWithDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", nil
	}
	dateParse, err := time.Parse("20060102", date)
	if err != nil {
		return "", err
	}
	repeatParse := strings.Split(repeat, " ")
	for {
		if repeatParse[0] == "y" {
			dateParse = dateParse.AddDate(1, 0, 0)
		} else if repeatParse[0] == "d" {
			if len(repeatParse) < 2 {
				return "", fmt.Errorf("error with reprat format")
			}
			interval, err := strconv.Atoi(repeatParse[1])
			if err != nil {
				return "", err
			}
			if interval < 0 {
				return "", fmt.Errorf("wrong interval")
			}
			if interval > 400 {
				return "", fmt.Errorf("wrong interval")
			}
			dateParse = dateParse.AddDate(0, 0, interval)
		} else {
			return "", fmt.Errorf("err with format")
		}
		if TrueDate(dateParse, now) {
			break
		}
	}
	resultDate := dateParse.Format("20060102")
	return resultDate, nil
}
