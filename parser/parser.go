package parser

import (
	"compiler/ast"
	"compiler/constants"
	"compiler/lexer"
	"compiler/token"
	"fmt"
)

/*
	Refactor functions such as expectPeek and peekTokenIs for better code simplicity
	Refer to PRATT PARSER: https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/
	Used Pratt parsing, Which is top down precedence parsing.
*/

type (
	prefixParsingFunction func() ast.Expression
	infixParsingFunction  func(ast.Expression) ast.Expression
)

type Parser struct {
	l                     lexer.Lexer
	curToken              token.Token
	peekToken             token.Token
	errors                []string
	prefixParsingFunction map[token.Type]prefixParsingFunction
	infixParsingFunction  map[token.Type]infixParsingFunction
}

func New(l lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.prefixParsingFunction = map[token.Type]prefixParsingFunction{
		token.IDENT: p.parseIdentifier,
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) registerInfixParsingMap(tokenType token.Type, fn infixParsingFunction) {
	p.infixParsingFunction[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: p.parseExpression(constants.LOWEST),
	}
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParsingFunction[p.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.curToken.Type)
		p.errors = append(p.errors, msg)
		return nil
	}
	expr := prefix()
	return expr
}

func (p *Parser) parseReturnStatement() ast.Statement {
	statement := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseLetStatement() ast.Statement {
	statement := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// Skip till we see a semicolon for the statement end.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	message := fmt.Sprintf("Expected next token is %s we got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, message)
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
