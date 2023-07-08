package parsing

import (
	"fmt"
	"github.com/mattellis91/glox/pkg/lexing"
)

type Expression interface {
	ToString() string
	accept(v visitor)	
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

type visitor interface {
	visitForBinaryExpression(be *BinaryExpression)
	visitForGroupingExpression(bg *GroupingExpression)
	visitForLiteralExpression(le *LiteralExpression)
	visitForUnaryExpression(ue *UnaryExpression)
}

func NewBinaryExpression(left Expression, operator lexing.Token, right Expression) *BinaryExpression {
	return &BinaryExpression{
		left: left,
		operator: operator,
		right: right,
	}	
}

func (be *BinaryExpression) accept(v visitor) {
	v.visitForBinaryExpression(be)
}

func NewGroupingExpression(expression Expression) *GroupingExpression {
	return &GroupingExpression{
		expression: expression,
	}
}

func (ge *GroupingExpression) accept(v visitor) {
	v.visitForGroupingExpression(ge)
}

func NewLiteralExpression(value any) *LiteralExpression {
	return &LiteralExpression{
		value: value,
	}
}

func (le *LiteralExpression) accept(v visitor) {
	v.visitForLiteralExpression(le)
}

func NewUnaryExpression(operator lexing.Token, right Expression) *UnaryExpression {
	return &UnaryExpression{
		operator: operator,
		right: right,
	}
} 

func (ue *UnaryExpression) accept(v visitor) {
	v.visitForUnaryExpression(ue)
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