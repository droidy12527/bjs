package parser

import (
	"compiler/ast"
	"compiler/lexer"
	"fmt"
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

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkforErrors(p, t)

	l := len(program.Statements)
	if l != 1 {
		t.Fatalf("program has not enough statements. got=%d", l)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	testIntegerLiteral(t, stmt.Expression, 5)
}

func TestPrefixOperationTesting(t *testing.T) {
	input := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}
	for _, tt := range input {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkforErrors(p, t)
		if len(program.Statements) != 1 {
			t.Fatalf("Program must contain %d statements got %d", 1, len(program.Statements))
		}
		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program statement must be ast.ExpressionStatement got back %T", statement)
		}
		prefix, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expected to get ast.PrefixExpression git back %T", statement.Expression)
		}
		if prefix.Operator != tt.operator {
			t.Fatalf("Operator is wrong expected %s got back %s", tt.operator, prefix.Operator)
		}
		testIntegerLiteral(t, prefix.Right, tt.integerValue)
	}
}

func TestOperatorPrecedingParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkforErrors(p, t)
		expected := program.String()
		if expected != tt.expected {
			t.Errorf("expected %q got %q", tt.expected, expected)
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range tests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkforErrors(p, t)
		l := len(program.Statements)
		if l != 1 {
			t.Errorf("program.Statements does not contain %d statements. got=%d", 1, l)
			continue
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
				program.Statements[0])
			continue
		}
		expr, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Errorf("expr not *ast.InfixExpression. got=%T", stmt.Expression)
			continue
		}
		testIntegerLiteral(t, expr.Left, tt.leftValue)
		if expr.Operator != tt.operator {
			t.Errorf("expr.Operator is not %q. got=%s", tt.operator, expr.Operator)
		}
		testIntegerLiteral(t, expr.Right, tt.rightValue)
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	i, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if i.Value != value {
		t.Errorf("i.Value not %d. got=%d", value, i.Value)
		return false
	}

	if i.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("i.TokenLiteral() not %d. got=%s", value, i.TokenLiteral())
		return false
	}
	return true
}

func checkforErrors(p *Parser, t *testing.T) {
	errors := p.Errors()
	for _, errors := range errors {
		t.Errorf("Errors are as follows %s", errors)
	}
}

// func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
// 	ident, ok := exp.(*ast.Identifier)
// 	if !ok {
// 		t.Errorf("Expression is not identifer got = %T", exp)
// 		return false
// 	}
// 	if ident.Value != value {
// 		t.Errorf("Identifier value is wrong , expected %s, and got %s", value, ident.Value)
// 		return false
// 	}
// 	if ident.TokenLiteral() != value {
// 		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
// 		return false
// 	}
// 	return true
// }

// func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
// 	switch v := expected.(type) {
// 	case int:
// 		return testIntegerLiteral(t, exp, int64(v))
// 	case int64:
// 		return testIntegerLiteral(t, exp, v)
// 	case string:
// 		return testIdentifier(t, exp, v)
// 	}
// 	t.Errorf("Expression type is not handled yet, got := %T", exp)
// 	return false
// }

// func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
// 	okExpression, ok := exp.(*ast.InfixExpression)
// 	if !ok {
// 		t.Errorf("Expression is not ast Expression got %T instead", exp)
// 		return false
// 	}
// 	if !testLiteralExpression(t, okExpression.Left, left) {
// 		return false
// 	}
// 	if okExpression.Operator != operator {
// 		t.Errorf("exp.Operator is not '%s'. got=%q", operator, okExpression.Operator)
// 		return false
// 	}
// 	if !testLiteralExpression(t, okExpression.Right, right) {
// 		return false
// 	}
// 	return true
// }

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
