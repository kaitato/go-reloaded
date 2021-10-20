package gr

import (
	"strings"
)

func ToHex() {
	var result string
	words := strings.SplitAfter(ReadFile(), " ")
	for i, a := range words {
		if a != "(hex)" {
			result += a
		} else if a[i+1] == "(hex)" {
		}
	}
}
