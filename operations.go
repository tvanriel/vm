package vm

import (
	"math"

	op "github.com/tvanriel/vm/operations"
)

type Operation func(c *CPU)

var Instructions map[op.Opcode]Operation

func getWordFrom(from Address) [2]byte {

	return [2]byte{
		GetFromAddressSpace(from + 1),
		GetFromAddressSpace(from),
	}
}

func getDwordFrom(from Address) [4]byte {

	return [4]byte{
		GetFromAddressSpace(from + 3),
		GetFromAddressSpace(from + 2),
		GetFromAddressSpace(from + 1),
		GetFromAddressSpace(from),
	}
}

func getQwordFrom(from Address) [8]byte {

	return [8]byte{
		GetFromAddressSpace(from + 7),
		GetFromAddressSpace(from + 6),
		GetFromAddressSpace(from + 5),
		GetFromAddressSpace(from + 4),
		GetFromAddressSpace(from + 3),
		GetFromAddressSpace(from + 2),
		GetFromAddressSpace(from + 1),
		GetFromAddressSpace(from),
	}
}

func getByte(c *CPU) byte {
	val := GetFromAddressSpace(Address(c.PC + 1))
	c.PC += 1
	return val
}

func getQword(c *CPU) [8]byte {
	val := getQwordFrom(Address(c.PC + 1))
	c.PC += 8
	return val
}

func getDword(c *CPU) [4]byte {
	val := getDwordFrom(Address(c.PC + 1))
	c.PC += 4
	return val
}

func getWord(c *CPU) [2]byte {
	val := getWordFrom(Address(c.PC + 1))
	c.PC += 2
	return val
}

func int64fromqword(bytes [8]byte) uint64 {
	return uint64(bytes[7])<<56 +
		uint64(bytes[6])<<48 +
		uint64(bytes[5])<<40 +
		uint64(bytes[4])<<32 +

		uint64(bytes[3])<<24 +
		uint64(bytes[2])<<16 +
		uint64(bytes[1])<<8 +
		uint64(bytes[0])
}

func int64fromdword(bytes [4]byte) uint64 {
	return uint64(bytes[3])<<24 +
		uint64(bytes[2])<<16 +
		uint64(bytes[1])<<8 +
		uint64(bytes[0])
}

func int64fromword(bytes [2]byte) uint64 {
	return uint64(bytes[1])<<8 +
		uint64(bytes[0])
}

