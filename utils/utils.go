package utils

import (
	"fmt"
	"os"
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

func Print(s string) {
	fmt.Println("****************************************************************************************************")
	fmt.Println(s)
	fmt.Println("****************************************************************************************************")
}

func IsTerminal() (bool, error) {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false, err
	}
	if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
