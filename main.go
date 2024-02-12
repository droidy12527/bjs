package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

/*
	Use pointer parsing for getting the keywords, Makes thing more easy.
*/

func main() {
	// fmt.Println("Hello world")
	// lexicalKeywords := []string{"if", "else", "if", "else", "+", "+", "-", "x", "helloworld", "nice", "}", "{", "}"}
	// lex := lexicalstack.Init()
	// for i := 0; i < len(lexicalKeywords); i++ {
	// 	lex.ParseASCII(lexicalKeywords[i])
	// }
	// lex.PrintStackTrace()
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
	file, err := os.Open(filename)
	if err != nil {
		panic("file not found, missing file.")
	}
	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fmt.Println(words)
	// lexer.Lex(words)
	tokens := token.NewFileSet()
	parsingFile, err := parser.ParseFile(tokens, "./lexicalstack/lexicalstack.go`", nil, parser.Trace)
	if err != nil {
		fmt.Println(err)
	}
	// Create AST (Abstract Syntax Tree) For The Module Declaration and then convert into Bytecode for more flexibilty
	fmt.Println(*parsingFile)
	// for i := 0; i < len(parsingFile.Imports); i++ {
	// 	fmt.Println(*parsingFile.Imports[i])
	// }
	fmt.Println(tokens)
}
