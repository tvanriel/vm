package codegen

type Register uint64

type ArgumentType uint64

const (
	REGISTER_A Register = iota
	REGISTER_B
	REGISTER_X
	REGISTER_Y
)

const (
	RegisterArgument ArgumentType = iota
	AddressArgument
	ConstantArgument
	LabelArgument
)

type Argument struct {
	Type ArgumentType

	Value       uint64 // Either of these will be filled, based on the type of argument.
	ValueString string
}

type Instruction struct {
	Name      string
	Arguments []*Argument
}

// Get the amount of bytes this instruction wil take in the binary.
// TODO: use a smarter approach for this.
func (i *Instruction) Size() int {
	return len(i.Bytes())
}

type Label struct {
	Name         string
	Instructions []*Instruction
}

// Get the amount of bytes this label will take in the binary.
func (l *Label) Size() int {
	sum := 0

	for _, i := range l.Instructions {
		sum += i.Size()
	}

	return sum
}
