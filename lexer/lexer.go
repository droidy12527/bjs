package lexer

import "compiler/tokens"

/*
	Lexer Interface has been defined
	Defined Functions:
		1. readSingleChar(): reads single char from the given input code.
*/

type Lexer interface {
	nextToken() tokens.Token
}

type lexer struct {
	input              string
	currentCharPointer int
	nextCharPointer    int
	char               byte
}

func New(input []byte) Lexer {
	localLexer := &lexer{
		input: string(input),
	}
	localLexer.readSingleChar()
	return localLexer
}

func (l *lexer) readSingleChar() {
	inputLen := len(l.input)
	if l.nextCharPointer >= inputLen {
		l.char = 0
	} else {
		l.char = l.input[l.nextCharPointer]
	}
	l.currentCharPointer = l.nextCharPointer
	l.nextCharPointer++
}

func (l *lexer) nextToken() tokens.Token {
	var tok tokens.Token
	return tok
}
