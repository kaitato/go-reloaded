package gr

import (
	"strconv"
	"strings"
)

func ToBin(s string) string {
	var result string
	words := strings.Split(s, " ")
	for a := 0; a < len(words); a++ {
		if a == len(words)-1 {
			result += " " + words[a]
		} else if words[a+1] != "(bin)" {
			result += " " + words[a]
		} else if words[a+1] == "(bin)" {
			var stringa string

			numberStr := strings.Replace(words[a], stringa, "", -1)
			numberStr = strings.Replace(numberStr, stringa, "", -1)
			n, _ := strconv.ParseInt(numberStr, 2, 64)
			m := strconv.FormatInt(n, 10)
			result = result + " " + m
			a += 1
		}
	}
	return result
}
