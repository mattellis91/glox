package parsing

import (
	"fmt"
	"github.com/mattellis91/zima/pkg/lexing"
)

type Expression interface {
	ToString() string	
}

type BinaryExpression struct {
	 left Expression
	 operator lexing.Token
	 right Expression 
}

type GroupingExpression struct {
	expression Expression
}

type LiteralExpression struct {
	value any
}

type UnaryExpression struct {
	operator lexing.Token
	right Expression 
}

func NewBinaryExpression(left Expression, operator lexing.Token, right Expression) *BinaryExpression {
	return &BinaryExpression{
		left: left,
		operator: operator,
		right: right,
	}	
}

func NewGroupingExpression(expression Expression) *GroupingExpression {
	return &GroupingExpression{
		expression: expression,
	}
}

func NewLiteralExpression(value any) *LiteralExpression {
	return &LiteralExpression{
		value: value,
	}
}

func NewUnaryExpression(operator lexing.Token, right Expression) *UnaryExpression {
	return &UnaryExpression{
		operator: operator,
		right: right,
	}
} 

func (be *BinaryExpression) ToString() string {
	return fmt.Sprintf("Binary Expression: %s | %s | %s", be.left.ToString(), be.operator.ToString(), be.right.ToString())
}

func (ge *GroupingExpression) ToString() string {
	return fmt.Sprintf("Grouping Expression: %s", ge.expression.ToString())
}

func (le *LiteralExpression) ToString() string {
	return fmt.Sprintf("Literal Expression: %v", le.value)
}

func (ue *UnaryExpression) ToString() string {
	return fmt.Sprintf("Unary Expression: %s | %s", ue.operator.ToString(), ue.right.ToString())
}