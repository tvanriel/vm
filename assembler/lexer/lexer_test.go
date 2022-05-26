package lexer_test

import (
	"io/ioutil"
	"testing"

	lex "github.com/bbuck/go-lexer"
	"github.com/tvanriel/vm/assembler/lexer"
)

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func Test_LexTest(t *testing.T) {
	expected := []lex.Token{
		{Type: lexer.CommentToken, Value: "; a comment here"},
		{Type: lexer.LabelToken, Value: "main"},
		{Type: lexer.ColonSign, Value: ":"},
		{Type: lexer.InstructionToken, Value: "lda"},
		{Type: lexer.HexadecimalToken, Value: "0xDEADBEEFDEADBEEF"},
		{Type: lexer.InstructionToken, Value: "ldb"},
		{Type: lexer.HexadecimalToken, Value: "0xC0FFEEC0FFEEC0FF"},
		{Type: lexer.InstructionToken, Value: "ldx"},
		{Type: lexer.BinaryToken, Value: "0b0101010101010101"},
		{Type: lexer.InstructionToken, Value: "ldy"},
		{Type: lexer.OctalToken, Value: "012345671234567123"},
		{Type: lexer.InstructionToken, Value: "jmp"},
		{Type: lexer.WordToken, Value: "label"},
		{Type: lexer.LabelToken, Value: "label"},
		{Type: lexer.ColonSign, Value: ":"},
	}
	tokens := lexer.Tokenize(ReadFile("./test/1/program.s"))
	if len(tokens) != len(expected) {
		t.Errorf("Expected number of tokens to be %d but got %d", len(expected), len(tokens))
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Value != expected[i].Value {
			t.Errorf("Expected Value of token %d to be %s but got %s", i, expected[i].Value, tokens[i].Value)
			return
		}

		if tokens[i].Type != expected[i].Type {
			t.Errorf("Expected Type of token %d to be %v but got %v", i, expected[i].Type, tokens[i].Type)
			return
		}

	}
}
