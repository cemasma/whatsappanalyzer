package wanalyzer

import (
	"reflect"
	"testing"
)

func TestSortFrequency(t *testing.T) {
	type args struct {
		frequence []MessageFrequence
	}
	tests := []struct {
		name string
		args args
		want []MessageFrequence
	}{
		{
			name: "TestSortFrequency",
			args: args{frequence: []MessageFrequence{
				{Date: "1.01.2017", Count: 1000},
				{Date: "2.01.2017", Count: 1001},
			}},
			want: []MessageFrequence{
				{Date: "2.01.2017", Count: 1001},
				{Date: "1.01.2017", Count: 1000},
			},
		},
	}
	for _, tt := range tests {
		if got := sortFrequency(tt.args.frequence); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SortFrequency() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetMessageFrequency(t *testing.T) {
	type args struct {
		lines []string
		dates []string
	}
	tests := []struct {
		name          string
		args          args
		wantFrequency []MessageFrequence
	}{
		{
			name: "TestGetMessageFrequency",
			args: args{
				lines: []string{
					"[27.11.2016 16:11] Cem Asma: iyi dersler",
					"[27.11.2016 16:11] Cem Asma: iyi çalışmalar",
					"[28.11.2016 16:11] Cem Asma: iyi günler",
					"[31.12.2016 16:11] Cem Asma: iyi seneler",
				},
				dates: []string{"27.11.2016", "28.11.2016", "31.12.2016"},
			},
			wantFrequency: []MessageFrequence{
				{Date: "27.11.2016", Count: 2},
				{Date: "28.11.2016", Count: 1},
				{Date: "31.12.2016", Count: 1},
			},
		},
	}
	for _, tt := range tests {
		if gotFrequency := GetMessageFrequency(tt.args.lines, tt.args.dates); !reflect.DeepEqual(gotFrequency, tt.wantFrequency) {
			t.Errorf("%q. GetMessageFrequency() = %v, want %v", tt.name, gotFrequency, tt.wantFrequency)
		}
	}
}

func TestGetDatesFromLines(t *testing.T) {
	type args struct {
		lines []string
	}

	tests := []struct {
		name      string
		args      args
		wantDates map[string]int
	}{
		{
			name: "TestGetDatesFromLines",
			args: args{
				lines: []string{
					"[27.11.2016 16:11] Cem Asma: iyi dersler",
					"[2.12.2016 16:11] Cem Asma: iyi günler",
					"[31.12.2016 16:11] Cem Asma: iyi seneler",
					"[31.12.2016 16:11] Cem Asma: iyi seneler",
				},
			},
			wantDates: map[string]int{"27.11.2016": 1, "2.12.2016": 1, "31.12.2016": 2},
		},
	}
	for _, tt := range tests {
		if gotDates := GetDatesFromLines(tt.args.lines); !reflect.DeepEqual(gotDates, tt.wantDates) {
			t.Errorf("%q. GetDatesFromLines() = %v, want %v", tt.name, gotDates, tt.wantDates)
		}
	}
}

func TestGetTimeFrequency(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name              string
		args              args
		wantTimeFrequence map[string]int
	}{
		{
			name: "TestGetTimeFrequency",
			args: args{
				lines: []string{
					"[19.02.2017, 17:49] Cem Asma: test",
					"[9.02.2017, 04:49] Cem Asma: test",
					"[9.02.2017, ..:49] Cem Asma: test",
				},
			},
			wantTimeFrequence: map[string]int{
				"Morning":   0,
				"Noon":      0,
				"Afternoon": 1,
				"Evening":   0,
				"Night":     1,
			},
		},
	}
	for _, tt := range tests {
		if gotTimeFrequence := GetTimeFrequency(tt.args.lines); !reflect.DeepEqual(gotTimeFrequence, tt.wantTimeFrequence) {
			t.Errorf("%q. GetTimeFrequency() = %v, want %v", tt.name, gotTimeFrequence, tt.wantTimeFrequence)
		}
	}
}
