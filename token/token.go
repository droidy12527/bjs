package token

/*
	Token Configuration for language.
	This file contains all of the tokens necessary for the language to work.
	More Tokens will be added in the future for reference.
	Things to understand here are:
		1. A new type is created here of string so that there is no mismatch between the internal string of golang and the token created for our pookie-lang.
		2. All the tokens are converted into Hashmap for better finding of tokens when necessary.
		3. Function lookup for identifier returns type identifier if not found in the basic tokenization.
*/

type Type string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	FLOAT     = "FLOAT"
	STRING    = "STRING"
	BANG      = "!"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTARISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	LE        = "<="
	GE        = ">="
	EQ        = "=="
	NEQ       = "!="
	AND       = "&&"
	OR        = "||"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	NIL       = "NIL"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	MACRO     = "MACRO"
)

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"nil":    NIL,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func ReadIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
