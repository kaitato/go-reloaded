package main

import (
	"fmt"
	"gr"
)

func main() {
	File := gr.ReadFile()
	File = gr.ToHex(File)
	File = gr.ToBin(File)
	fmt.Println(File)
}