func qwordfromint64(v uint64) [8]byte {
	return [8]byte{
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

func dwordfromint64(v uint64) [4]byte {
	return [4]byte{
		byte(0xff & (v >> 24)),
		byte(0xff & (v >> 16)),
		byte(0xff & (v >> 8)),
		byte(0xff & v),
	}
}
func wordfromint64(v uint64) [2]byte {
	return [2]byte{
		byte(0xff & (v >> 8)),
		byte(0xff & v),
	}
}

func addByteToStack(value byte, c *CPU) {
	PutInAddressSpace(Address(c.SP-1), value)
	c.SP--
}
func addWordToStack(value [2]byte, c *CPU) {
	PutInAddressSpace(Address(c.SP-1), value[0])
	PutInAddressSpace(Address(c.SP-2), value[1])
	c.SP -= 2
}
func addDwordToStack(value [4]byte, c *CPU) {
	PutInAddressSpace(Address(c.SP-1), value[0])
	PutInAddressSpace(Address(c.SP-2), value[1])
	PutInAddressSpace(Address(c.SP-3), value[2])
	PutInAddressSpace(Address(c.SP-4), value[3])
	c.SP -= 4
}
func addQwordToStack(value [8]byte, c *CPU) {
	PutInAddressSpace(Address(c.SP-1), value[0])
	PutInAddressSpace(Address(c.SP-2), value[1])
	PutInAddressSpace(Address(c.SP-3), value[2])
	PutInAddressSpace(Address(c.SP-4), value[3])
	PutInAddressSpace(Address(c.SP-5), value[4])
	PutInAddressSpace(Address(c.SP-6), value[5])
	PutInAddressSpace(Address(c.SP-7), value[6])
	PutInAddressSpace(Address(c.SP-8), value[7])
	c.SP -= 8
}

func popByteFromStack(c *CPU) byte {
	val := GetFromAddressSpace(Address(c.SP))
	c.SP += 1
	return val
}

func popWordFromStack(c *CPU) [2]byte {
	val := [2]byte{
		GetFromAddressSpace(Address(c.SP)),
		GetFromAddressSpace(Address(c.SP + 1)),
	}
	c.SP += 2
	return val
}

func popDWordFromStack(c *CPU) [4]byte {
	val := [4]byte{
		GetFromAddressSpace(Address(c.SP)),
		GetFromAddressSpace(Address(c.SP + 1)),
		GetFromAddressSpace(Address(c.SP + 2)),
		GetFromAddressSpace(Address(c.SP + 3)),
	}
	c.SP += 4
	return val
}
func popQWordFromStack(c *CPU) [8]byte {
	val := [8]byte{
		GetFromAddressSpace(Address(c.SP)),
		GetFromAddressSpace(Address(c.SP + 1)),
		GetFromAddressSpace(Address(c.SP + 2)),
		GetFromAddressSpace(Address(c.SP + 3)),
		GetFromAddressSpace(Address(c.SP + 4)),
		GetFromAddressSpace(Address(c.SP + 5)),
		GetFromAddressSpace(Address(c.SP + 6)),
		GetFromAddressSpace(Address(c.SP + 7)),
	}
	c.SP += 8
	return val
}

func storeQword(value [8]byte, addr Address) {
	PutInAddressSpace(addr, value[0])
	PutInAddressSpace(addr+1, value[1])
	PutInAddressSpace(addr+2, value[2])
	PutInAddressSpace(addr+3, value[3])
	PutInAddressSpace(addr+4, value[4])
	PutInAddressSpace(addr+5, value[5])
	PutInAddressSpace(addr+6, value[6])
	PutInAddressSpace(addr+7, value[7])
}

func storeDword(value [4]byte, addr Address) {
	PutInAddressSpace(addr, value[0])
	PutInAddressSpace(addr+1, value[1])
	PutInAddressSpace(addr+2, value[2])
	PutInAddressSpace(addr+3, value[3])
}

func storeWord(value [2]byte, addr Address) {
	PutInAddressSpace(addr, value[0])
	PutInAddressSpace(addr+1, value[1])
}

func OpNoop(c *CPU) {}

func OpAddAToB(c *CPU) {
	c.A += c.B
}
func OpAddAbsA(c *CPU) {
	c.A += int64fromqword(getQword(c))
}

func OpSubAFromB(c *CPU) {
	c.A -= c.B
}

func OpSubAbsA(c *CPU) {
	c.A -= int64fromqword(getQword(c))
}
func OpMulAByB(c *CPU) {
	c.A *= c.B
}
func OpMulAbsA(c *CPU) {
	c.A += int64fromqword(getQword(c))
}

func OpDivAByB(c *CPU) {
	c.A /= c.B
}
func OpDivAbsA(c *CPU) {
	c.A /= int64fromqword(getQword(c))
}

func OpRolA(c *CPU) {
	highestBit := uint64(0x10_00_00_00_00_00_00_00)
	if c.A&highestBit == highestBit {
		c.Flags |= FLAG_CARRY
	}
	c.A = c.A << 1
}
func OpRolB(c *CPU) {
	highestBit := uint64(0x10_00_00_00_00_00_00_00)
	if c.B&highestBit == highestBit {
		c.Flags |= FLAG_CARRY
	}
	c.B = c.B << 1
}
func OpRolX(c *CPU) {
	highestBit := uint64(0x10_00_00_00_00_00_00_00)
	if c.B&highestBit == highestBit {
		c.Flags |= FLAG_CARRY
	}
	c.B = c.B << 1
}
func OpRolY(c *CPU) {
	highestBit := uint64(0x10_00_00_00_00_00_00_00)
	if c.Y&highestBit == highestBit {
		c.Flags |= FLAG_CARRY
	}
	c.Y = c.Y << 1
}

func OpXorAB(c *CPU) {
	c.A ^= c.B
}
func OpXorAX(c *CPU) {
	c.A ^= c.X
}

func OpFaddAToB(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) + math.Float64frombits(c.B))
}

