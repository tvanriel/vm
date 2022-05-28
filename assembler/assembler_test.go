package assembler

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	lex "github.com/bbuck/go-lexer"
	"github.com/tvanriel/vm/assembler/lexer"
	"github.com/tvanriel/vm/assembler/parser"
	"github.com/tvanriel/vm/assembler/preprocessor"
)

var expected1 = "; comment\n\nDEFINITION = 0x100\n\nlabdef:\n    lda DEFINITION\n    sta 0x100\n\nmain:\n    jmp labdef\n    lda 0x100"

func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func TestPreprocess1(t *testing.T) {
	str, err := preprocessor.Process("test/preprocessor_include/prg.s")
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	if len(str) != len(expected1) {
		t.Errorf("Expected string length to be %d but got %d", len(expected1), len(str))
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected1); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
	}
}

var expected2 = "\n\nmain:\n    lda 123123"

func TestPreprocess2(t *testing.T) {
	str, err := preprocessor.Process("test/preprocessor_ifndef/prg.s")
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	if len(str) != len(expected2) {
		t.Errorf("Expected string length to be %d but got %d", len(expected2), len(str))
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected2); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
	}
}

var expected3 = "\n\nmain:\n    lda 123123\n"

func TestPreprocess3(t *testing.T) {
	str, err := preprocessor.Process("test/preprocessor_ifdef/prg.s")
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	if len(str) != len(expected3) {
		t.Errorf("Expected string length to be %d but got %d", len(expected3), len(str))
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected3); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
	}
}

func TestLexer(t *testing.T) {
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
	tokens := lexer.Tokenize(ReadFile("./test/lexer_1/prg.s"))
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

func TestParser(t *testing.T) {
	str, _ := preprocessor.Process("./test/parser_1/prg.s")
	ast, err := parser.Parse(lexer.Tokenize(str))
	if err != nil {
		t.Error(err)
	}
	t.Log(ast)
}
