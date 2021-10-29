package gr

import "strings"

func QuoMarks(s string) string {
	counter := 0
	var result string
	words := strings.Split(s, " ")
	for a := len(words) - 1; a >= 0; a-- {
		if a == 0 {
			result = words[a] + " " + result
		} else if a == len(words)-1 {
			result = words[a]
		} else if words[a] != "'" {
			result = words[a] + " " + result
		} else if words[a] == "'" {
			counter++
			if counter%2 == 1 {
				result = words[a-1] + words[a] + result
				a--
			} else {
				result = " " + words[a] + result
			}
		}
	}
	return result
}
