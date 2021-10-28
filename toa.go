package gr

import "strings"

func ToA(s string) string {
	var result string
	words := strings.Split(s, " ")
	for a := 0; a < len(words)-1; a++ {
		if a == 0 {
			result += words[a]
		} else if words[a] != "A" && words[a] != "a" {
			result += " " + words[a]
		} else if words[a] == "A" {
			if strings.HasPrefix(words[a+1], "a") || strings.HasPrefix(words[a+1], "e") || strings.HasPrefix(words[a+1], "i") || strings.HasPrefix(words[a+1], "o") || strings.HasPrefix(words[a+1], "u") || strings.HasPrefix(words[a+1], "h") {
				result += " " + "An"
			} else {
				result += " " + words[a]
			}
		} else if words[a] == "a" {
			if strings.HasPrefix(words[a+1], "a") || strings.HasPrefix(words[a+1], "e") || strings.HasPrefix(words[a+1], "i") || strings.HasPrefix(words[a+1], "o") || strings.HasPrefix(words[a+1], "u") || strings.HasPrefix(words[a+1], "h") {
				result += " " + "an"
			} else {
				result += " " + words[a]
			}
		}
	}
	return result
}
