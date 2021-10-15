package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	File := strings.Split(string(ReadFile()), " ")
	for i, char := range File {
		if strings.Contains(char, "(hex)") {
			File[i] = ""
		}
	}

	output := strings.Join(lines, "\n")

	err = ioutil.WriteFile(result, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error creating", result)
		fmt.Println(err)
		return
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

func toHex() {
}
