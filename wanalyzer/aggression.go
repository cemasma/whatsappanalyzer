package wanalyzer

// CalculateAggression calculates aggression by use counts of negative words
func CalculateAggression(lines, negatives []string) map[string]int {
	usernames := GetUsernames(lines)

	aggression := make(map[string]int)

	for _, username := range usernames {
		aggression[username] = AggressionCount(GetUserLines(lines, username), negatives)
	}

	aggression["total"] = AggressionCount(lines, negatives)
	return aggression
}

// AggressionCount collect points for every negative word
func AggressionCount(lines, negatives []string) (aggression int) {
	for _, word := range negatives {
		aggression += CountWordInLines(word, lines)
	}

	return
}
