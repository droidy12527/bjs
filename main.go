package main

import (
	"os"
	"strings"
)

/*
	Use pointer parsing for getting the keywords, Makes thing more easy.
	While loading the program, Make sure to dynamically load it rather than loading it through consuming heap memory.
	Use lexical stack for loading the text
*/

func main() {
	readFile()
}

func readFile() {
	osArgs := os.Args
	if len(osArgs) < 1 {
		panic("not enough arguments to compile file")
	}
	filename := osArgs[1]
	fileExtension := strings.Split(filename, ".")[1]
	if fileExtension != "pookie" {
		panic("file extension is wrong, make sure you are using pookie file extension.")
	}
}
