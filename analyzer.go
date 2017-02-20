package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Word struct {
	Key   string
	Value int
}

func main() {
	byt, _ := ioutil.ReadFile("C:\\Users\\cemasma\\Desktop\\langüage öf ıngilazca ile WhatsApp Sohbeti.txt")
	//fmt.Println(string(byt))

	words := getWordsByValue(getWordsWithCounts(getWords(string(byt))))
	fmt.Println(words)

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
