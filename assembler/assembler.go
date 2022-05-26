package assembler

import (
	"fmt"
	"log"

	"github.com/tvanriel/vm/assembler/lexer"
	"github.com/tvanriel/vm/assembler/preprocessor"
)

func Assemble(entryPath string, out string) {
	str, err := preprocessor.Process(entryPath)
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	tokens := lexer.Tokenize(str)
	// ast, err := parser.Parse(tokens)
	fmt.Println(tokens)
	if err != nil {
		log.Fatalln("Failed to parse: " + err.Error())
	}

}
