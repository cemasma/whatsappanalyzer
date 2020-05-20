package main

import (
	"flag"
	"fmt"

	"whatsappanalyzer/wanalyzer"
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
		"\tLetters of the word must be lower.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --word \"test\"")

	aggression := flag.Bool("aggression", false, "If you want to measure aggression use it.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\"  --username \"Cem Asma\" --aggression true")

	messageFrequency := flag.Bool("messagef", false, "It measure the messaging frequency in date by date.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --messagef")

	timeFrequency := flag.Bool("timef", false, "It measure messaging frequency in time periods.\n"+
		"\tExample: analyzer --file --file \"C:\\filename.txt\" --timef")

	printFrequency := flag.Bool("printf", false, "It sorts frequency by activity and prints.\n"+
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
	} else if *aggression == true && len(*username) > 0 {
		count := wanalyzer.AggressionCount(lines)

		fmt.Printf("%s's aggression count is: %f", *username, count)
	} else if *aggression {
		calculatedAgg := wanalyzer.CalculateAggression(lines)

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

			fmt.Println("Graph of messaging frequency was created.")
		}

	} else if *timeFrequency {
		timeFrequences := wanalyzer.GetTimeFrequency(lines)

		if *printFrequency {
			wanalyzer.PrintTimeFrequence(timeFrequences)
		} else {
			drawer := wanalyzer.NewGraph("timefrequences.png")
			drawer.DrawTimeFrequence(timeFrequences)

			fmt.Println("Graph of messaging frequency in time periods was created.")
		}
	} else if len(*file) > 0 {

		words := wanalyzer.GetWordsWithOrder(lines)
		wanalyzer.PrintWords(words, *start, *limit)
	}
}
