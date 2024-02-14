package parser

import (
	"compiler/lexer"
	"compiler/token"
)

type (
	prefixParseFunction func() ast.
)


type Parser struct {
	l lexer.Lexer
	errors []string
	currToken token.Token
	peekToken token.Token
	prefixAST map[token.Type]prefixParseFunction
	infixAST map[token.Token] infixParseFunction
}