func OpFaddAbsA(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) + math.Float64frombits(int64fromqword(getQword(c))))
}

func OpFSubAToB(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) - math.Float64frombits(c.B))
}

func OpFSubAbsA(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) - math.Float64frombits(int64fromqword(getQword(c))))
}

func OpFmulAByB(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) * math.Float64frombits(c.B))
}

func OpFmulAbsA(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) * math.Float64frombits(int64fromqword(getQword(c))))
}

func OpFDivAByB(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) / math.Float64frombits(c.B))
}
func OpFDivAbsA(c *CPU) {
	c.A = math.Float64bits(math.Float64frombits(c.A) / math.Float64frombits(int64fromqword(getQword(c))))
}

func OpStop(c *CPU) {
	c.Flags |= FLAG_HALT
}

// LEA
func OpLeaByteA(c *CPU) { c.A = uint64(GetFromAddressSpace(Address(c.A))) }
func OpLeaByteB(c *CPU) { c.B = uint64(GetFromAddressSpace(Address(c.B))) }
func OpLeaByteX(c *CPU) { c.X = uint64(GetFromAddressSpace(Address(c.X))) }
func OpLeaByteY(c *CPU) { c.Y = uint64(GetFromAddressSpace(Address(c.Y))) }

func OpLeaWordA(c *CPU) { c.A = int64fromword(getWordFrom(Address(c.A))) }
func OpLeaWordB(c *CPU) { c.B = int64fromword(getWordFrom(Address(c.B))) }
func OpLeaWordX(c *CPU) { c.X = int64fromword(getWordFrom(Address(c.X))) }
func OpLeaWordY(c *CPU) { c.Y = int64fromword(getWordFrom(Address(c.Y))) }

func OpLeaDWordA(c *CPU) { c.A = int64fromdword(getDwordFrom(Address(c.A))) }
func OpLeaDWordB(c *CPU) { c.B = int64fromdword(getDwordFrom(Address(c.B))) }
func OpLeaDWordX(c *CPU) { c.X = int64fromdword(getDwordFrom(Address(c.X))) }
func OpLeaDWordY(c *CPU) { c.Y = int64fromdword(getDwordFrom(Address(c.Y))) }

func OpLeaQWordA(c *CPU) { c.A = int64fromqword(getQwordFrom(Address(c.A))) }
func OpLeaQWordB(c *CPU) { c.B = int64fromqword(getQwordFrom(Address(c.B))) }
func OpLeaQWordX(c *CPU) { c.X = int64fromqword(getQwordFrom(Address(c.X))) }
func OpLeaQWordY(c *CPU) { c.Y = int64fromqword(getQwordFrom(Address(c.Y))) }

// SET
func OpSetByteA(c *CPU) { c.A = uint64(getByte(c)) }
func OpSetByteB(c *CPU) { c.B = uint64(getByte(c)) }
func OpSetByteX(c *CPU) { c.X = uint64(getByte(c)) }
func OpSetByteY(c *CPU) { c.Y = uint64(getByte(c)) }

func OpSetWordA(c *CPU) { c.A = int64fromword(getWord(c)) }
func OpSetWordB(c *CPU) { c.B = int64fromword(getWord(c)) }
func OpSetWordX(c *CPU) { c.X = int64fromword(getWord(c)) }
func OpSetWordY(c *CPU) { c.Y = int64fromword(getWord(c)) }

