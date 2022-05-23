package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var flagRomFile string

func main() {

	flag.StringVar(&flagRomFile, "rom", "prg.bin", "Read the following rom file")
	flag.Parse()

	cpu := &CPU{}
	cpu.PC = PC_START_LOCATION

	content, err := ioutil.ReadFile(flagRomFile)
	if err != nil {
		log.Fatal(err)
	}
	if len(content) != int(ROM_SIZE_MAX) {
		log.Fatalf("Failed to load rom (%d!=%d)", len(content), ROM_SIZE_MAX)
	}
	copy(Rom[:], content)

	for !cpu.IsHalted() {
		cpu.Execute()
	}
}
