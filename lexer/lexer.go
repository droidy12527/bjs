package lexer

import "compiler/tokens"

/*
	Lexer Interface defination.
	Defined Functions:
		1. readSingleChar(): reads single char from the given input code.
		2. nextToken(): returns the next token in the char.
		3. skipWhiteSpace(): skips white space until a character is found.
		4. skipComments(): skips the comments in the code.
		5. peekCharacter(): returns the next character in the code.
		6. newToken(tokenType, tokenLiteral): returns new token based on token type and literal.
		7. readTwoCharcters(tokenType): reads two characters and returns a new token.
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
	l.skipWhiteSpace()

	if l.char == '#' {
		l.skipComments()
	}

	return tok
}

func (l *lexer) skipWhiteSpace() {
	for l.char == ' ' || l.char == '\n' || l.char == '\r' || l.char == '\t' {
		l.readSingleChar()
	}
}

func (l *lexer) skipComments() {
	for l.char != '\n' && l.char != '\r' {
		l.readSingleChar()
	}
	l.skipWhiteSpace()
}

func (l *lexer) peekCharacter() byte {
	if len(l.input) > l.nextCharPointer {
		return 0
	}
	return l.input[l.nextCharPointer]
}

func (l *lexer) readTwoCharcters(tokenType tokens.TokenType) tokens.Token {
	currectChar := l.char
	l.readSingleChar()
	return tokens.Token{
		Type:    tokenType,
		Literal: string(currectChar) + string(l.char),
	}
}

func (l *lexer) newToken(tokenType tokens.TokenType, tokenCharacter byte) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: string(tokenCharacter),
	}
}
