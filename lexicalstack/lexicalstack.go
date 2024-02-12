package lexicalstack

import (
	"compiler/constants"
	"compiler/identifiers"
	"fmt"
)

/*
	Using heap memory for parsing the language with less complexity for pushing and popping.
	Gives access to free memory rather than allocating a large amount of memory at start for parsing the code.

	TODO:
	1. Memory leaks and issues to be kept in mind.
	2. Add the keyword specifier in the map itself
*/

type StackNode struct {
	OperatorType string	
	Operator     string
	NextNode     *StackNode
}

type LexicalStack struct {
	StackHead         *StackNode
	LexicalStackDepth int
}

type asciiMetadata struct {
	Type   string
	String string
}

/*
	Helper Functions
*/

func getOperatorType(asciiString string) asciiMetadata {
	_, ok := identifiers.GetKeywords()[asciiString]
	if ok {
		return asciiMetadata{
			Type:   constants.KEYWORD,
			String: asciiString,
		}
	}
	_, ok = identifiers.GetOperators()[asciiString]
	if ok {
		return asciiMetadata{
			Type:   constants.OPERATOR,
			String: asciiString,
		}
	}
	_, ok = identifiers.GetSeperators()[asciiString]
	if ok {
		return asciiMetadata{
			Type:   constants.SEPERATOR,
			String: asciiString,
		}
	}
	return asciiMetadata{
		Type:   constants.IDENTFIER,
		String: asciiString,
	}
}

func recStacktrace(stacknode *StackNode) {
	fmt.Printf("%+v \n", stacknode)
	if stacknode.NextNode == nil {
		return
	}
	recStacktrace(stacknode.NextNode)
}

/*
	Pointer functions for the main code.
*/

func Init() *LexicalStack {
	return &LexicalStack{
		StackHead:         nil,
		LexicalStackDepth: 0,
	}
}

func (ls *LexicalStack) ParseASCII(asciiString string) {
	// Return for blank spaces as we don't need stack for parsing empty ascii nodes
	if asciiString == " " || asciiString == "" {
		return
	}
	// Check for the first node of the stack
	parsedOperator := getOperatorType(asciiString)
	if ls.StackHead == nil && ls.LexicalStackDepth == 0 {
		node := &StackNode{
			OperatorType: parsedOperator.Type,
			Operator:     parsedOperator.String,
			NextNode:     nil,
		}
		ls.StackHead = node
	} else {
		prevNode := ls.StackHead
		node := &StackNode{
			OperatorType: parsedOperator.Type,
			Operator:     parsedOperator.String,
			NextNode:     prevNode,
		}
		ls.StackHead = node
	}
	ls.LexicalStackDepth = ls.LexicalStackDepth + 1
}

func (ls *LexicalStack) PrintStackTrace() {
	recStacktrace(ls.StackHead)
}

func (ls *LexicalStack) AnalyseTuringMachine() {

}
