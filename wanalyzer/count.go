package wanalyzer

import (
	"regexp"
	"strings"
)

func CountWordInLines(word string, lines []string) (count int) {
	if strings.Index(word, " ") > -1 {
		for _, value := range lines {
			if strings.Index(value, word) > -1 {
				count += len(regexp.MustCompile("\\b"+word+"\\b").FindAllString(value, -1))
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

func SortWordsByCount(pureWords []string) []Word {
	words := []Word{}
	wordCountMap := makeValuesAsKey(countWords(pureWords))

	for count, wordArr := range wordCountMap {
		for _, word := range wordArr {
			if isItIgnored(word) {
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

func getMessageCount(lines []string, date string) (count int) {
	for _, line := range lines {
		if (len(date) == 10 && len(line) > 9 && line[:10] == date) || (len(date) == 9 && len(line) > 8 && line[:9] == date) {
			count++
		}
	}
	return
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
