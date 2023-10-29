package validation

import (
	"errors"
	"fmt"
	"time"
)

func SevenDayLimitVal(BookingDate string) (string, string, error) {
	format := "2006-01-02"
	parsedTime, err := time.Parse(format, BookingDate)
	if err != nil {
		return "", "", errors.New(err.Error())
	}
	bookTime := parsedTime
	sevenDaysLater := bookTime.Add(7 * 24 * time.Hour)
	sevenDaysAgo := bookTime.Add(-7 * 24 * time.Hour)
	// Format the results
	// layout := "2006-01-02 15:04:05"
	// strMonthCur := int(bookTime.Month())
	strMonthSDL := int(sevenDaysLater.Month())
	strMonthSDA := int(sevenDaysAgo.Month())
	// bookTimeStr := fmt.Sprintf("%d-%d-%d", bookTime.Year(), strMonthCur, bookTime.Day())
	sevenDaysLaterStr := fmt.Sprintf("%d-%d-%d", sevenDaysLater.Year(), strMonthSDL, sevenDaysLater.Day())
	sevenDaysAgoStr := fmt.Sprintf("%d-%d-%d", sevenDaysAgo.Year(), strMonthSDA, sevenDaysAgo.Day())
	return sevenDaysLaterStr, sevenDaysAgoStr, nil
}
