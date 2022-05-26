package parser

import (
	"fmt"
	"strconv"

	lex "github.com/bbuck/go-lexer"
	"github.com/tvanriel/vm/assembler/lexer"
)

type register int64

const (
	REGISTER_A register = iota + 1
	REGISTER_B
	REGISTER_X
	REGISTER_Y
)

type Argument struct {
	Value int64
}

type Instruction struct {
	Name      string
	Arguments []*Argument
}

type Label struct {
	Name         string
	Position     int
	Instructions []*Instruction
}

type ParseContext struct {
	Current int
	Labels  []*Label
}

type Constants map[string]int64
type Labels map[string]int64

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
			switch val.Type {
			case lexer.OctalToken:
				n, err := ParseOct(val.Value)
				if err != nil {
					return nil, err
				}
				consts[token.Value] = n
			case lexer.BinaryToken:
				n, err := ParseBin(val.Value)
				if err != nil {
					return nil, err
				}
				consts[token.Value] = n
			case lexer.HexadecimalToken:
				n, err := ParseHex(val.Value)
				if err != nil {
					return nil, err
				}
				consts[token.Value] = n
			case lexer.DecimalToken:
				n, err := ParseDec(val.Value)
				if err != nil {
					return nil, err
				}
				consts[token.Value] = n
			}
		}
	}
	return consts, nil
}

func ParseHex(str string) (int64, error) {
	return strconv.ParseInt(str[2:], 16, 64)
}

func ParseBin(str string) (int64, error) {
	return strconv.ParseInt(str[2:], 2, 64)
}

func ParseOct(str string) (int64, error) {
	return strconv.ParseInt(str[2:], 8, 64)
}

func ParseDec(str string) (int64, error) {
	return strconv.ParseInt(str[2:], 10, 64)
}
