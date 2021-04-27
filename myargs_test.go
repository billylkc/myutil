package myutil

import (
	"testing"
	"time"

	"github.com/jinzhu/now"
)

func TestParseDateRange(t *testing.T) {
	type args struct {
		s        string
		nrecords int
		freq     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "same days",
			args: args{
				s:        "2021-04-19",
				nrecords: 1,
				freq:     "d",
			},
			want:  "2021-04-19",
			want1: "2021-04-19",
		},
		{
			name: "3 previous days",
			args: args{
				s:        "2021-04-19",
				nrecords: 3,
				freq:     "d",
			},
			want:  "2021-04-17",
			want1: "2021-04-19",
		},
		{
			name: "2 previous week",
			args: args{
				s:        "2021-04-19",
				nrecords: 2,
				freq:     "w",
			},
			want:  "2021-04-12",
			want1: "2021-04-25",
		},
		{
			name: "This week",
			args: args{
				s:        "2021-04-25",
				nrecords: 1,
				freq:     "w",
			},
			want:  "2021-04-19",
			want1: "2021-04-25",
		},
		{
			name: "This week",
			args: args{
				s:        "2021-04-25",
				nrecords: 0,
				freq:     "w",
			},
			want:  "2021-04-19",
			want1: "2021-04-25",
		},
		{
			name: "same month - April",
			args: args{
				s:        "2021-04-19",
				nrecords: 1,
				freq:     "m",
			},
			want:  "2021-04-01",
			want1: "2021-04-30",
		},
		{
			name: "same month - March",
			args: args{
				s:        "2021-03-01",
				nrecords: 1,
				freq:     "m",
			},
			want:  "2021-03-01",
			want1: "2021-03-31",
		},
		{
			name: "last month",
			args: args{
				s:        "2021-04-05",
				nrecords: 2,
				freq:     "m",
			},
			want:  "2021-03-01",
			want1: "2021-04-30",
		},
		{
			name: "last 2 months across year",
			args: args{
				s:        "2021-01-02",
				nrecords: 2,
				freq:     "m",
			},
			want:  "2020-12-01",
			want1: "2021-01-31",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ParseDateRange(tt.args.s, tt.args.nrecords, tt.args.freq)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDateRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseDateRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseDateInput(t *testing.T) {
	type args struct {
		s    string
		freq string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "zero week",
			args: args{
				s:    "0",
				freq: "w",
			},
			want: time.Now().Format("2006-01-02"),
		},
		{
			name: "one week",
			args: args{
				s:    "1",
				freq: "w",
			},
			want: time.Now().AddDate(0, 0, -7).Format("2006-01-02"),
		},
		{
			name: "two weeks",
			args: args{
				s:    "2",
				freq: "w",
			},
			want: time.Now().AddDate(0, 0, -2*7).Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateInput(tt.args.s, tt.args.freq)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleMonthArgs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "January",
			args: args{
				s: "JAN",
			},
			want: now.BeginningOfYear().Format("2006-01-02"),
		},
		{
			name: "This month",
			args: args{
				s: "0",
			},
			want: now.BeginningOfMonth().Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleMonthArgs(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleMonthArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HandleMonthArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
