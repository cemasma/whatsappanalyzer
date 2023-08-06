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

	date := flag.String("date", "", "You must specify the date information when you want to observe topics.")

	topic := flag.Bool("topic", false, "If you want to observer topics you must specify the date.\n"+
		"\tExample: analyzer --file \"C:\\filename.txt\" --date \"14.02.2018\" --topic true")

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
				percent := (val * 100) / calculatedAgg["total"]
				fmt.Printf("%s %f percent aggressive by total. \n", key, percent)
				fmt.Printf("%s's using count is %f\n\n", key, val)
			}
		}

		fmt.Printf("The total negative aggression count is: %f", calculatedAgg["total"])
	} else if *topic && len(*date) > 0 {
		lines = wanalyzer.GetLinesByDate(lines, *date)

		topicsMap := wanalyzer.ObserverTopics(lines)

		for key, value := range topicsMap {
			fmt.Println("Topic: ", key, " ", value)
		}
	} else if *messageFrequency {

		dates := wanalyzer.GetDatesFromLines(lines)

		if *printFrequency {
			wanalyzer.PrintMessageFrequency(dates, *start, *limit)
		} else {
			drawer := wanalyzer.NewGraph(*username + " mf.png")
			drawer.DrawFrequency(dates)

			fmt.Println("Graph of messaging frequency was created.")
		}

	} else if *timeFrequency {
		timeFrequencies := wanalyzer.GetTimeFrequency(lines)

		if *printFrequency {
			wanalyzer.PrintTimeFrequency(timeFrequencies)
		} else {
			drawer := wanalyzer.NewGraph("timefrequences.png")
			drawer.DrawFrequency(timeFrequencies)

			fmt.Println("Graph of messaging frequency in time periods was created.")
		}
	} else if len(*file) > 0 {

		words := wanalyzer.GetWordsWithOrder(lines)
		wanalyzer.PrintWords(words, *start, *limit)
	}
}
