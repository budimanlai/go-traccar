package helpers

import (
	"fmt"
	"time"
)

func NowToString() string {
	return time.Now().Format("2006-01-02")
}

func NormalizeDateFormat(date string) string {
	d, error := time.Parse("2006-01-02 15:04:05", date)

	if error != nil {
		fmt.Println(error)
		return ""
	}

	return d.Format("2006-01-02T15:04:05.999Z")
}
