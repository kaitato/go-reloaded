package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1]
	result := os.Args[2]

	contents, err := ioutil.ReadFile(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(contents), "\n")

	for i, line := range lines {
		if strings.Contains(line, "(cap)") {
			lines[i] = "LOL"
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
