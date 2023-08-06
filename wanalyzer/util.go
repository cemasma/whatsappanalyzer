package wanalyzer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Word a struct, contains use count information
type Word struct {
	Content string
	Value   int
}

type time struct {
	Name  string
	Start int
	End   int
}

var morning = time{Name: "Morning", Start: 6, End: 12}
var noon = time{Name: "Noon", Start: 12, End: 16}
var afternoon = time{Name: "Afternoon", Start: 16, End: 20}
var evening = time{Name: "Evening", Start: 20, End: 24}
var night = time{Name: "Night", Start: 24, End: 6}

// Read reads file and returns as string
func Read(fileAddress string) string {
	byt, err := ioutil.ReadFile(fileAddress)

	if err != nil {
		panic(err)
	}

	return string(byt)
}

// GetLines returns chat record's lines in an array
func GetLines(chatRecord string) []string {
	return strings.Split(chatRecord, "\n")
}

// GetUserLines filters the lines by username
func GetUserLines(lines []string, username string) (specifiedLines []string) {
	for _, value := range lines {
		if strings.Index(value, username) > -1 {
			specifiedLines = append(specifiedLines, value)
		}
	}

	return specifiedLines
}

// GetLinesByDate filters lines by date information and returns filtered lines
func GetLinesByDate(lines []string, date string) (filteredLines []string) {
	formattedDate := strings.ReplaceAll(date, ".", " ")

	for _, value := range lines {
		if strings.Index(value, formattedDate) > -1 {
			filteredLines = append(filteredLines, value)
		}
	}

	return
}

// GetTopicsFromResult provides topics by input
func GetTopicsFromResult(input string) (topics []string) {
	r := regexp.MustCompile(`'.*?'`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		topics = append(topics, strings.ReplaceAll(v[0], "'", ""))
	}

	return
}

// GetUsernames finds the usernames in lines and returns
func GetUsernames(lines []string) (usernames []string) {
	for _, line := range lines {
		value := regexp.MustCompile(`] (.*):`).FindString(line)
		if len(value) > 2 && !Contains(usernames, value[2:len(value)-1]) {
			usernames = append(usernames, value[2:len(value)-1])
		}
	}
	return
}

// SeparateWords parses the words from the lines and returns these in an array
func SeparateWords(lines []string) []string {
	replacer := strings.NewReplacer("\r", "", "\n", "")

	words := []string{}
	for _, value := range lines {
		_, sentence, found := strings.Cut(value, ": ")
		sentence = replacer.Replace(sentence)
		if found {
			wordArr := strings.Split(sentence, " ")
			for _, word := range wordArr {
				if word != " " && len(word) > 0 {
					words = append(words, word)
				}
			}
		}
	}

	return words
}

// GetWordsWithOrder returns sorted Word array by lines
func GetWordsWithOrder(lines []string) []Word {
	words := SeparateWords(lines)
	return SortWordsByCount(words)
}

// PrintWords prints the word array
func PrintWords(words []Word, start, limit int) {
	for i := start; i < limit; i++ {
		fmt.Printf("%d.\t %s\t\t\t\t\tCount: %d\n", i+1, words[i].Content, words[i].Value)
	}
}

// Contains controls the array that is contains the array
func Contains(arr []string, elem string) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}
	return false
}

func getHourTime(hour int) string {
	if hour >= morning.Start && hour <= morning.End {
		return morning.Name
	} else if hour >= noon.Start && hour <= noon.End {
		return noon.Name
	} else if hour >= afternoon.Start && hour <= afternoon.End {
		return afternoon.Name
	} else if hour >= evening.Start && hour <= evening.End {
		return evening.Name
	} else if hour <= night.Start && hour <= night.End {
		return night.Name
	}
	return "time not found"
}

func parseHour(line string) int {
	re := regexp.MustCompile(`\[(.*?)]`)
	foundStr := re.FindString(line)
	if foundStr != "" {
		hour, err := strconv.Atoi(strings.Split(foundStr, " ")[1][0:2])
		if err != nil {
			return 25
		}
		return hour
	}
	return 25
}

func isItIgnored(word string) bool {
	for _, value := range getIgnoredWords() {
		if word == value {
			return false
		}
	}
	return true
}

func makeValuesAsKey(wordsWithCounts map[string]int) map[int][]string {
	wordsByCounts := make(map[int][]string)

	for key, value := range wordsWithCounts {
		wordsByCounts[value] = append(wordsByCounts[value], key)
	}

	return wordsByCounts
}

func getIgnoredWords() []string {
	return []string{"<medya", "atlanmış>", "http", "https", "", "<media", "omitted>"}
}
