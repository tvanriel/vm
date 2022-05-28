package codegen

import (
	"errors"

	goparser "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/parser"
)

// Get the set of constants with their values from the AST.
func ResolveLabels(ast *goparser.AST, constants *map[string]uint64) (map[string]*Label, error) {
	labels := map[string]*Label{}
	for _, child := range ast.Children {
		err := walkAstForLabels(child, &labels, constants)
		if err != nil {
			return nil, err
		}
	}
	return labels, nil
}

func walkAstForLabels(ast *goparser.AST, labels *map[string]*Label, constants *map[string]uint64) error {
	if ast.ValueType == parser.LabelNode {
		if _, has := (*labels)[ast.ValueString]; has {
			return errors.New("Duplicate label declaration " + ast.ValueString)
		}

		i, err := getInstructions(ast, constants)
		if err != nil {
			return err
		}
		label := &Label{
			Name:         ast.ValueString,
			Instructions: i,
		}

		(*labels)[ast.ValueString] = label
	}
	return nil
}

func getInstructions(ast *goparser.AST, constants *map[string]uint64) ([]*Instruction, error) {
	instructions := []*Instruction{}
	for _, child := range ast.Children {
		if child.ValueType == parser.InstructionNode {
			i, err := makeInstruction(child, constants)
			if err != nil {
				return nil, err
			}
			instructions = append(instructions, i)
		}
	}
	return instructions, nil
}

func makeInstruction(ast *goparser.AST, constants *map[string]uint64) (*Instruction, error) {
	args, err := makeArguments(ast, constants)
	if err != nil {
		return nil, err
	}
	return &Instruction{
		Name:      ast.ValueString,
		Arguments: args,
	}, nil
}

func makeArguments(ast *goparser.AST, constants *map[string]uint64) ([]*Argument, error) {
	arguments := []*Argument{}
	for _, child := range ast.Children {
		arg, err := makeArgument(child, constants)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, arg)
	}
	return arguments, nil
}

func makeArgument(ast *goparser.AST, constants *map[string]uint64) (*Argument, error) {
	if ast.ValueType == parser.ConstantNode {
		if ast.ValueString == "A" {
			return &Argument{
				Type:  RegisterArgument,
				Value: uint64(REGISTER_A),
			}, nil
		}
		if ast.ValueString == "B" {
			return &Argument{
				Type:  RegisterArgument,
				Value: uint64(REGISTER_B),
			}, nil
		}
		if ast.ValueString == "X" {
			return &Argument{
				Type:  RegisterArgument,
				Value: uint64(REGISTER_X),
			}, nil
		}
		if ast.ValueString == "Y" {
			return &Argument{
				Type:  RegisterArgument,
				Value: uint64(REGISTER_Y),
			}, nil
		}
		val, ok := (*constants)[ast.ValueString]
		if !ok {
			return nil, errors.New("Undeclared constant " + ast.ValueString)
		}
		return &Argument{
			Type:  ConstantArgument,
			Value: val,
		}, nil
	}

	if isNumberNode(ast.ValueType) {
		n, err := ParseNumber(ast)
		if err != nil {
			return nil, err
		}
		return &Argument{
			Type:  ConstantArgument,
			Value: n,
		}, nil
	}

	if ast.ValueType == parser.ExpressionNode {
		n, err := evaluateExpression(ast, constants)
		if err != nil {
			return nil, err
		}
		return &Argument{
			Type:  ConstantArgument,
			Value: n,
		}, nil
	}

	if ast.ValueType == parser.WordNode {
		return &Argument{
			Type:        LabelArgument,
			ValueString: ast.ValueString,
		}, nil
	}
	return nil, errors.New("expected number, expression or constant: got " + ast.ValueString)
}
