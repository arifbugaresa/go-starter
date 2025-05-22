package common

import "time"

var defaultLayout = "2006-01-02 15:04:05"

func DefaultFormatDate(time time.Time) string {
	return time.Format(defaultLayout)
}

func FormatDate(time time.Time, layout string) string {
	return time.Format(layout)
}
