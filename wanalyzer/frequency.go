package wanalyzer

import "fmt"

// MessageFrequence a struct for storing message count per day
type MessageFrequence struct {
	Date  string
	Count int
}

// SortFrequency sorts the frequency array by message count
func SortFrequency(frequence []MessageFrequence) []MessageFrequence {
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

// GetDatesFromLines separates the dates from lines and returns these in array
func GetDatesFromLines(lines []string) (dates []string) {
	for _, line := range lines {
		if len(line) > 10 && (string(line[2]) == "." && string(line[5]) == "." && string(line[10]) == ",") &&
			string(line[14]) == ":" && !Contains(dates, string(line[:10])) {

			dates = append(dates, line[:10])
		} else if len(line) > 9 && string(line[1]) == "." && string(line[4]) == "." && string(line[9]) == "," &&
			string(line[13]) == ":" && !Contains(dates, string(line[:9])) {

			dates = append(dates, line[:9])
		}
	}
	return
}

// PrintMessageFrequence prints the frequencies
func PrintMessageFrequence(frequence []MessageFrequence, start, end int) {
	for i := start; i < end; i++ {
		fmt.Printf("%d. %s \t %d\n", i+1, frequence[i].Date, frequence[i].Count)
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
		hourTime := getHourTime(getHourInLine(line))
		if _, ok := timeFrequence[hourTime]; ok {
			timeFrequence[hourTime]++
		}
	}

	return
}

// PrintTimeFrequence prints the frequencies
func PrintTimeFrequence(frequence map[string]int) {
	for key, value := range frequence {
		fmt.Printf("%s message count: %d\n", key, value)
	}
}
