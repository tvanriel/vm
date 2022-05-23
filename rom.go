package main

const ROM_SIZE_MAX uint64 = 0xFF_FF
const ROM_ADDR_START Address = 0x1_00_00
const ROM_ADDR_END Address = 0x1_FF_FF

type RomDevice [ROM_SIZE_MAX]byte

func (m *RomDevice) Get(addr Address) byte {
	return m[addr-ROM_ADDR_START]
}

func (m *RomDevice) Put(addr Address, val byte) {}

func (m *RomDevice) Range() AddressRange {
	return [2]Address{ROM_ADDR_START, ROM_ADDR_END}
}

func (*RomDevice) Name() string { return "ROM" }

var _ Device = &RomDevice{}
var Rom *RomDevice = &RomDevice{}

func init() {
	RegisterDevice(Rom)
}
