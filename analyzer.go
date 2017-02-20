package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type Word struct {
	Content string
	Value   int
}

func main() {
	file := flag.String("file", "", "Whatsapp chat record file address.")
	username := flag.String("username", "", "Provides specified analysis by username.")
	limit := flag.Int("limit", 10, "It limits list from start to sended value.")
	flag.Parse()

	byt, _ := ioutil.ReadFile(*file)

	if len(*username) > 0 {
		// TODO: User's lines will be under review.
	} else if len(*file) > 0 {
		words := sortWordsByCount(getWordsByValue(getWordsWithCounts(getWords(string(byt)))))

		for i := 0; i < *limit; i++ {
			fmt.Printf("Word: %s 		Count: %d\n", words[i].Content, words[i].Value)
		}
	}
}

func getWords(chatRecord string) []string {
	lines := strings.Split(chatRecord, "\n")
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

func getWordsWithCounts(words []string) map[string]int {
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

func getWordsByValue(wordsWithCounts map[string]int) map[int][]string {
	wordsByCounts := make(map[int][]string)

	for key, value := range wordsWithCounts {
		wordsByCounts[value] = append(wordsByCounts[value], key)
	}

	return wordsByCounts
}

func sortWordsByCount(wordsByCount map[int][]string) []Word {
	words := []Word{}

	for count, wordArr := range wordsByCount {
		for _, word := range wordArr {
			if word != "" && word != "atlanmış>" && word != "<medya" {
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
