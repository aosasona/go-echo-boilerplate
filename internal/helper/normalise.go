package helper

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
)

func CapitalizeFirst(str string) string {
	if len(str) <= 1 {
		return str
	}
	str = strings.ToLower(str)
	return strings.ToUpper(string(str[0])) + string(str[1:])
}

func JSONMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JSONUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func ToHumanTime(rawDateInHours float64) string {
	var (
		dateInHours = int(rawDateInHours)
		years       = dateInHours / 8760
		months      = dateInHours / 730
		weeks       = dateInHours / 168
		days        = dateInHours / 24
		hours       = dateInHours % 24
	)

	if years > 0 {
		return Pluralize(years, "year")
	}

	if months > 0 {
		return Pluralize(months, "month")
	}

	if weeks > 0 {
		return Pluralize(weeks, "week")
	}

	if days > 0 {
		return Pluralize(days, "day")
	}

	if hours > 0 {
		return Pluralize(hours, "hour")
	}

	return "today"
}

func Pluralize(count int, word string) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, word)
	}

	return fmt.Sprintf("%d %ss", count, word)
}
