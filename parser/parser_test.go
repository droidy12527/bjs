package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
		let x = 1;
		let y = 2;
		let foobar = 10;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Program returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Program does not contain 3 statements got %d in return", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.Tolenliteral not 'let' got back %q", s.TokenLiteral())
		return false
	}
	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.Letstatement got %T", s)
		return false
	}
	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value is '%s' got %s", name, letStatement.Name.Value)
		return false
	}
	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name not %s. got %s", name, letStatement.Name)
		return false
	}
	return true
}
