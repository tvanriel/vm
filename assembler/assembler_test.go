package assembler

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	lex "github.com/bbuck/go-lexer"
	"github.com/tvanriel/vm/assembler/codegen"
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
		t.Error("Failed to preprocess file: " + err.Error())
		return
	}
	if len(str) != len(expected1) {
		t.Errorf("Expected string length to be %d but got %d", len(expected1), len(str))
		return
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected1); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
		return
	}
}

var expected2 = "\n\nmain:\n    lda 123123"

func TestPreprocess2(t *testing.T) {
	str, err := preprocessor.Process("test/preprocessor_ifndef/prg.s")
	if err != nil {
		t.Error("Failed to preprocess file: " + err.Error())
		return
	}
	if len(str) != len(expected2) {
		t.Errorf("Expected string length to be %d but got %d", len(expected2), len(str))
		return
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected2); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
		return
	}
}

var expected3 = "\n\nmain:\n    lda 123123\n"

func TestPreprocess3(t *testing.T) {
	str, err := preprocessor.Process("test/preprocessor_ifdef/prg.s")
	if err != nil {
		t.Error("Failed to preprocess file: " + err.Error())
		return
	}
	if len(str) != len(expected3) {
		t.Errorf("Expected string length to be %d but got %d", len(expected3), len(str))
		return
	}

	if a, e := strings.TrimSpace(str), strings.TrimSpace(expected3); a != e {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(e, a))
		return
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
		return
	}
	t.Log(ast)
}

var expectedConsts = map[string]uint64{
	"ROM_START":      0x10000,
	"ROM_END":        0x10000 + 0xffff,
	"DISPLAY_ENABLE": 0x10000 + 0x10000 + 0xffff,
}

func TestConstResolver(t *testing.T) {
	str, _ := preprocessor.Process("./test/resolver_1/prg.s")
	ast, err := parser.Parse(lexer.Tokenize(str))
	if err != nil {
		t.Error(err)
		return
	}
	consts, err := codegen.ResolveConsts(ast)
	if err != nil {
		t.Error(err)
		return
	}
	for k, v := range consts {
		if expectedConsts[k] != v {
			t.Errorf("Expected value for %s to be but got %x", k, v)
		}
	}
}

func TestLabelResolver(t *testing.T) {
	str, _ := preprocessor.Process("./test/resolver_1/prg.s")
	ast, err := parser.Parse(lexer.Tokenize(str))
	if err != nil {
		t.Error(err)
		return
	}
	consts, err := codegen.ResolveConsts(ast)
	if err != nil {
		t.Error(err)
		return
	}
	labels, err := codegen.ResolveLabels(ast, &consts)
	fmt.Println(labels)
	if err != nil {
		t.Error(err)
		return
	}
}
func TestStringify(t *testing.T) {
	str, _ := preprocessor.Process("./test/resolver_1/prg.s")
	ast, err := parser.Parse(lexer.Tokenize(str))
	if err != nil {
		t.Error(err)
		return
	}
	consts, err := codegen.ResolveConsts(ast)
	if err != nil {
		t.Error(err)
		return
	}
	labels, err := codegen.ResolveLabels(ast, &consts)
	fmt.Println(labels)
	if err != nil {
		t.Error(err)
		return
	}
	instructions, err := codegen.Stringify(labels)
	if err != nil {
		t.Error(err)
	}
	rom := codegen.Assemble(instructions)
	fmt.Println(rom)
}
