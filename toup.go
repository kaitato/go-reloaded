package gr

import (
	"strings"
)

func ToUp(s string) string {
	var result string
	words := strings.Split(s, " ")
	for a := 0; a < len(words); a++ {
		if a == len(words)-1 {
			result += " " + words[a]
		} else if !strings.Contains(words[a+1], "(up") {
			result += " " + words[a]
		} else if words[a+1] == "(up)" {
			words[a] = strings.ToUpper(words[a])
			result += " " + words[a]
			a += 1
		} else if words[a+1] == "(up," {
			str := []rune(words[a+2])
			pram := 0
			for i := 0; i < len(str); i++ {
				if str[i] >= '0' && str[i] <= '9' {
					pram *= 10
					pram += int(str[i] - 48)
				}
			}
			for b := 0; b <= pram; b++ {
				words[a-b] = strings.ToUpper(words[a-b])
			}
		}
	}
	return result
}
