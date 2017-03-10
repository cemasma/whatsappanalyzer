package wanalyzer

import (
	"reflect"
	"testing"
)

func TestCalculateAggression(t *testing.T) {
	type args struct {
		lines     []string
		negatives []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "TestCalculateAggression",
			args: args{
				lines:     []string{"27.11.2016, 16:11 - Cem Asma: iyi dersler", "27.11.2016, 16:11 - Cem Asma: mal"},
				negatives: []string{"mal", "aptal"},
			},
			want: map[string]int{"Cem Asma": 1, "total": 1},
		},
	}
	for _, tt := range tests {
		if got := CalculateAggression(tt.args.lines, tt.args.negatives); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. CalculateAggression() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAggressionCount(t *testing.T) {
	type args struct {
		lines     []string
		negatives []string
	}
	tests := []struct {
		name           string
		args           args
		wantAggression int
	}{
		{
			name: "TestAggressionCount",
			args: args{
				lines:     []string{"27.11.2016, 16:11 - Cem Asma: iyi dersler", "27.11.2016, 16:11 - Cem Asma: mal"},
				negatives: []string{"mal", "aptal"},
			},
			wantAggression: 1,
		},
	}
	for _, tt := range tests {
		if gotAggression := AggressionCount(tt.args.lines, tt.args.negatives); gotAggression != tt.wantAggression {
			t.Errorf("%q. AggressionCount() = %v, want %v", tt.name, gotAggression, tt.wantAggression)
		}
	}
}
