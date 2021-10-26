package gr

import "strings"

func Punctuation(s string) string {
	var result string
	words := strings.Split(s, " ")
	for a := 0; a < len(words); a++ {
		if a == 0 {
			result += words[a]
		} else if !strings.HasPrefix(words[a], ",") && !strings.HasPrefix(words[a], ".") && !strings.HasPrefix(words[a], "!") && !strings.HasPrefix(words[a], "?") && !strings.HasPrefix(words[a], ";") && !strings.HasPrefix(words[a], ":") {
			result += " " + words[a]
		} else if words[a] == "," || words[a] == "." || words[a] == "!" || words[a] == "?" || words[a] == ":" || words[a] == ";" || words[a] == "..." || words[a] == "!?" {
			result += words[a]
		} else if strings.HasPrefix(words[a], ",") || strings.HasPrefix(words[a], ".") || strings.HasPrefix(words[a], "!") || strings.HasPrefix(words[a], "?") || strings.HasPrefix(words[a], ";") || strings.HasPrefix(words[a], ":") {
			for _, c := range words[a] {
				switch c {
				case '.', ',', '!', '?', ':', ';':
					result += string(c) + " "
				}
				if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
					result += string(c)
				}
			}
		}
	}
	return result
}
