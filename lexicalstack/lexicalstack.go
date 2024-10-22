package lexicalstack

import (
	"compiler/constants"
	"compiler/identifiers"
	"fmt"
)

// TODO: Remove this code block as this is not longer needed
// Lexical stack is not in use

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
