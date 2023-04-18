package helper

import (
	"new-apex-api/entities"
	"strings"
	"time"
)

const (
	YYYYMMDD   = "2006-01-02"
	YYYYMMDDV2 = "20060102"
	DDMMYYY    = "02/01/2006"
	DDMMMYYY   = ""
)

func ParseTimeStrToDate(dateStr string) (yesterday string, er error) {
	date, er := time.Parse(YYYYMMDD, dateStr)
	if er != nil {
		entities.PrintError(er)
		return yesterday, er
	}

	today := date.AddDate(0, 0, -1)
	yesterday = string(today.Format(YYYYMMDD))

	return yesterday, nil
}

func FormatTimeStrDDMMYYY(input string) (dateStr string, er error) {
	date, er := time.Parse(YYYYMMDD, input)
	if er != nil {
		entities.PrintError(er)
		return dateStr, er
	}

	dateStr = string(date.Format(DDMMYYY))

	return dateStr, nil
}

func AddSeparatorsOnDateStr(input string) (result string, er error) {

	// extract payload
	year := input[0:4]
	month := input[4:6]
	day := input[6:8]

	// array of strings.
	dt := []string{year, month, day}

	// joining the string by separator
	DtWithSeparators := strings.Join(dt, "-")

	return DtWithSeparators, nil
}
