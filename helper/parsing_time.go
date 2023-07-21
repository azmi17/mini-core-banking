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

func ParseTimeStrToDateOnYesterday(dateStr string) (yesterday string, er error) {
	date, er := time.Parse(YYYYMMDD, dateStr)
	if er != nil {
		entities.PrintError(er)
		return yesterday, er
	}

	today := date.AddDate(0, 0, -1)
	yesterday = string(today.Format(YYYYMMDD))

	return yesterday, nil
}

// Rev
func FormatTimeStrDDMMYYY(input string) (dateStr string, er error) {
	date, er := time.Parse(YYYYMMDD, input)
	if er != nil {
		entities.PrintError(er)
		return dateStr, er
	}

	dateStr = string(date.Format(DDMMYYY))

	return dateStr, nil
}

func StringFormatDMYToYMDWithoutSeparators(input string) (ymd string, er error) {
	dmy := strings.Replace(input, "/", "", 2)

	day := dmy[0:2]
	month := dmy[2:4]
	year := dmy[4:8]

	ymd = year + month + day

	return ymd, nil

}

// Rev
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

func GetCurrentDate(timeFormatStr string) string {
	now := time.Now()
	var getCurrentDate = string(now.Format(timeFormatStr))
	return getCurrentDate
}
