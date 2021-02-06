package ast

import "zlang/token"

// Node requires every node to provide a token literal
type Node interface {
	TokenLiteral() string
}

// Statement is a statement node
type Statement interface {
	Node
	statementNode()
}

// Expression is an expression node
type Expression interface {
	Node
	expressionNode()
}

// Program node is the root node of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral for Program node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is a let statement node (let a = 1;)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns token literal for let statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is an identifier node
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns a token literal for identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ReturnStatement is a return statement node
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns a token literal for return statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
