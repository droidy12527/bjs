package ast

import (
	"bytes"
	"compiler/token"
)

/*
	Node is an AST node and it should always return a token literal.
	All of the other methods for node should Implement the Node interface.
	example ast: let i = 0 the AST will be:
		token: token.LET, Name: { token: token.IDEN, value: i }
		value: null
	tokenliteral interface method is used for debugging the program: makes easy to print out strings.
	Prefix statement should be written , For eg Prefix statement AST parsing is 5 * 5 where * is prefix for parsing AST
	String method for debug spits out string concat buffer out for debugging.
*/

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var bytes bytes.Buffer
	for _, buffer := range p.Statements {
		bytes.WriteString(buffer.String())
	}
	return bytes.String()
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var bytes bytes.Buffer
	bytes.WriteString(ls.TokenLiteral() + " ")
	bytes.WriteString(ls.Name.String())
	if ls.Value != nil {
		bytes.WriteString(" = " + ls.Value.String())
	}
	bytes.WriteString(";")
	return bytes.String()
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var bytes bytes.Buffer
	bytes.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		bytes.WriteString(rs.TokenLiteral())
	}
	bytes.WriteString(";")
	return bytes.String()
}

func (e *ExpressionStatement) statementNode()       {}
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }
func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return ""
}
