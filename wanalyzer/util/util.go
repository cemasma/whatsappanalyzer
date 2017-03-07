package util

import "strings"

type Word struct {
	Content string
	Value   int
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

func SortWordsByCount(pureWords []string) []Word {
	words := []Word{}
	wordCountMap := makeValuesAsKey(countWords(pureWords))

	for count, wordArr := range wordCountMap {
		for _, word := range wordArr {
			if isNotIgnored(word) {
				words = append(words, Word{Content: word, Value: count})
			}
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i].Value < words[j].Value {
				backup := words[i]
				words[i] = words[j]
				words[j] = backup
			}
		}
	}

	return words
}

func GetMessageCount(lines []string, date string) (count int) {
	for _, line := range lines {
		if (len(date) == 10 && len(line) > 9 && line[:10] == date) || (len(date) == 9 && len(line) > 8 && line[:9] == date) {
			count++
		}
	}
	return
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

func countWords(words []string) map[string]int {
	wordsWithCounts := make(map[string]int)

	for _, value := range words {
		value = strings.ToLower(value)
		if _, ok := wordsWithCounts[value]; ok {
			wordsWithCounts[value]++
		} else {
			wordsWithCounts[value] = 1
		}
	}
	return wordsWithCounts
}
