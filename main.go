package main

import (
	"compiler/constants"
	"compiler/relp"
	"fmt"
	"os"
	"strings"
)

/*
	Use pointer parsing for getting the keywords, Makes thing more easy.
	While loading the program, Make sure to dynamically load it rather than loading it through consuming heap memory.
	Use lexical stack for loading the text
	Read file for compiling.
*/

func main() {
	// readFile()
	tokenDebug := os.Getenv("BJS_DEBUG")
	tokenCompilationMode := os.Getenv("BJS_COMPILE_MODE")
	if tokenDebug == "true" {
		fmt.Printf(constants.LOGO)
		fmt.Println("JavaScript For Servers, Blazingly Fast and Compiled")
		fmt.Println("Welcome to BJS lang debugger, This is your new RELP for debugging purpose")
		fmt.Println("Please type in expressions to Debug: ")
		if tokenCompilationMode == "true" {
			fmt.Println("Running in Compilation mode")
			relp.StartRELP(os.Stdin, os.Stdout, true)
		} else {
			fmt.Println("Running in Interpreter mode: default mode")
			relp.StartRELP(os.Stdin, os.Stdout, false)
		}
	} else {
		readFile()
	}
}

func readFile() {
	osArgs := os.Args
	if len(osArgs) < 2 {
		panic("not enough arguments to read file")
	}
	filename := osArgs[0]
	fileExtension := strings.Split(filename, ".")[1]
	if fileExtension != "bjs" {
		panic("file extension is wrong, make sure you are using bjs file extension.")
	}
}
