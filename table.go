package myutil

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
)

// InterfaceSlice converts a list of struct to list of interface
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		// panic("InterfaceSlice() given a non-slice type")
		ret := make([]interface{}, 1)
		ret = append(ret, slice)
		return ret
	}

	ret := make([]interface{}, s.Len())
	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

// PrintTable prints the table with interface
func PrintTable(data []interface{}, headers, ignore []string, colMerge int) error {

	// Build ignore column map from the input list
	// use to filter fields like ID from the result interface
	m := make(map[string]bool)
	for _, key := range ignore {
		if _, ok := m[key]; !ok {
			m[key] = true
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Handle headers
	var headersR table.Row
	for _, h := range headers {
		headersR = append(headersR, h)
	}
	t.AppendHeader(headersR)

	// Construct rows
	var rows []table.Row
	for _, x := range data {
		var r table.Row
		values := reflect.ValueOf(x)
		types := reflect.TypeOf(x)
		for i := 0; i < values.NumField(); i++ {
			field := values.Field(i).Interface()
			name := types.Field(i).Name

			// if name in the ignore map, skip
			if _, ok := m[name]; ok {
				continue
			}
			switch t := field.(type) {
			case time.Time: // for time, format to YYYY-MM-DD
				v := t.Format("2006-01-02")
				r = append(r, v)
			case float64:
				v := humanize.CommafWithDigits(t, 1)
				r = append(r, v)
			case int64:
				v := ""
				if strings.Contains(name, "ID") { // Handle ID column, no comma
					v = fmt.Sprintf("%d", t)
				} else {
					v = humanize.Comma(t)
				}
				r = append(r, v)
			case int:
				v := ""
				if strings.Contains(name, "ID") { // Handle ID column, no comma
					v = fmt.Sprintf("%d", t)
				} else {
					v = humanize.Comma(int64(t))
				}
				r = append(r, v)
			default:
				r = append(r, field)
			}
		}
		rows = append(rows, r)
	}

	// Generate merged column config
	var cc []table.ColumnConfig
	if colMerge >= 1 {
		for i := 1; i <= colMerge; i++ {
			c := table.ColumnConfig{
				Number:    i,
				AutoMerge: true,
				WidthMax:  80,
			}
			cc = append(cc, c)
		}
	}

	// Append rows
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t.AppendRows(rows, rowConfigAutoMerge)
	t.SetColumnConfigs(cc)
	t.AppendSeparator()
	t.Style().Options.SeparateRows = true
	t.Render()
	return nil
}
