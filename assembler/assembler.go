package assembler

import (
	"fmt"
	"log"

	"github.com/tvanriel/vm/assembler/lexer"
	"github.com/tvanriel/vm/assembler/parser"
	"github.com/tvanriel/vm/assembler/preprocessor"
)

func Assemble(entryPath string, out string) {
	str, err := preprocessor.Process(entryPath)
	if err != nil {
		log.Fatalln("Failed to preprocess file: " + err.Error())
	}
	tokens := lexer.Tokenize(str)
	ast, err := parser.Parse(tokens)

	var prgid [24]byte
	copy(prgid[:], "standardprogram         ")

	fmt.Println(ast)
	// stringcode := codegen.Stringify(ast, prgid)
	// bin := codegen.Generate(stringcode)
	if err != nil {
		log.Fatalln("Failed to parse: " + err.Error())
	}

}
