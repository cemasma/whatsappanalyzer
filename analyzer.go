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
	file := flag.String("file", "", "You must send it with Whatsapp chat record file address.\n"+
		"\tIf you send just file address it prints top 10 most used words.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\"")

	username := flag.String("username", "", "If you want to query specific user use it.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --username \"Cem Asma\"")

	limit := flag.Int("limit", 10, "It limits list from start to sended value.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --limit 20")

	start := flag.Int("start", 0, "It sets the starting index for list.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --start 10")

	flag.Parse()

	byt, _ := ioutil.ReadFile(*file)

	if len(*username) > 0 {
		// TODO: User's lines will be under review.
		lines := getUserLines(string(byt), *username)
		words := getWordsWithOrder(lines)

		printWords(words, *start, *limit)

	} else if len(*file) > 0 {
		lines := getLines(string(byt))
		words := getWordsWithOrder(lines)

		printWords(words, *start, *limit)
	}
}

func getWordsWithOrder(lines []string) []Word {
	return sortWordsByCount(getWordsByValue(getWordsWithCounts(getWordsInLines(lines))))
}

func printWords(words []Word, start, limit int) {
	for i := start; i < limit; i++ {
		fmt.Printf("%d.\t%s\t\t\t\tCount: %d\n", i+1, words[i].Content, words[i].Value)
	}
}

func getLines(chatRecord string) []string {
	return strings.Split(chatRecord, "\n")
}

func getUserLines(chatRecord, username string) []string {
	lines := strings.Split(chatRecord, "\n")
	specifiedLines := []string{}

	for _, value := range lines {
		if strings.Index(value, username) > -1 {
			specifiedLines = append(specifiedLines, value)
		}
	}

	return specifiedLines
}

func getWordsInLines(lines []string) []string {
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
