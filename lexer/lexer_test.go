package lexer

import (
	"compiler/tokens"
	"testing"
)

func TestNexttokens(t *testing.T) {
	input := []byte(`
	let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	!-/*0;
	2 < 10 > 7;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 <= 11;
	10 >= 9;
	10 == 10;
	10 != 9;

	true && false;
	true || false;

	"foobar";
	"foo bar";

	[1, 2];

	{"foo": "bar"};

	# comment
	let a = 1; # inline comment

	let b = 123.45;
	let c = 0.678;
	let d = 9.0;

	a = 2;
	b = nil;

	macro(x, y) { x + y; };
	`)

	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.LET, "let"},
		{tokens.IDENT, "five"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "ten"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "add"},
		{tokens.ASSIGN, "="},
		{tokens.FUNCTION, "fn"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "x"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "y"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.IDENT, "x"},
		{tokens.PLUS, "+"},
		{tokens.IDENT, "y"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "result"},
		{tokens.ASSIGN, "="},
		{tokens.IDENT, "add"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "five"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "ten"},
		{tokens.RPAREN, ")"},
		{tokens.SEMICOLON, ";"},
		{tokens.BANG, "!"},
		{tokens.MINUS, "-"},
		{tokens.SLASH, "/"},
		{tokens.ASTARISK, "*"},
		{tokens.INT, "0"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "2"},
		{tokens.LT, "<"},
		{tokens.INT, "10"},
		{tokens.GT, ">"},
		{tokens.INT, "7"},
		{tokens.SEMICOLON, ";"},
		{tokens.IF, "if"},
		{tokens.LPAREN, "("},
		{tokens.INT, "5"},
		{tokens.LT, "<"},
		{tokens.INT, "10"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.RETURN, "return"},
		{tokens.TRUE, "true"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.ELSE, "else"},
		{tokens.LBRACE, "{"},
		{tokens.RETURN, "return"},
		{tokens.FALSE, "false"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.INT, "10"},
		{tokens.LE, "<="},
		{tokens.INT, "11"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "10"},
		{tokens.GE, ">="},
		{tokens.INT, "9"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "10"},
		{tokens.EQ, "=="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},
		{tokens.INT, "10"},
		{tokens.NEQ, "!="},
		{tokens.INT, "9"},
		{tokens.SEMICOLON, ";"},
		{tokens.TRUE, "true"},
		{tokens.AND, "&&"},
		{tokens.FALSE, "false"},
		{tokens.SEMICOLON, ";"},
		{tokens.TRUE, "true"},
		{tokens.OR, "||"},
		{tokens.FALSE, "false"},
		{tokens.SEMICOLON, ";"},
		{tokens.STRING, "foobar"},
		{tokens.SEMICOLON, ";"},
		{tokens.STRING, "foo bar"},
		{tokens.SEMICOLON, ";"},
		{tokens.LBRACKET, "["},
		{tokens.INT, "1"},
		{tokens.COMMA, ","},
		{tokens.INT, "2"},
		{tokens.RBRACKET, "]"},
		{tokens.SEMICOLON, ";"},
		{tokens.LBRACE, "{"},
		{tokens.STRING, "foo"},
		{tokens.COLON, ":"},
		{tokens.STRING, "bar"},
		{tokens.RBRACE, "}"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "a"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "1"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "b"},
		{tokens.ASSIGN, "="},
		{tokens.FLOAT, "123.45"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "c"},
		{tokens.ASSIGN, "="},
		{tokens.FLOAT, "0.678"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.IDENT, "d"},
		{tokens.ASSIGN, "="},
		{tokens.FLOAT, "9.0"},
		{tokens.SEMICOLON, ";"},
		{tokens.IDENT, "a"},
		{tokens.ASSIGN, "="},
		{tokens.INT, "2"},
		{tokens.SEMICOLON, ";"},
		{tokens.IDENT, "b"},
		{tokens.ASSIGN, "="},
		{tokens.NIL, "nil"},
		{tokens.SEMICOLON, ";"},
		{tokens.MACRO, "macro"},
		{tokens.LPAREN, "("},
		{tokens.IDENT, "x"},
		{tokens.COMMA, ","},
		{tokens.IDENT, "y"},
		{tokens.RPAREN, ")"},
		{tokens.LBRACE, "{"},
		{tokens.IDENT, "x"},
		{tokens.PLUS, "+"},
		{tokens.IDENT, "y"},
		{tokens.SEMICOLON, ";"},
		{tokens.RBRACE, "}"},
		{tokens.SEMICOLON, ";"},
		{tokens.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.nextToken()

		if tok.Type != tt.expectedType {
			t.Logf("tests[%d] - tok: %#v", i, tok)
			t.Fatalf("tests[%d] - tokens type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Logf("tests[%d] - tok: %#v", i, tok)
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
