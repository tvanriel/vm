package codegen

import (
	"errors"

	goparser "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/parser"
)

// Get the set of constants with their values from the AST.
func ResolveLabels(ast *goparser.AST) (map[string]*Label, error) {
	labels := map[string]*Label{}
	for _, child := range ast.Children {
		err := walkAstForLabels(child, &labels)
		if err != nil {
			return nil, err
		}
	}
	return labels, nil
}

func walkAstForLabels(ast *goparser.AST, labels *map[string]*Label) error {
	if ast.ValueType == parser.LabelNode {
		if _, has := (*labels)[ast.ValueString]; has {
			return errors.New("Duplicate label declaration " + ast.ValueString)
		}

		label := &Label{
			Name:         ast.ValueString,
			Instructions: getInstructions(ast.Children),
		}

		(*labels)[ast.ValueString] = label
	}
	return nil
}

func getInstructions(ast []*goparser.AST) []*Instruction {
	return []*Instruction{}
}
