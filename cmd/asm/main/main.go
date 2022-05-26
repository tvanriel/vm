package main

import (
	"flag"

	"github.com/tvanriel/vm/assembler"
)

func main() {
	entryFile := flag.String("entry", "prg.s", "File to assemble and link")
	outFile := flag.String("out", "prg.bin", "Output file")
	assembler.Assemble(*entryFile, *outFile)
}
