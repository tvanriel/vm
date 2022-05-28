package codegen

import "github.com/tvanriel/vm/assembler/parser"

// Stringify's job is to find a place for all the instructions and to put a place for all the labels and to resolve jumps.
func Stringify(ctx *parser.ParseContext, programId [PROGRAM_ID_LENGTH]byte) {

	// calculate the start of the first instruction we're able to put in.
	// start := rom_start + len(magic) + 1 + CHECKSUM_LENGTH + PROGRAM_ID_LENGTH

	// namedLabels = getNamedLabels(ctx)
	// current := start
	
}

func getNamedLabels(list []*parser.Label) map[string]*parser.Label {
	labels := make(map[string]*parser.Label)
	for _, label := range list {
		labels[label.Name] = label
	}
	return labels
}
