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
		return "hello";
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkforErrors(p, t)
	if program == nil {
		t.Fatalf("Program returned nil")
	}
	if len(program.Statements) != 4 {
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

func TestIndetifierExpression(t *testing.T) {
	input := "pookie;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkforErrors(p, t)
	if len(program.Statements) != 1 {
		t.Fatalf("program has statement mismatch. got %d stataments", len(program.Statements))
	}
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program statement is not ast expression, got %T", program.Statements[0])
	}
	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expected the value to be identifier got %T", statement.Expression)
	}
	if identifier.Value != "pookie" {
		t.Errorf("identifier value not %s got %s instead", "pookie", identifier.Value)
	}
	if identifier.TokenLiteral() != "pookie" {
		t.Errorf("got wrong token literal for indeitifer, Expected %s got %s", "pookie", identifier.TokenLiteral())
	}
}

func checkforErrors(p *Parser, t *testing.T) {
	errors := p.Errors()
	for _, errors := range errors {
		t.Errorf("Errors are as follows %s", errors)
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
