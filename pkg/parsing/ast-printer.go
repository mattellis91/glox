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
	//TODO: implement visit for binary expression
	astp.parenthesize(be.operator.Lexeme, be.left, be.right)
}

func (astp *AstPrinter) visitForGroupingExpression(ge *GroupingExpression) {
	//TODO: implement visit for grouping expression
	astp.parenthesize("Group", ge.expression)
}

func (astp *AstPrinter) visitForLiteralExpression(le *LiteralExpression) {
	//TODO: implement visit for literal expression
	if le == nil {
		fmt.Print("nil")
		return
	}
	fmt.Print(le.value)
}

func (astp *AstPrinter) visitForUnaryExpression(ue *UnaryExpression) {
	//TODO: implement visit for unary expression
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
