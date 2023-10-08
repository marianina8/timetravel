package utils

import (
	"strings"
	"time"
)

func FormatDestination(t time.Time) string {
	return strings.ToUpper(t.Format("Jan 02 2006 (pm) 03:04"))
}

func ParseDestination(s string) (time.Time, error) {
	if s == "" {
		return time.Now(), nil
	}
	const layout = "010220061504"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
