package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/tvanriel/vm"
)

var flagRomFile string

func main() {

	flag.StringVar(&flagRomFile, "rom", "prg.bin", "Read the following rom file")
	flag.Parse()

	cpu := &vm.CPU{}
	cpu.PC = vm.PC_START_LOCATION

	content, err := ioutil.ReadFile(flagRomFile)
	if err != nil {
		log.Fatal(err)
	}
	if len(content) != int(vm.ROM_SIZE_MAX) {
		log.Fatalf("Failed to load rom (%d!=%d)", len(content), vm.ROM_SIZE_MAX)
	}
	copy(vm.Rom[:], content)

	for !cpu.IsHalted() {
		cpu.Execute()
	}
}
