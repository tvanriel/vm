package codegen

import (
	"errors"
	"strings"

	goparser "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/parser"
)

func evaluateExpression(expression *goparser.AST, consts *map[string]uint64) (uint64, error) {
	if isNumberNode(expression.ValueType) {
		return ParseNumber(expression)
	}

	if expression.ValueType == parser.ConstantNode {
		if n, ok := (*consts)[expression.ValueString]; ok {
			return n, nil
		}
		return 0, errors.New("Undeclared constant " + expression.ValueString)

	}

	if len(expression.Children) == 1 {
		return evaluateExpression(expression.Children[0], consts)
	}

	if len(expression.Children) == 3 {
		left, err := evaluateExpression(expression.Children[0], consts)
		if err != nil {
			return 0, err
		}
		right, err := evaluateExpression(expression.Children[2], consts)
		if err != nil {
			return 0, err
		}
		if !isOperator(expression.Children[1].ValueString) {
			return 0, errors.New("eval error: expected operator, got " + expression.Children[1].ValueString)
		}
		return operation(left, right, expression.Children[1].ValueString), nil
	}
	return 0, errors.New("Cannot evaluate expression: " + expression.Print(0))

}

func operation(left uint64, right uint64, op string) uint64 {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return 0
}

func isOperator(op string) bool {
	return strings.Contains("+-*/", op)
}