func OpSetDWordA(c *CPU) { c.A = int64fromdword(getDword(c)) }
func OpSetDWordB(c *CPU) { c.B = int64fromdword(getDword(c)) }
func OpSetDWordX(c *CPU) { c.X = int64fromdword(getDword(c)) }
func OpSetDWordY(c *CPU) { c.Y = int64fromdword(getDword(c)) }

func OpSetQWordA(c *CPU) { c.A = int64fromqword(getQword(c)) }
func OpSetQWordB(c *CPU) { c.B = int64fromqword(getQword(c)) }
func OpSetQWordX(c *CPU) { c.X = int64fromqword(getQword(c)) }
func OpSetQWordY(c *CPU) { c.Y = int64fromqword(getQword(c)) }

// STR
func OpStrByte(c *CPU)  { PutInAddressSpace(Address(c.B), byte(0xff&c.A)) }
func OpStrWord(c *CPU)  { storeWord(wordfromint64(c.A), Address(c.B)) }
func OpStrDWord(c *CPU) { storeDword(dwordfromint64(c.A), Address(c.B)) }
func OpStrQWord(c *CPU) { storeQword(qwordfromint64(c.A), Address(c.B)) }

// Jumps
func OpJsrAbs(c *CPU) {
	addQwordToStack(qwordfromint64(c.PC), c)
	c.PC = int64fromqword(getQword(c))
}

func OpJsrInd(c *CPU) {
	addQwordToStack(qwordfromint64(c.A), c)
	c.PC = int64fromqword(getQwordFrom(Address(c.A)))
}

// RET
func OpRet(c *CPU) {
	c.PC = int64fromqword(popQWordFromStack(c))
}

// PUSH

// Branches
func OpBps(c *CPU) {
	if int64(c.A) >= 0 {
		c.PC = int64fromqword(getQword(c))
	}
}
func OpBng(c *CPU) {
	if int64(c.A) < 0 {
		c.PC = int64fromqword(getQword(c))
	}
}

func OpBcc(c *CPU) {
	if c.Flags&FLAG_CARRY == 0 {
		c.PC = int64fromqword(getQword(c))
	}
}

func OpBcs(c *CPU) {
	if c.Flags&FLAG_CARRY == 1 {
		c.PC = int64fromqword(getQword(c))
	}
}
func OpBneAB(c *CPU) {
	if c.A != c.B {
		c.PC = int64fromqword(getQword(c))
	}
}
func OpBeqAB(c *CPU) {
	if c.A == c.B {
		c.PC = int64fromqword(getQword(c))
	}
}

func OpBneAbsA(c *CPU) {
	if c.A != int64fromqword(getQword(c)) {
		c.PC = int64fromqword(getQword(c))
	}
}
func OpBeqAbsA(c *CPU) {
	if c.A == int64fromqword(getQword(c)) {
		c.PC = int64fromqword(getQword(c))
	}
}

// Transfer
func OpTrAB(c *CPU) { c.B = c.A }
func OpTrAX(c *CPU) { c.X = c.A }
func OpTrAY(c *CPU) { c.Y = c.A }

func OpTrBA(c *CPU) { c.A = c.B }
func OpTrBX(c *CPU) { c.X = c.B }
func OpTrBY(c *CPU) { c.Y = c.B }

func OpTrXA(c *CPU) { c.A = c.X }
func OpTrXB(c *CPU) { c.B = c.X }
func OpTrXY(c *CPU) { c.X = c.Y }

func OpTrYA(c *CPU) { c.A = c.Y }
func OpTrYB(c *CPU) { c.B = c.Y }
func OpTrYX(c *CPU) { c.X = c.Y }

// Decs
func OpDecA(c *CPU) { c.A-- }
func OpDecB(c *CPU) { c.B-- }
func OpDecX(c *CPU) { c.X-- }
func OpDecY(c *CPU) { c.Y-- }

// Incs
func OpIncA(c *CPU) { c.A++ }
func OpIncB(c *CPU) { c.B++ }
func OpIncX(c *CPU) { c.X++ }
func OpIncY(c *CPU) { c.Y++ }

