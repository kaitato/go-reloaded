package gr

import "strings"

func QuoMarks(s string) string {
	counter := 0
	var result string
	words := strings.Split(s, " ")
	for a := 0; a < len(words); a++ {
		if a == 0 {
			result += words[a]
		} else if a == len(words) && words[a] != "'" {
			result += " " + words[a]
		} else if words[a] != "'" {
			result += " " + words[a]
		} else if words[a] == "'" {
			counter++
			if counter%2 == 1 {
				result += " " + words[a] + words[a+1]
				a++
			} else {
				result += words[a]
			}
		}
	}
	return result
}
