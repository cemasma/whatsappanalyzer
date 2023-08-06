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
	}{
		{
			name: "TestCalculateAggression",
			args: args{
				lines: []string{"[27.11.2016 16:11] Cem Asma: iyi dersler", "[27.11.2016 16:11] Cem Asma: test"},
			},
		},
	}
	for _, tt := range tests {
		got := CalculateAggression(tt.args.lines)
		if _, ok := got["Cem Asma"]; !ok {
			t.Errorf("%q. CalculateAggression() = %v, key is not exists = %v", tt.name, got, "Cem Asma")
		}
		if _, ok := got["total"]; !ok {
			t.Errorf("%q. CalculateAggression() = %v, key is not exists = %v", tt.name, got, "total")
		}
	}
}

func TestAggressionCount(t *testing.T) {
	type args struct {
		lines     []string
		negatives []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestAggressionCount",
			args: args{
				lines: []string{"27.11.2016, 16:11 - Cem Asma: iyi dersler", "27.11.2016, 16:11 - Cem Asma: test"},
			},
		},
	}
	for _, tt := range tests {
		if gotAggression := AggressionCount(tt.args.lines); reflect.TypeOf(gotAggression).Kind() != reflect.Float64 {
			t.Errorf("%q. AggressionCount() = %v, type is not matching = float64", tt.name, gotAggression)
		}
	}
}
