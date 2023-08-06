package wanalyzer

import (
	"fmt"
	"regexp"
	"strings"
)

// MessageFrequence a struct for storing message count per day
type MessageFrequence struct {
	Date  string
	Count int
}

func sortFrequency(frequence []MessageFrequence) []MessageFrequence {
	for i := 0; i < len(frequence); i++ {
		for k := i + 1; k < len(frequence); k++ {
			if frequence[k].Count > frequence[i].Count {
				backup := frequence[i]
				frequence[i] = frequence[k]
				frequence[k] = backup
			}
		}
	}
	return frequence
}

// GetMessageFrequency provides messages' frequencies
func GetMessageFrequency(lines, dates []string) (frequency []MessageFrequence) {
	for _, date := range dates {
		count := getMessageCount(lines, date)
		frequency = append(frequency, MessageFrequence{Date: date, Count: count})
	}

	return frequency
}

// GetDatesFromLines parses dates from the lines and returns
func GetDatesFromLines(lines []string) map[string]int {
	re := regexp.MustCompile(`\[(.*?)]`)

	dateMap := make(map[string]int, 0)

	for _, line := range lines {
		splitDate := strings.Split(re.FindString(line), " ")[0]
		date := strings.TrimLeft(splitDate, "[")

		if val, ok := dateMap[date]; ok {
			dateMap[date] = val + 1
		} else {
			dateMap[date] = 1
		}
	}

	return dateMap
}

// PrintMessageFrequency prints messaging the frequency
func PrintMessageFrequency(dates map[string]int, start, end int) {
	frequencyArray := make([]MessageFrequence, 0)
	for key, value := range dates {
		frequencyArray = append(frequencyArray, MessageFrequence{
			Date:  key,
			Count: value,
		})
	}

	frequencyArray = sortFrequency(frequencyArray)

	for i := start; i < end; i++ {
		fmt.Printf("%d. %s \t %d\n", i+1, frequencyArray[i].Date, frequencyArray[i].Count)
	}
}

// GetTimeFrequency prepares and provides messaging frequency by time periods
func GetTimeFrequency(lines []string) (timeFrequence map[string]int) {
	timeFrequence = map[string]int{
		"Morning":   0,
		"Noon":      0,
		"Afternoon": 0,
		"Evening":   0,
		"Night":     0,
	}

	for _, line := range lines {
		hourTime := getHourTime(parseHour(line))
		if _, ok := timeFrequence[hourTime]; ok {
			timeFrequence[hourTime]++
		}
	}

	return
}

// PrintTimeFrequency prints the frequencies
func PrintTimeFrequency(frequencyMap map[string]int) {
	for key, value := range frequencyMap {
		fmt.Printf("%s message count: %d\n", key, value)
	}
}
