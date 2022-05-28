package codegen

import (
	"errors"
	"strconv"

	goparser "github.com/tvanriel/go-parser"
	"github.com/tvanriel/vm/assembler/parser"
)

func ParseNumber(val *goparser.AST) (uint64, error) {
	switch val.ValueType {
	case parser.OctalNumberNode:
		n, err := ParseOct(val.ValueString)
		if err != nil {
			return 0, err
		}
		return n, nil
	case parser.BinaryNumberNode:
		n, err := ParseBin(val.ValueString)
		if err != nil {
			return 0, err
		}
		return n, nil
	case parser.HexadecimalNumberNode:
		n, err := ParseHex(val.ValueString)
		if err != nil {
			return 0, err
		}
		return n, nil
	case parser.DecimalNumberNode:
		n, err := ParseDec(val.ValueString)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return 0, errors.New("Invalid number type: " + val.ValueString)
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
