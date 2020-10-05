package ast

import (
	"bytes"
	"strings"

	"github.com/masaaki008/go_interpreter/token"
)

// Node ノード
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement 文
type Statement interface {
	Node
	statementNode()
}

// Expression 式
type Expression interface {
	Node
	expressionNode()
}

// Program nodeインタフェース実装
type Program struct {
	Statements []Statement
}

// TokenLiteral プログラムがもつリテラルを返却
func (p *Program) TokenLiteral() string {
	var result string
	if len(p.Statements) > 0 {
		result = p.Statements[0].TokenLiteral()
	} else {
		result = ""
	}
	return result
}

// String 文字列を返却
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement statementインタフェース実装
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String 文字列を返却
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

// ReturnStatement statementインタフェース実装
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String 文字列を返却
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement statementインタフェース実装
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String 文字列を返却
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement statementインタフェース実装
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

// String 文字列を返却
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Identifier expressionインタフェース実装
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String 文字列を返却
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral expressionインタフェース実装
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String 文字列を返却
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression expressionインタフェース実装
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String 文字列を返却
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression expressionインタフェース実装
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

// String 文字列を返却
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean expressionインタフェース実装
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String 文字列を返却
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression expressionインタフェース実装
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String 文字列を返却
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

// FunctionLiteral expressionインタフェース実装
type FunctionLiteral struct {
	Token     token.Token
	Prameters []*Identifier
	Body      *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

// String 文字列を返却
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Prameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression expressionインタフェース実装
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral トークンがもつリテラルを返却
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// String 文字列を返却
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
