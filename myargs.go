package myutil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/now"
)

// HandleDateArgs handles non flag input arguments
// mostly handle nrecords only, where n could be days, weeks, or months
func HandleDateArgs(date *string, nrecords *int, defaultN int, args ...string) error {
	var err error
	if len(args) == 1 {
		*date = args[0]
		*nrecords = defaultN
	} else if len(args) == 2 {
		*date = args[0]
		*nrecords, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}
	} else {
		*nrecords = defaultN
	}
	return nil
}

// ParseDateInput parse the input for past n days, or actual day string in YYYY-MM-DD format
// result depends on freq, daily, monthly -> 2021-03-01, weekly -> start from monday
func ParseDateInput(s, freq string) (string, error) {
	var (
		t     time.Time
		dateF string
		err   error
	)

	// Set config
	location, _ := time.LoadLocation("Asia/Shanghai")
	tconfig := &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02"},
	}

	// Check if input is in YYYY-MM-DD format
	if len(s) == 10 {
		t, err = now.Parse(s)
		if err != nil {
			return dateF, fmt.Errorf("Invalid input for date. Need a date in YYYY-MM-DD format or number for past n days/weeks/months.")
		}
	} else {
		// Convert to date
		n, err := strconv.Atoi(s)
		if err != nil {
			return dateF, fmt.Errorf("Invalid input for date. Need a date in YYYY-MM-DD format or number for past n days/weeks/months")
		}
		switch freq {
		case "d": // Daily
			t = time.Now().AddDate(0, 0, -n)

		case "w": // Weekly
			t = time.Now().AddDate(0, 0, -n*7+6)

		case "m": //Monthly
			t = time.Now().AddDate(0, -n, 0)

		default:
			t = time.Now()
		}
	}

	// Handle frequency
	switch freq {
	case "d": // Daily
		dateF = t.Format("2006-01-02")

	case "w": // Weekly.
		t := tconfig.With(t).BeginningOfWeek()
		dateF = tconfig.With(t).Format("2006-01-02")

	case "m": // Monthly
		dateF = tconfig.With(t).BeginningOfMonth().Format("2006-01-02")

	default:
		dateF = time.Now().Format("2006-01-02")
	}

	return dateF, nil
}
