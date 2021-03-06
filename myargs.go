package myutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

// HandleMonthArgs handles months related arguments
func HandleMonthArgs(s string) (string, error) {

	if s == "" {
		return "", nil
	}

	m := map[string]int{ // mapping for month
		"JAN": 1,
		"FEB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}

	var (
		res   string // result in yyyy-mm-dd format
		month int
	)

	t := time.Now()
	s = strings.ToUpper(s)
	if val, ok := m[s]; ok { // by string, e.g. jan
		month = val
		year := t.Year()
		res = fmt.Sprintf("%d-%02d-01", year, month)

	} else { // Past n month
		n, err := strconv.Atoi(s)
		if err != nil {
			return "", fmt.Errorf("Cant parse month. Should be somthing like 1 (last month) or JAN. \n")
		}
		t = now.With(t).BeginningOfMonth()
		t = t.AddDate(0, -n, 0)
		res = t.Format("2006-01-02")
	}
	return res, nil
}

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

// ParseDateRange TBA
func ParseDateRange(s string, nrecords int, freq string) (string, string, error) {
	var (
		n     int
		err   error
		t     time.Time // input date in date format
		tS    time.Time // start in date format
		tE    time.Time // end date in date format
		start string
		end   string
	)

	t, err = now.Parse(s)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}
	n = nrecords - 1
	if n <= 0 {
		n = 1
	}

	// Set config
	location, _ := time.LoadLocation("Asia/Shanghai")
	tconfig := &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02"},
	}

	end = t.Format("2006-01-02")

	switch freq {
	case "d": // Daily
		if nrecords <= 1 {
			start = t.Format("2006-01-02")
		} else {
			start = t.AddDate(0, 0, -n).Format("2006-01-02")
		}

	case "w": // Weekly
		tS = tconfig.With(t).BeginningOfWeek()
		if nrecords <= 1 {
			// pass
		} else {
			tS = tS.AddDate(0, 0, -n*7)
		}
		tE = tconfig.With(t).EndOfWeek()
		start = tS.Format("2006-01-02")
		end = tE.Format("2006-01-02")

	case "m": //Monthly
		tS = now.With(t).BeginningOfMonth()
		tE = now.With(t).EndOfMonth()
		end = tE.Format("2006-01-02")
		if nrecords <= 1 {
			start = tS.Format("2006-01-02")
		} else {
			start = tS.AddDate(0, -n, 0).Format("2006-01-02")
		}

	default:
		t = time.Now()
	}
	return start, end, nil
}

// ParseDateInput parse the input for past n days, or actual day string in YYYY-MM-DD format
// result depends on freq, daily, monthly -> 2021-03-01
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
			t = time.Now().AddDate(0, 0, -n*7)

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
		dateF = tconfig.With(t).Format("2006-01-02")

	case "m": // Monthly
		dateF = tconfig.With(t).BeginningOfMonth().Format("2006-01-02")

	default:
		dateF = time.Now().Format("2006-01-02")
	}

	return dateF, nil
}
