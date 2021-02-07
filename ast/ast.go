package ast

import (
	"bytes"
	"zlang/token"
)

// Node requires every node to provide a token literal
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns String() from each node in Program
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String stringifies LetStatement node
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// ReturnStatement is a return statement node
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns a token literal for return statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String stringifies a return statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is an expression statement node
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns a token literal for expression statement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String stringifies an expression statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier is an identifier node
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns a token literal for identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String returns identifier value
func (i *Identifier) String() string { return i.Value }

// IntegerLiteral is an integer literal node
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns a token literal for integer literal
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String stringifies an integer literal
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// PrefixExpression is a prefix expression node
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns a token literal for prefix expression
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String stringifies a prefix expression
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
