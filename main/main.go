package main

import (
	"fmt"
	"gr"
)

func main() {
	File := gr.ReadFile()
	File = gr.ToHex(File)
	File = gr.ToBin(File)
	File = gr.ToUp(File)
	fmt.Println(File)
}
