package assembler

import (
	"log"
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/tvanriel/vm/assembler/preprocessor"
)

var expected1 = "; comment\n\nDEFINITION = 0x100\n\nlabdef:\n    lda DEFINITION\n    sta 0x100\n\nmain:\n    jmp labdef\n    lda 0x100"

func TestPreprocess1(t *testing.T) {
	str, err := preprocessor.Process("test/1/prg.s")
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
	str, err := preprocessor.Process("test/2/prg.s")
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
	str, err := preprocessor.Process("test/3/prg.s")
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
