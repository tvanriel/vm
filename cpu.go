package main

import (
	"fmt"

	"github.com/tvanriel/cpu-emulator/operations"
)

const PC_START_LOCATION = 0x1_ff_f6 // 9-bytes under rom-top.  Enough for JMP instr + 8 byte addressing.
const SP_START_LOCATION = 0xFF_FF   // Starts at last address of memory.

type CPU struct {
	// Flags is the sum of all flags that are currently set in the CPU.
	Flags Flag

	// Program-counter.  Contains the address of the next instruction.
	PC uint64
	SP uint64

	// Accumulator register
	A uint64

	// General-purpose registers
	B uint64
	X uint64
	Y uint64
}

func (c *CPU) IsHalted() bool {
	return c.Flags&FLAG_HALT == 1
}

func (c *CPU) Execute() {
	// Fetch the instruction from memory.
	fmt.Println("fetch")
	instr := operations.Opcode(GetFromAddressSpace(Address(c.PC)))

	fmt.Printf("instr: %x\n", instr)
	// Decode the instruction
	if action, ok := Instructions[instr]; ok {
		// Execute
		action(c)
		c.PC++
		return
	}
	c.Flags &= FLAG_HALT
}
