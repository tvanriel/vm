package codegen

import "hash/crc32"

var magic = []byte("EXECVM")

const rom_start = 0x10000
const version = 0x1

const CHECKSUM_LENGTH = 8
const PROGRAM_ID_LENGTH = 24
const ROM_SIZE = 0xFF_FF

var HEADER_SIZE = uint64(len(magic) + 1)
var OP_START = HEADER_SIZE + 4 // 4 bytes for the CRC

type ROM [ROM_SIZE]byte

func Assemble(instructions []*Instruction) *ROM {
	buf := &ROM{}
	WriteHeader(buf)
	WriteBody(buf, instructions)
	WriteFooter(buf)
	WriteCRC(buf)

	return buf
}

func WriteBody(buf *ROM, instructions []*Instruction) {
	i := OP_START
	for _, instruction := range instructions {
		WriteAt(buf, int(i), instruction.Bytes())
		i += uint64(instruction.Size())
	}
}

func WriteHeader(buf *ROM) {
	WriteAt(buf, 0, magic)
	WriteAt(buf, len(magic), []byte{version})
}

func WriteFooter(buf *ROM) {
	WriteUint64At(buf, ROM_SIZE-8, rom_start+HEADER_SIZE)
}

func WriteAt(buf *ROM, at int, bytes []byte) {
	for _, b := range bytes {
		buf[at] = b
		at++
	}
}

func WriteCRC(buf *ROM) {
	WriteAt(buf, int(HEADER_SIZE), dwordfromint64(uint64(crc32.ChecksumIEEE(buf[HEADER_SIZE:ROM_SIZE-9]))))
}

func WriteUint64At(buf *ROM, at int, value uint64) {
	WriteAt(buf, at, qwordfromint64(value))
}

func qwordfromint64(v uint64) []byte {
	return []byte{
		byte(0xff & (v >> 56)),
		byte(0xff & (v >> 48)),
		byte(0xff & (v >> 40)),
		byte(0xff & (v >> 32)),
		byte(0xff & (v >> 24)),
		byte(0xff & (v >> 16)),
		byte(0xff & (v >> 8)),
		byte(0xff & v),
	}
}

func dwordfromint64(v uint64) []byte {
	return []byte{
		byte(0xff & (v >> 24)),
		byte(0xff & (v >> 16)),
		byte(0xff & (v >> 8)),
		byte(0xff & v),
	}
}

func wordfromint64(v uint64) []byte {
	return []byte{
		byte(0xff & (v >> 8)),
		byte(0xff & v),
	}
}

func bytefromint64(v uint64) []byte {
	return []byte{
		byte(0xff & v),
	}
}
