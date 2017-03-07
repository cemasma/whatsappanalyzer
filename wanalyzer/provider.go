package wanalyzer

import "strings"
import "fmt"
import "./util"

type MessageFrequence struct {
	Date  string
	Count int
}

func GetMessageFrequency(lines, dates []string) (frequency []MessageFrequence) {
	for _, date := range dates {
		count := util.GetMessageCount(lines, date)
		frequency = append(frequency, MessageFrequence{Date: date, Count: count})
	}

	return frequency
}

func GetDatesFromLines(lines []string) (dates []string) {
	for _, line := range lines {
		if len(line) > 10 && (string(line[2]) == "." && string(line[5]) == "." && string(line[10]) == ",") &&
			string(line[14]) == ":" && !util.Contains(dates, string(line[:10])) {

			dates = append(dates, line[:10])
		} else if len(line) > 9 && string(line[1]) == "." && string(line[4]) == "." && string(line[9]) == "," &&
			string(line[13]) == ":" && !util.Contains(dates, string(line[:9])) {

			dates = append(dates, line[:9])
		}
	}
	return
}

func CalculateAggression(lines, negatives []string) (aggression int) {
	for _, word := range negatives {
		aggression += FindWordInLines(word, lines)
	}

	return
}

func FindWordInLines(word string, lines []string) (count int) {
	if strings.Index(word, " ") > -1 {
		for _, value := range lines {
			if strings.Index(value, word) > -1 {
				count++
			}
		}
	} else {
		separatedWords := GetWordsWithOrder(lines)
		for _, value := range separatedWords {
			if value.Content == word {
				count = value.Value
				return
			}
		}
	}

	return
}

func GetWordsWithOrder(lines []string) []util.Word {
	words := util.SeparateWords(lines)
	return util.SortWordsByCount(words)
}

func PrintWords(words []util.Word, start, limit int) {
	for i := start; i < limit; i++ {
		fmt.Printf("%d.\t %s\t\t\t\t\tCount: %d\n", i+1, words[i].Content, words[i].Value)
	}
}

func GetLines(chatRecord string) []string {
	return strings.Split(chatRecord, "\n")
}

func GetUserLines(lines []string, username string) (specifiedLines []string) {
	for _, value := range lines {
		if strings.Index(value, username) > -1 {
			specifiedLines = append(specifiedLines, value)
		}
	}

	return specifiedLines
}
