package ast

import (
	"compiler/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "pookie"},
					Value: "pookie",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "rose"},
					Value: "rose",
				},
			},
		},
	}

	if program.String() != "let pookie = rose;" {
		t.Errorf("program.String() wrong. got=%T", program.String())
	}
}
