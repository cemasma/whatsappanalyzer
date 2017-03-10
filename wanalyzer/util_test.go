package wanalyzer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		fileAddress string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestRead",
			args: args{
				fileAddress: "testfile",
			},
			want: "test\r\nasd",
		},
	}
	for _, tt := range tests {
		if got := Read(tt.args.fileAddress); got != tt.want {
			t.Errorf("%q. Read() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetLines(t *testing.T) {
	type args struct {
		chatRecord string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetLines",
			args: args{"test\nasd\ndeneme"},
			want: []string{"test", "asd", "deneme"},
		},
	}
	for _, tt := range tests {
		if got := GetLines(tt.args.chatRecord); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetLines() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetUserLines(t *testing.T) {
	type args struct {
		lines    []string
		username string
	}
	tests := []struct {
		name               string
		args               args
		wantSpecifiedLines []string
	}{
		{
			name: "TestGetUserLines",
			args: args{
				lines:    []string{"Ali Ozan Asma: selam", "Cem Asma: test", "Ali Ozan Asma: test"},
				username: "Ali Ozan Asma",
			},
			wantSpecifiedLines: []string{"Ali Ozan Asma: selam", "Ali Ozan Asma: test"},
		},
	}
	for _, tt := range tests {
		if gotSpecifiedLines := GetUserLines(tt.args.lines, tt.args.username); !reflect.DeepEqual(gotSpecifiedLines, tt.wantSpecifiedLines) {
			t.Errorf("%q. GetUserLines() = %v, want %v", tt.name, gotSpecifiedLines, tt.wantSpecifiedLines)
		}
	}
}

func TestGetUsernames(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name          string
		args          args
		wantUsernames []string
	}{
		{
			name: "TestGetUsernames",
			args: args{
				lines: []string{
					"27.11.2016, 16:11 - Cem Asma: iyi dersler",
					"5.12.2016, 12:27 - Ali Ozan Asma: selam",
					"6.12.2016, 12:27 - Ali Ozan Asma: test",
				},
			},
			wantUsernames: []string{"Cem Asma", "Ali Ozan Asma"},
		},
	}
	for _, tt := range tests {
		if gotUsernames := GetUsernames(tt.args.lines); !reflect.DeepEqual(gotUsernames, tt.wantUsernames) {
			t.Errorf("%q. GetUsernames() = %v, want %v", tt.name, gotUsernames, tt.wantUsernames)
		}
	}
}

func TestSeparateWords(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestSeparateWords",
			args: args{
				lines: []string{
					"27.11.2016, 16:11 - Cem Asma: iyi dersler",
					"5.12.2016, 12:27 - Ali Ozan Asma: selam",
					"6.12.2016, 12:27 - Ali Ozan Asma: test test",
				},
			},
			want: []string{"iyi", "dersler", "selam", "test", "test"},
		},
	}
	for _, tt := range tests {
		if got := SeparateWords(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
			for i, val := range got {
				fmt.Println(i, val)
			}
			t.Errorf("%q. SeparateWords() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetWordsWithOrder(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{
			name: "TestGetWordsWithOrder",
			args: args{
				lines: []string{
					"27.11.2016, 16:11 - Cem Asma: selam iyi dersler",
					"5.12.2016, 12:27 - Ali Ozan Asma: selam",
					"6.12.2016, 12:27 - Ali Ozan Asma: test test",
				},
			},
			want: []Word{
				Word{Content: "selam", Value: 2},
				Word{Content: "test", Value: 2},
				Word{Content: "iyi", Value: 1},
				Word{Content: "dersler", Value: 1},
			},
		},
	}
	for _, tt := range tests {
		if got := GetWordsWithOrder(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetWordsWithOrder() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestContains(t *testing.T) {
	type args struct {
		arr  []string
		elem string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestContains",
			args: args{arr: []string{"test", "deneme"}, elem: "test"},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := Contains(tt.args.arr, tt.args.elem); got != tt.want {
			t.Errorf("%q. Contains() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_isItIgnored(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isItIgnored",
			args: args{word: "<media"},
			want: false,
		},
		{
			name: "Test_isItIgnored",
			args: args{word: "test"},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := isItIgnored(tt.args.word); got != tt.want {
			t.Errorf("%q. isItIgnored() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_makeValuesAsKey(t *testing.T) {
	type args struct {
		wordsWithCounts map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[int][]string
	}{
		{
			name: "Test_makeValueAsKey",
			args: args{
				wordsWithCounts: map[string]int{"test": 2, "deneme": 1, "cem": 1},
			},
			want: map[int][]string{2: []string{"test"}, 1: []string{"deneme", "cem"}},
		},
	}
	for _, tt := range tests {
		if got := makeValuesAsKey(tt.args.wordsWithCounts); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. makeValuesAsKey() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
