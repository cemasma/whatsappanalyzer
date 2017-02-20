package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	byt, _ := ioutil.ReadFile("C:\\Users\\cemasma\\Desktop\\langüage öf ıngilazca ile WhatsApp Sohbeti.txt")
	//fmt.Println(string(byt))

	words := getWordsWithCounts(getWords(string(byt)))
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

}
