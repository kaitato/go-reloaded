package gr

import (
	"strings"
)

func ToCap(s string) string {
	var result string
	words := strings.Split(s, " ")
	for a := len(words) - 1; a >= 0; a-- {
		if a == 0 {
			result = words[a] + " " + result
		} else if !strings.Contains(words[a], ")") {
			result = words[a] + " " + result
		} else if words[a] == "(cap)" {
			words[a-1] = strings.Title(words[a-1])
			result = words[a-1] + " " + result
			a -= 1
		} else if strings.Contains(words[a], ")") && words[a-1] == "(cap," {
			str := []rune(words[a])
			num := 0
			for i := 0; i < len(str); i++ {
				if str[i] >= '0' && str[i] <= '9' {
					num *= 10
					num += int(str[i] - 48)
				}
			}
			for b := 2; b < num+2; b++ {
				words[a-b] = strings.Title(words[a-b])
				result = words[a-b] + " " + result
			}
			a -= num + 1
		} else if words[a-1] != "(cap," {
			result = words[a] + " " + result
		}
	}
	return result
}
