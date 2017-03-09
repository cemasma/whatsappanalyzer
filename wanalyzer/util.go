package wanalyzer

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Word struct {
	Content string
	Value   int
}

func Read(fileAddress string) string {
	byt, err := ioutil.ReadFile(fileAddress)

	if err != nil {
		panic(err)
	}

	return string(byt)
}

func GetWordsWithOrder(lines []string) []Word {
	words := SeparateWords(lines)
	return SortWordsByCount(words)
}

func PrintWords(words []Word, start, limit int) {
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

func GetUsernames(lines []string) (usernames []string) {
	for _, line := range lines {
		value := regexp.MustCompile(`- .*?:`).FindString(line)
		if len(value) > 2 && !Contains(usernames, value[2:len(value)-1]) {
			usernames = append(usernames, value[2:len(value)-1])
		}
	}
	return
}

func Contains(arr []string, elem string) bool {
	for _, val := range arr {
		if val == elem {
			return true
		}
	}
	return false
}

func SeparateWords(lines []string) []string {
	words := []string{}
	for _, value := range lines {
		sentence := strings.Split(value, ":")
		if len(sentence) > 2 {
			wordArr := strings.Split(sentence[2], " ")
			words = append(words, wordArr...)
		}
	}

	return words
}

func isNotIgnored(word string) bool {
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
