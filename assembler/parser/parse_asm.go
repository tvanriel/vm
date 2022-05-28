package parser

import (
	"fmt"
	"strconv"

	lex "github.com/bbuck/go-lexer"
	"github.com/tvanriel/vm/assembler/lexer"
	parse "github.com/tvanriel/go-parser"
)

type register int64

const (
	REGISTER_A register = iota + 1
	REGISTER_B
	REGISTER_X
	REGISTER_Y
)

type ArgumentType int

const (
	ArgumentTypeConstant
	ArgumentTypeNumber
)

type Argument struct {
	Value uint64
	Type  ArgumentType
}

type Instruction struct {
	Name      string
	Arguments []*Argument
}

type Label struct {
	Name         string
	Instructions []*Instruction
}

type ParseContext struct {
	Current int
	Labels  map[string]*Label
}

type Constants map[string]uint64

func Parse(tokens []*lex.Token) (*ParseContext, error) {
	consts, err := ExploreConstants(tokens)
	if err != nil {
		return nil, err
	}
	fmt.Println(consts)
	ctx := &ParseContext{}
	return ctx, nil
}

func ExploreConstants(tokens []*lex.Token) (Constants, error) {
	consts := Constants{}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type == lexer.ConstantToken {
			if tokens[i+1].Type != lexer.EqualsSign {
				// Not an assignment.
				continue
			}
			val := tokens[i+2]
			n, err := ParseNumber(val)
			if err != nil {
				return nil, err
			}
			consts[token.Value] = n
		}
	}
	return consts, nil
}


func ParseNumber(val *lex.Token) {
	switch val.Type {
	case lexer.OctalToken:
		n, err := ParseOct(val.Value)
		if err != nil {
			return nil, err
		}
		return n, nil
	case lexer.BinaryToken:
		n, err := ParseBin(val.Value)
		if err != nil {
			return nil, err
		}
		return n, nil
	case lexer.HexadecimalToken:
		n, err := ParseHex(val.Value)
		if err != nil {
			return nil, err
		}
		return n, nil
	case lexer.DecimalToken:
		n, err := ParseDec(val.Value)
		if err != nil {
			return nil, err
		}
		return n, nil
	}
}

func ParseHex(str string) (uint64, error) {
	return strconv.ParseUint(str[2:], 16, 64)
}

func ParseBin(str string) (uint64, error) {
	return strconv.ParseUint(str[2:], 2, 64)
}

func ParseOct(str string) (uint64, error) {
	return strconv.ParseUint(str[2:], 8, 64)
}

func ParseDec(str string) (uint64, error) {
	return strconv.ParseUint(str[2:], 10, 64)
}

func compileLabels(tokens []*lex.Token, constants map[string]uint64, ctx *ParseContext) {
	var currentLabel string
	var currentInstruction *Instruction

	
		if token.Type == lexer.LabelToken {
			currentLabel = token.Value
			ctx.Labels[currentLabel] = &Label{}
		}
		if token.Type == lexer.InstructionToken {
			currentInstruction = &Instruction{}
			ctx.Labels[currentLabel].Instructions = append(ctx.Labels[currentLabel].Instructions, currentInstruction)
		}
	
}
