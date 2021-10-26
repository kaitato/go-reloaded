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
	File = gr.ToLow(File)
	File = gr.ToCap(File)
	File = gr.Punctuation(File)
	File = gr.QuoMarks(File)
	File = gr.ToA(File)
	fmt.Println(File)
}
