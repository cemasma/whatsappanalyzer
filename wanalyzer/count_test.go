package wanalyzer

import (
	"reflect"
	"testing"
)

func TestCountWordInLines(t *testing.T) {
	type args struct {
		word  string
		lines []string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "TestCountWordInLines",
			args: args{
				word:  "iyi",
				lines: []string{"27.11.2016, 16:11 - Cem Asma: iyi dersler", "27.11.2016, 16:11 - Cem Asma: mal"},
			},
			wantCount: 1,
		},
		{
			name: "TestCountWordInLines",
			args: args{
				word:  "iyi",
				lines: []string{"27.11.2016, 16:11 - Cem Asma: iyi dersler iyi", "27.11.2016, 16:11 - Cem Asma: mal iyi"},
			},
			wantCount: 3,
		},
	}
	for _, tt := range tests {
		if gotCount := CountWordInLines(tt.args.word, tt.args.lines); gotCount != tt.wantCount {
			t.Errorf("%q. CountWordInLines() = %v, want %v", tt.name, gotCount, tt.wantCount)
		}
	}
}

func TestSortWordsByCount(t *testing.T) {
	type args struct {
		pureWords []string
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{
			name: "TestSortWordsByCount",
			args: args{pureWords: []string{"selam", "selam", "abi", "abi", "naber", "iyi", "selam", "selam", "iyi"}},
			want: []Word{
				Word{Content: "selam", Value: 4},
				Word{Content: "abi", Value: 2},
				Word{Content: "iyi", Value: 2},
				Word{Content: "naber", Value: 1},
			},
		},
	}
	for _, tt := range tests {
		if got := SortWordsByCount(tt.args.pureWords); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. SortWordsByCount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getMessageCount(t *testing.T) {
	type args struct {
		lines []string
		date  string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "Test_getMessageCount",
			args: args{
				lines: []string{"11.11.2016, 16:11 - Cem Asma: iyi dersler", "1.11.2016, 16:11 - Cem Asma: iyi"},
				date:  "1.11.2016",
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		if gotCount := getMessageCount(tt.args.lines, tt.args.date); gotCount != tt.wantCount {
			t.Errorf("%q. getMessageCount() = %v, want %v", tt.name, gotCount, tt.wantCount)
		}
	}
}

func Test_countWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Test_countWords",
			args: args{words: []string{"test", "deneme", "example", "test", "deneme", "example", "example"}},
			want: map[string]int{"test": 2, "deneme": 2, "example": 3},
		},
	}
	for _, tt := range tests {
		if got := countWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. countWords() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
