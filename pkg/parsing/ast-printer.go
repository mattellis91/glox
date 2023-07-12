package parsing

import (
	"fmt"
)

type AstPrinter struct {}

func (astp *AstPrinter) Print(expr Expression) {
	expr.accept(astp)
	fmt.Print("\n")
}

func (astp *AstPrinter) visitForBinaryExpression(be *BinaryExpression) {
	astp.parenthesize(be.operator.Lexeme, be.left, be.right)
}

func (astp *AstPrinter) visitForGroupingExpression(ge *GroupingExpression) {
	astp.parenthesize("Group", ge.expression)
}

func (astp *AstPrinter) visitForLiteralExpression(le *LiteralExpression) {
	if le == nil {
		fmt.Print("nil")
		return
	}
	fmt.Print(le.value)
}

func (astp *AstPrinter) visitForUnaryExpression(ue *UnaryExpression) {
	astp.parenthesize(ue.operator.Lexeme, ue.right)
}

func (astp *AstPrinter) parenthesize(name string, exprs ...Expression) {
	fmt.Print("(")
	fmt.Print(name)
	for _, expr := range exprs {
		fmt.Print(" ")
		expr.accept(astp)
	}
	fmt.Print(")")
}
