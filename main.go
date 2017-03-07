package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"./wanalyzer"
)

func main() {
	file := flag.String("file", "", "You must send it with Whatsapp chat record file address.\n"+
		"\tIf you send just file address it prints top 10 most used words.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\"")

	username := flag.String("username", "", "If you want to query specific user then use it.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --username \"Cem Asma\"")

	limit := flag.Int("limit", 10, "It limits list from start to sended value.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --limit 20")

	start := flag.Int("start", 0, "It sets the starting index for list.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --start 10")

	word := flag.String("word", "", "If you want to see information about the specific word use it.\n"+
		"\tLetters of the word must be lower."+
		"\tExample: analyzer --file \"C:\\filename.txt\" --word \"test\"")

	negativesFileAddress := flag.String("negatives", "", "If you want to measure aggression use it.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --negatives \"C:\\negativewords.txt\" ")

	messageFrequency := flag.Bool("mf", false, "It measure the talking frequency in date by date")

	flag.Parse()

	chatByt, _ := ioutil.ReadFile(*file)
	lines := wanalyzer.GetLines(string(chatByt))

	if len(*username) > 0 {
		lines = wanalyzer.GetUserLines(lines, *username)
	}

	if len(*word) > 0 {
		count := wanalyzer.FindWordInLines(*word, lines)
		fmt.Printf("%s\t\t\t\t\tCount: %d", *word, count)
	} else if len(*negativesFileAddress) > 0 {
		negativesByt, _ := ioutil.ReadFile(*negativesFileAddress)
		calculatedAgg := wanalyzer.CalculateAggression(lines, strings.Split(string(negativesByt), "\r\n"))

		fmt.Printf("The aggression level is: %d", calculatedAgg)
	} else if *messageFrequency {
		dates := wanalyzer.GetDatesFromLines(lines)
		frequence := wanalyzer.GetMessageFrequency(lines, dates)

		wanalyzer.DrawFrequence(frequence, *username+" mf.png")
	} else if len(*file) > 0 {
		words := wanalyzer.GetWordsWithOrder(lines)
		wanalyzer.PrintWords(words, *start, *limit)
	}
}
