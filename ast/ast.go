package ast

/*
	Node is an AST node and it should always return a token literal.
	All of the other methods for node should Implement the Node interface.
*/

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
