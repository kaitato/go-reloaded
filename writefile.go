package gr

import (
	"log"
	"os"
)

func WriteFile(s string) {
	File2 := os.Args[2]
	f, err := os.Create(File2)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	data := []byte(s)

	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}

}
