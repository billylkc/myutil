package myutil

import "testing"

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
