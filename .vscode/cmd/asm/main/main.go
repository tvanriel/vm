package main

import (
	"github.com/tvanriel/cpu-emu/assembler/lexer"
)

func main() {
	// Get a file
	content := ReadFile(inFile)
	// Let the lexer read it
	tokens := lexer.Tokenize(content)
	ast := parser.Parse(tokens)
	opcodes := codegen.GenerateFromAst(ast)
	WriteFiles(outFile)
}
