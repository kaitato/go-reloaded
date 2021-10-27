package gr

import "os"

func WriteFile(s string) {
	os.WriteFile("result.txt", s)
}
