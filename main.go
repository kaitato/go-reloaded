package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	words := strings.SplitAfter(ReadFile(), " ")
	for _, a := range words {
		if a == "(hex)" {
			fmt.Print(a)
		}
	}
}

func ReadFile() string {
	file := os.Args[1]

	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Print("Error!")
	} else {
		return string(contents)
	}
	return ""
}

func toHex(s string) string {
	numberStr := strings.Replace(s, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	n, err := strconv.ParseUint(numberStr, 16, 64)
	if err != nil {
		panic(err)
	}
	return string(rune(n))
}
