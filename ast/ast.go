package ast

import (
	"bytes"
	"strings"
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

// InfixExpression is an infix expression node
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns a token literal for infix expression
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String stringifies an infix expression
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean is a boolean node
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns a token literal for boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String stringifies a boolean
func (b *Boolean) String() string { return b.Token.Literal }

// IfExpression is an if expression node
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns a token literal for if expression
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String stringifies an if expression
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// BlockStatement is a block statement node
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns token literal for block statement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String stringifies a block statement node
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral is a function literal node
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns a token literal for function literal
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// String stringifies a function literal node
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression is a call expression node
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns a token literal for call expression
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

// String stringifies a call expression node
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// StringLiteral is a string node
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns a token literal for string literal
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// String stringifies a string literal
func (sl *StringLiteral) String() string { return sl.Token.Literal }

// ArrayLiteral is an array node
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral returns a token literal for array literal
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// String stringifies an array literal
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// IndexExpression is an index expression node
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral returns a token literal for index expression
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

// String stringifies an index expression
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