func OpHalt(c *CPU) {
	c.Flags |= FLAG_HALT
}

func init() {
	// Jumptable
	Instructions = map[op.Opcode]Operation{
		op.OP_NOOP: OpNoop,

		op.OP_ADD_A_TO_B:   OpAddAToB,
		op.OP_ADD_ABS_A:    OpAddAbsA,
		op.OP_SUB_A_FROM_B: OpSubAFromB,
		op.OP_SUB_ABS_A:    OpSubAbsA,
		op.OP_MUL_A_BY_B:   OpMulAByB,
		op.OP_MUL_ABS_A:    OpMulAbsA,
		op.OP_DIV_A_BY_B:   OpDivAByB,
		op.OP_DIV_ABS_A:    OpDivAbsA,

		op.OP_ROL_A: OpRolA,
		op.OP_ROL_B: OpRolB,
		op.OP_ROL_X: OpRolX,
		op.OP_ROL_Y: OpRolY,

		op.OP_XOR_A_B: OpXorAB,
		op.OP_XOR_A_X: OpXorAX,

		op.OP_FADD_A_TO_B: OpFaddAToB,
		op.OP_FADD_ABS_A:  OpFaddAbsA,
		op.OP_FSUB_A_TO_B: OpFSubAToB,
		op.OP_FSUB_ABS_A:  OpFSubAbsA,
		op.OP_FMUL_A_BY_B: OpFmulAByB,
		op.OP_FMUL_ABS_A:  OpFmulAbsA,
		op.OP_FDIV_A_BY_B: OpFDivAByB,
		op.OP_FDIV_ABS_A:  OpFDivAbsA,

		op.OP_LEA_BYTE_A: OpLeaByteA,
		op.OP_LEA_BYTE_B: OpLeaByteB,
		op.OP_LEA_BYTE_X: OpLeaByteX,
		op.OP_LEA_BYTE_Y: OpLeaByteY,

		op.OP_LEA_WORD_A: OpLeaWordA,
		op.OP_LEA_WORD_B: OpLeaWordB,
		op.OP_LEA_WORD_X: OpLeaWordX,
		op.OP_LEA_WORD_Y: OpLeaWordY,

		op.OP_LEA_DWORD_A: OpLeaDWordA,
		op.OP_LEA_DWORD_B: OpLeaDWordB,
		op.OP_LEA_DWORD_X: OpLeaDWordX,
		op.OP_LEA_DWORD_Y: OpLeaDWordY,

		op.OP_LEA_QWORD_A: OpLeaQWordA,
		op.OP_LEA_QWORD_B: OpLeaQWordB,
		op.OP_LEA_QWORD_X: OpLeaQWordX,
		op.OP_LEA_QWORD_Y: OpLeaQWordY,

		op.OP_SET_BYTE_A: OpSetByteA,
		op.OP_SET_BYTE_B: OpSetByteB,
		op.OP_SET_BYTE_X: OpSetByteX,
		op.OP_SET_BYTE_Y: OpSetByteY,

		op.OP_SET_WORD_A: OpSetWordA,
		op.OP_SET_WORD_B: OpSetWordB,
		op.OP_SET_WORD_X: OpSetWordX,
		op.OP_SET_WORD_Y: OpSetWordY,

		op.OP_SET_DWORD_A: OpSetDWordA,
		op.OP_SET_DWORD_B: OpSetDWordB,
		op.OP_SET_DWORD_X: OpSetDWordX,
		op.OP_SET_DWORD_Y: OpSetDWordY,

		op.OP_SET_QWORD_A: OpSetQWordA,
		op.OP_SET_QWORD_B: OpSetQWordB,
		op.OP_SET_QWORD_X: OpSetQWordX,
		op.OP_SET_QWORD_Y: OpSetQWordY,

		op.OP_STR_BYTE:  OpStrByte,
		op.OP_STR_WORD:  OpStrWord,
		op.OP_STR_DWORD: OpStrDWord,
		op.OP_STR_QWORD: OpStrQWord,

		op.OP_JSR_ABS: OpJsrAbs,
		op.OP_JSR_IND: OpJsrInd,

		op.OP_BPS:       OpBps,
		op.OP_BNG:       OpBng,
		op.OP_BCC:       OpBcc,
		op.OP_BCS:       OpBcs,
		op.OP_BNE_A_B:   OpBneAB,
		op.OP_BEQ_A_B:   OpBeqAbsA,
		op.OP_BNE_A_ABS: OpBneAbsA,
		op.OP_BEQ_A_ABS: OpBeqAbsA,

		op.OP_TR_A_TO_B: OpTrAB,
		op.OP_TR_A_TO_X: OpTrAX,
		op.OP_TR_A_TO_Y: OpTrAY,

		op.OP_TR_B_TO_A: OpTrBA,
		op.OP_TR_B_TO_X: OpTrBX,
		op.OP_TR_B_TO_Y: OpTrBY,

		op.OP_TR_X_TO_A: OpTrXA,
		op.OP_TR_X_TO_B: OpTrXB,
		op.OP_TR_X_TO_Y: OpTrXY,

		op.OP_TR_Y_TO_A: OpTrYA,
		op.OP_TR_Y_TO_B: OpTrYB,
		op.OP_TR_Y_TO_X: OpTrYX,

		op.OP_DEC_A: OpDecA,
		op.OP_DEC_B: OpDecB,
		op.OP_DEC_X: OpDecX,
		op.OP_DEC_Y: OpDecY,

		op.OP_INC_A: OpIncA,
		op.OP_INC_B: OpIncB,
		op.OP_INC_X: OpIncX,
		op.OP_INC_Y: OpIncY,

		op.OP_PUSH_BYTE_A: OpPushByteA,
		op.OP_PUSH_BYTE_B: OpPushByteB,
		op.OP_PUSH_BYTE_X: OpPushByteX,
		op.OP_PUSH_BYTE_Y: OpPushByteY,

		op.OP_POP_BYTE_A: OpPopByteA,
		op.OP_POP_BYTE_B: OpPopByteB,
		op.OP_POP_BYTE_X: OpPopByteX,
		op.OP_POP_BYTE_Y: OpPopByteY,

		op.OP_PUSH_WORD_A: OpPushWordA,
		op.OP_PUSH_WORD_B: OpPushWordB,
		op.OP_PUSH_WORD_X: OpPushWordX,
		op.OP_PUSH_WORD_Y: OpPushWordY,

		op.OP_POP_WORD_A: OpPopWordY,
		op.OP_POP_WORD_B: OpPopWordY,
		op.OP_POP_WORD_X: OpPopWordY,
		op.OP_POP_WORD_Y: OpPopWordY,

		op.OP_PUSH_DWORD_A: OpPushDWordA,
		op.OP_PUSH_DWORD_B: OpPushDWordA,
		op.OP_PUSH_DWORD_X: OpPushDWordX,
		op.OP_PUSH_DWORD_Y: OpPushDWordY,

		op.OP_POP_DWORD_A: OpPopDWordA,
		op.OP_POP_DWORD_B: OpPopDWordY,
		op.OP_POP_DWORD_X: OpPopDWordY,
		op.OP_POP_DWORD_Y: OpPopDWordY,

		op.OP_PUSH_QWORD_A: OpPushQWordA,
		op.OP_PUSH_QWORD_B: OpPushQWordB,
		op.OP_PUSH_QWORD_X: OpPushQWordX,
		op.OP_PUSH_QWORD_Y: OpPushQWordY,

		op.OP_POP_QWORD_A: OpPopQWordY,
		op.OP_POP_QWORD_B: OpPopQWordY,
		op.OP_POP_QWORD_X: OpPopQWordY,
		op.OP_POP_QWORD_Y: OpPopQWordY,

		op.OP_HALT: OpHalt,
	}
}
