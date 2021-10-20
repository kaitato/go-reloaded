package gr

import (
	"fmt"
	"os"
)

func ReadFile() string {
	File := os.Args[1]

	contents, err := os.ReadFile(File)
	if err != nil {
		fmt.Print("Error!")
	} else {
		return string(contents)
	}
	return ""
}
