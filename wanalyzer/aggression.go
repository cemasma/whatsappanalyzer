package wanalyzer

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// CalculateAggression calculates aggression by use counts of negative words
func CalculateAggression(lines []string) map[string]float64 {
	usernames := GetUsernames(lines)

	aggression := make(map[string]float64)

	for _, username := range usernames {
		aggression[username] = AggressionCount(GetUserLines(lines, username))
	}

	aggression["total"] = AggressionCount(lines)
	return aggression
}

// AggressionCount collect points for every negative word
func AggressionCount(lines []string) (aggression float64) {

	for _, line := range lines {
		splittedLine := strings.Split(line, ":")

		if len(splittedLine) > 2 {

			sentence := splittedLine[2]

			cmd := exec.Command("python", "analysis_scripts/analysis.py", "--sentence=\""+sentence+"\"", "--type=sentiment")

			out, err := cmd.Output()

			if err != nil {
				println(err)
				return
			}

			strOutput := string(out)

			if strings.Contains(strOutput, "'negative'") == true {
				score, _ := getScore(strOutput, "negative")
				aggression += score
			} else if strings.Contains(strOutput, "'positive'") == true {
				score, _ := getScore(strOutput, "positive")
				aggression -= score
			}
			fmt.Println(sentence)
			fmt.Println(aggression)
		}
	}

	return
}

func getScore(text, token string) (score float64, err error) {
	splittedText := strings.Split(text, "classification: [['"+token+"', ")[1]
	strScore := strings.Split(splittedText, "]]")[0]
	score, err = strconv.ParseFloat(strScore, 64)

	return
}
