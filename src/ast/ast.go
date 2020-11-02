package ast

import (
	"bytes"
	"strings"

	"github.com/ythosa/pukiclang/src/token"
)

// Node is interface for simple node (anyone) element in the AST tree
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is interface for statements elements in the AST tree
type Statement interface {
	Node
	statementNode()
}

// Expression is interface for expressions elements in the AST tree
type Expression interface {
	Node
	expressionNode()
}

// Program is type for program - higher element of AST tree
type Program struct {
	Statements []Statement
}

// TokenLiteral returns token literal of the node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// String returns string representation of the node
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement is type for let statements in the AST tree
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns token literal of the node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns string representation of the node
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

// ReturnStatement is type for return statements in the AST tree
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns token literal of the node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns string representation of the node
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is type for expression statements in the AST tree
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns token literal of the node
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns string representation of the node
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Identifier is type for identifier expressions in the AST tree
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns token literal of the node
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns string representation of the node
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral is type for integer expressions in the AST tree
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns token literal of the node
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns string representation of the node
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression is type for prefix expressions in the AST tree
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns token literal of the node
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns string representation of the node
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is type for infix expressions in the AST tree
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns token literal of the node
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns string representation of the node
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean is type for boolean expressions in the AST tree
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns token literal of the node
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String returns string representation of the node
func (b *Boolean) String() string {
	return b.Token.Literal
}

// StringLiteral is type for string literals
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns token literal of the node
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// String returns string representation of the node
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

// IfExpression is type for if expressions in the AST tree
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns token literal of the node
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns string representation of the node
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

// BlockStatement is type for block statements in the AST tree
type BlockStatement struct {
	Token      token.Token // The '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns token literal of the node
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String returns string representation of the node
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// FunctionLiteral is type for function literals in the AST tree
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns token literal of the node
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String returns string representation of the node
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression is type for call expressions in the AST tree
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns token literal of the node
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String returns string representation of the node
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// ArrayLiteral is type for array literals
type ArrayLiteral struct {
	Token    token.Token // The '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral returns token literal of the node
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

// String returns string representation of the node
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// IndexExpression is type for `<expression>[<expression>]` index expressions
type IndexExpression struct {
	Token token.Token // The '[' token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral returns token literal of the node
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns string representation of the node
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
