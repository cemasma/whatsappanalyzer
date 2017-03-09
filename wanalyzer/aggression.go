package wanalyzer

func CalculateAggression(lines, negatives []string) map[string]int {
	usernames := GetUsernames(lines)

	aggression := make(map[string]int)

	for _, username := range usernames {
		aggression[username] = AggressionCount(GetUserLines(lines, username), negatives)
	}

	aggression["total"] = AggressionCount(lines, negatives)
	return aggression
}

func AggressionCount(lines, negatives []string) (aggression int) {
	for _, word := range negatives {
		aggression += CountWordInLines(word, lines)
	}

	return
}
