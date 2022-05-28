package parser

import (
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

func IsNumber(t lex.TokenType) bool {
	return t == lexer.OctalToken ||
		t == lexer.BinaryToken ||
		t == lexer.HexadecimalToken ||
		t == lexer.DecimalToken
}
