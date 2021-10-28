package gr

<<<<<<< HEAD
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

=======
import "os"

func WriteFile(s string) {
	os.WriteFile("result.txt", s)
>>>>>>> fb5f9ff10a8ccd2ba48d534c07c74c3e617170b9
}
