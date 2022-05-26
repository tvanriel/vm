package vm

const MEM_SIZE_MAX uint64 = 0xFF_FF
const MEM_ADDR_START = 0x00_00
const MEM_ADDR_END = 0xFF_FF

type MemDevice [MEM_SIZE_MAX]byte

func (m *MemDevice) Get(addr Address) byte {
	return m[addr]
}

func (m *MemDevice) Put(addr Address, val byte) {
	m[addr] = val
}
func (m *MemDevice) Range() AddressRange {
	return [2]Address{MEM_ADDR_START, MEM_ADDR_END}
}
func (m *MemDevice) Content() {

}

func (*MemDevice) Name() string { return "Memory" }

var _ Device = &MemDevice{}
var Memory = &MemDevice{}

func init() {
	RegisterDevice(Memory)
}
