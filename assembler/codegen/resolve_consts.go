package codegen

import (
	"errors"

	goparser "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/parser"
)

// Get the set of constants with their values from the AST.
func ResolveConsts(ast *goparser.AST) (map[string]uint64, error) {
	consts := map[string]uint64{}
	for _, child := range ast.Children {
		err := walkAstForConsts(child, &consts)
		if err != nil {
			return nil, err
		}
	}
	return consts, nil
}

func walkAstForConsts(ast *goparser.AST, consts *map[string]uint64) error {
	if ast.ValueType == parser.ConstantNode {

		if _, has := (*consts)[ast.ValueString]; has {
			return errors.New("Duplicate constant declaration " + ast.ValueString)
		}

		n, err := getConstValue(ast, consts)

		if err != nil {
			return err
		}

		(*consts)[ast.ValueString] = n
	}

	return nil
}

func getConstValue(ast *goparser.AST, consts *map[string]uint64) (uint64, error) {
	if ast.Children[0].ValueType == parser.ConstantNode {
		if _, has := (*consts)[ast.Children[0].ValueString]; !has {
			return 0, errors.New("Undeclared constant " + ast.Children[0].ValueString)
		}

		return (*consts)[ast.Children[0].ValueString], nil
	}

	if isNumberNode(ast.Children[0].ValueType) {
		return ParseNumber(ast.Children[0])
	}
	if ast.Children[0].ValueType == parser.ExpressionNode {
		return evaluateExpression(ast.Children[0], consts)
	}

	return 0, errors.New("Cannot get value for constant " + ast.ValueString)
}

func isNumberNode(t goparser.NodeType) bool {
	return t == parser.OctalNumberNode ||
		t == parser.HexadecimalNumberNode ||
		t == parser.BinaryNumberNode ||
		t == parser.DecimalNumberNode
}
