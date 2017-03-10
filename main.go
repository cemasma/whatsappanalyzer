package main

import (
	"flag"
	"fmt"
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

	negativesFile := flag.String("negatives", "", "If you want to measure aggression use it.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --negatives \"C:\\negativewords.txt\" ")

	messageFrequency := flag.Bool("messagef", false, "It measure the talking frequency in date by date."+
		"\tExample: analyzer --file \"C:\\filename.txt\" --messagef")

	printFrequency := flag.Bool("printf", false, "It sorts frequency by activity and prints."+
		"\tExample: analyzer --file \"C:\\filename.txt\" --messagef --printf")
	flag.Parse()

	var lines []string

	if len(*file) > 0 {
		lines = wanalyzer.GetLines(wanalyzer.Read(*file))
	}

	if len(*username) > 0 {
		lines = wanalyzer.GetUserLines(lines, *username)
	}

	if len(*word) > 0 {

		count := wanalyzer.CountWordInLines(*word, lines)
		fmt.Printf("%s\t\t\t\t\tCount: %d", *word, count)
	} else if len(*negativesFile) > 0 && len(*username) > 0 {
		count := wanalyzer.AggressionCount(lines, strings.Split(wanalyzer.Read(*negativesFile), "\r\n"))

		fmt.Printf("%s's aggression count is: %d", *username, count)
	} else if len(*negativesFile) > 0 {
		calculatedAgg := wanalyzer.CalculateAggression(lines, strings.Split(wanalyzer.Read(*negativesFile), "\r\n"))

		for key, val := range calculatedAgg {
			if key != "total" {
				percent := (float64((val * 100)) / float64(calculatedAgg["total"]))
				fmt.Printf("%s %f percent aggressive by total. \n", key, percent)
				fmt.Printf("%s's using count is %d\n\n", key, val)
			}
		}

		fmt.Printf("The total negative aggression count is: %d", calculatedAgg["total"])

	} else if *messageFrequency {

		dates := wanalyzer.GetDatesFromLines(lines)
		frequence := wanalyzer.GetMessageFrequency(lines, dates)

		if *printFrequency {
			frequence = wanalyzer.SortFrequency(frequence)
			wanalyzer.PrintMessageFrequence(frequence, *start, *limit)
		} else {
			drawer := wanalyzer.NewGraph(*username + " mf.png")
			drawer.DrawFrequence(frequence)

			fmt.Println("Graph of messaging frequency is created.")
		}

	} else if len(*file) > 0 {

		words := wanalyzer.GetWordsWithOrder(lines)
		wanalyzer.PrintWords(words, *start, *limit)
	}
}
