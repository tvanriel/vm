package assembler

import (
	"log"
	"testing"

	"github.com/tvanriel/vm/assembler/preprocessor"
)

func TestPreprocess(t *testing.T) {
	str, err := preprocessor.Process("test/1/prg.s")
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	t.Log(str)

	// ast, err := parser.Parse(tokens)
	// if err != nil {
	// log.Fatalln("Failed to parse: " + err.Error())
	// }

}
