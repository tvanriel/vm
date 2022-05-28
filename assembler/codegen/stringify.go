package codegen

import "errors"

// import (
// goparser "github.com/tvanriel/go-parser"
// )

// Stringify's job is to find a place for all the instructions and to put a place for all the labels and to resolve jumps.
func Stringify(labels map[string]*Label) ([]*Instruction, error) {

	res := []*Instruction{}
	labelPos := map[string]int{}

	start, ok := labels["main"]
	if !ok {
		return nil, errors.New("missing main routine")
	}
	delete(labels, "main")
	labelPos["main"] = 0
	i := 0
	for _, instruction := range start.Instructions {
		i += instruction.Size()
		res = append(res, instruction)
	}
	for _, label := range labels {
		labelPos[label.Name] = i
		for _, instruction := range label.Instructions {
			i += instruction.Size()
			res = append(res, instruction)
		}
	}
	for _, instruction := range res {
		if instruction.Name == "jmp" {
			if instruction.Arguments[0].Type == LabelArgument {
				instruction.Arguments[0] = &Argument{
					Type:  ConstantArgument,
					Value: uint64(labelPos[instruction.Arguments[0].ValueString]),
				}
			}

		}
	}

	return res, nil
}
