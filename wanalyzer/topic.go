package wanalyzer

import (
	"fmt"
	"os/exec"
	"strings"
	"unicode"
)

//ObserverTopics provides topics for lines
func ObserverTopics(lines []string) map[string]int {
	topicMap := make(map[string]int)

	for _, line := range lines {
		splittedLine := strings.Split(line, ":")

		if len(splittedLine) > 2 {

			sentence := splittedLine[2]

			sentence = strings.ToLowerSpecial(unicode.TurkishCase, sentence)

			cmd := exec.Command("python", "analysis_scripts/analysis.py", "--sentence=\""+sentence+"\"", "--type=topic")

			out, err := cmd.Output()

			if err != nil {
				println(err)
				return nil
			}

			strOutput := string(out)

			fmt.Println(strOutput)

			topics := GetTopicsFromResult(strOutput)

			for _, topic := range topics {
				topicMap[topic]++
			}
		}
	}

	// fmt.Println(topicMap)

	return topicMap
}
