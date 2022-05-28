package parser

import (
	"errors"
	"strconv"

	lex "github.com/bbuck/go-lexer"
	parse "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/lexer"
)

var tokennumtonodenum = map[lex.TokenType]parse.NodeType{
	lexer.HexadecimalToken: HexadecimalNumberNode,
	lexer.DecimalToken:     DecimalNumberNode,
	lexer.OctalToken:       OctalNumberNode,
	lexer.BinaryToken:      BinaryNumberNode,
}

func ParseNumber(val *lex.Token) (uint64, error) {
	switch val.Type {
	case lexer.OctalToken:
		n, err := ParseOct(val.Value)
		if err != nil {
			return 0, err
		}
		return n, nil
	case lexer.BinaryToken:
		n, err := ParseBin(val.Value)
		if err != nil {
			return 0, err
		}
		return n, nil
	case lexer.HexadecimalToken:
		n, err := ParseHex(val.Value)
		if err != nil {
			return 0, err
		}
		return n, nil
	case lexer.DecimalToken:
		n, err := ParseDec(val.Value)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, errors.New("Invalid number type: " + val.Value)
}

func IsNumber(t lex.TokenType) bool {
	return t == lexer.OctalToken ||
		t == lexer.BinaryToken ||
		t == lexer.HexadecimalToken ||
		t == lexer.DecimalToken
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
