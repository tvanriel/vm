package codegen

import (
	op "github.com/tvanriel/vm/operations"
)

func (i *Instruction) Bytes() []byte {
	switch i.Name {

	case "nop":
		return opcodetobytes(op.OP_NOOP)

	case "add":
		if i.Arguments[0].Type == RegisterArgument && i.Arguments[0].Value == uint64(REGISTER_B) {
			return opcodetobytes(op.OP_ADD_A_TO_B)
		}
		return opcodeqwordtobytes(op.OP_ADD_A_TO_B, i.Arguments[0].Value)

	case "sub":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_SUB_B_FROM_A)
			}

		}
		return opcodeqwordtobytes(op.OP_SUB_ABS_A, i.Arguments[0].Value)
	case "mul":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_MUL_A_BY_B)
			}
		}
		return opcodeqwordtobytes(op.OP_DIV_ABS_A, i.Arguments[0].Value)
	case "div":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_DIV_A_BY_B)
			}
		}
		return opcodeqwordtobytes(op.OP_MUL_ABS_A, i.Arguments[0].Value)

	case "rola":
		return opcodetobytes(op.OP_ROL_A)
	case "rolb":
		return opcodetobytes(op.OP_ROL_B)
	case "rolx":
		return opcodetobytes(op.OP_ROL_X)
	case "roly":
		return opcodetobytes(op.OP_ROL_Y)

	case "rora":
		return opcodetobytes(op.OP_ROR_A)
	case "rorb":
		return opcodetobytes(op.OP_ROR_B)
	case "rorx":
		return opcodetobytes(op.OP_ROR_X)
	case "rory":
		return opcodetobytes(op.OP_ROR_Y)

	case "rol":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_ROL_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_ROL_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_ROL_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_ROL_Y)
			}
		}

	case "ror":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_ROR_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_ROR_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_ROR_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_ROR_Y)
			}
		}
	case "xor":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_XOR_A_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_XOR_A_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_XOR_A_Y)
			}
		}

	case "fadd":
		if i.Arguments[0].Type == RegisterArgument && i.Arguments[0].Value == uint64(REGISTER_B) {
			return opcodetobytes(op.OP_FADD_A_TO_B)
		}
		return opcodeqwordtobytes(op.OP_FADD_A_TO_B, i.Arguments[0].Value)

	case "fsub":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_FSUB_B_FROM_A)
			}

		}
		return opcodeqwordtobytes(op.OP_FSUB_ABS_A, i.Arguments[0].Value)
	case "fmul":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_FMUL_A_BY_B)
			}
		}
		return opcodeqwordtobytes(op.OP_FMUL_ABS_A, i.Arguments[0].Value)
	case "fdiv":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_FDIV_A_BY_B)
			}
		}
		return opcodeqwordtobytes(op.OP_FDIV_ABS_A, i.Arguments[0].Value)

	case "lda":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodeqwordtobytes(op.OP_SET_QWORD_A, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_TR_B_TO_A)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_TR_X_TO_A)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_TR_Y_TO_A)
			}
		}
	case "ldb":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodeqwordtobytes(op.OP_SET_QWORD_B, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_TR_A_TO_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_TR_X_TO_B)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_TR_Y_TO_B)
			}
		}
	case "ldx":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodeqwordtobytes(op.OP_SET_QWORD_X, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_TR_A_TO_X)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_TR_B_TO_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_TR_Y_TO_X)
			}
		}
	case "ldy":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodeqwordtobytes(op.OP_SET_QWORD_Y, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_TR_A_TO_Y)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_TR_X_TO_Y)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_TR_B_TO_Y)
			}
		}
	case "leab":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_LEA_BYTE_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_LEA_BYTE_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_LEA_BYTE_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_LEA_BYTE_Y)
			}
		}
	case "leaw":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_LEA_WORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_LEA_WORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_LEA_WORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_LEA_WORD_Y)
			}
		}
	case "lead":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_LEA_DWORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_LEA_DWORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_LEA_DWORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_LEA_DWORD_Y)
			}
		}
	case "leaq":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_LEA_QWORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_LEA_QWORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_LEA_QWORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_LEA_QWORD_Y)
			}
		}
	case "stab":
		return opcodebytetobytes(op.OP_STR_BYTE, i.Arguments[0].Value)
	case "staw":
		return opcodewordtobytes(op.OP_STR_WORD, i.Arguments[0].Value)
	case "stad":
		return opcodedwordtobytes(op.OP_STR_DWORD, i.Arguments[0].Value)
	case "staq":
		return opcodeqwordtobytes(op.OP_STR_QWORD, i.Arguments[0].Value)

	case "jsr":
		if i.Arguments[0].Type == LabelArgument {
			return opcodeqwordtobytes(op.OP_JSR_ABS, 0)
		}
		return opcodeqwordtobytes(op.OP_JSR_ABS, i.Arguments[0].Value)
	case "jmp":
		if i.Arguments[0].Type == LabelArgument {
			return opcodeqwordtobytes(op.OP_JSR_ABS, 0)
		}
		return opcodeqwordtobytes(op.OP_JSR_IND, i.Arguments[0].Value)

	case "ret":
		return opcodetobytes(op.OP_RET)

	case "bps":
		return opcodetobytes(op.OP_BPS)
	case "bng":
		return opcodetobytes(op.OP_BNG)
	case "bcc":
		return opcodetobytes(op.OP_BCC)
	case "bcs":
		return opcodetobytes(op.OP_BCS)
	case "beq":
		if len(i.Arguments) == 2 {
			return append(opcodeqwordtobytes(op.OP_BEQ_A_ABS, i.Arguments[0].Value), qwordfromint64(i.Arguments[1].Value)...)
		}
		return opcodetobytes(op.OP_BEQ_A_B)
	case "bne":
		if len(i.Arguments) == 2 {
			return append(opcodeqwordtobytes(op.OP_BNE_A_ABS, i.Arguments[0].Value), qwordfromint64(i.Arguments[1].Value)...)
		}
		return opcodetobytes(op.OP_BNE_A_B)

	case "dec":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_DEC_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_DEC_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_DEC_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_DEC_Y)
			}
		}
	case "inc":
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_INC_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_INC_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_INC_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_INC_Y)
			}
		}
	case "popb":

	case "popw":

	case "popd":

	case "popq":

	case "pushb":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodebytetobytes(op.OP_PUSH_ABS_BYTE, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_PUSH_BYTE_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_PUSH_BYTE_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_PUSH_BYTE_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_PUSH_BYTE_Y)
			}
		}
	case "pushw":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodewordtobytes(op.OP_PUSH_ABS_WORD, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_PUSH_WORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_PUSH_WORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_PUSH_WORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_PUSH_WORD_Y)
			}
		}
	case "pushd":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodedwordtobytes(op.OP_PUSH_ABS_DWORD, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_PUSH_DWORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_PUSH_DWORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_PUSH_DWORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_PUSH_DWORD_Y)
			}
		}
	case "pushq":
		if i.Arguments[0].Type == ConstantArgument {
			return opcodeqwordtobytes(op.OP_PUSH_ABS_QWORD, i.Arguments[0].Value)
		}
		if i.Arguments[0].Type == RegisterArgument {
			switch i.Arguments[0].Value {
			case uint64(REGISTER_A):
				return opcodetobytes(op.OP_PUSH_QWORD_A)
			case uint64(REGISTER_B):
				return opcodetobytes(op.OP_PUSH_QWORD_B)
			case uint64(REGISTER_X):
				return opcodetobytes(op.OP_PUSH_QWORD_X)
			case uint64(REGISTER_Y):
				return opcodetobytes(op.OP_PUSH_QWORD_Y)
			}
		}
	case "halt":
		opcodetobytes(op.OP_HALT)
	}

	return []byte{}
}

func opcodetobytes(o op.Opcode) []byte {
	return []byte{byte(o)}
}

func opcodeqwordtobytes(o op.Opcode, qw uint64) []byte {
	return byteqwordtobytes(byte(o), qw)
}

func byteqwordtobytes(b byte, qw uint64) []byte {
	return append([]byte{b}, qwordfromint64(qw)...)
}

func opcodedwordtobytes(o op.Opcode, dw uint64) []byte {
	return bytedwordtobytes(byte(o), dw)
}

func bytedwordtobytes(b byte, qw uint64) []byte {
	return append([]byte{b}, dwordfromint64(qw)...)
}

func opcodewordtobytes(o op.Opcode, dw uint64) []byte {
	return bytedwordtobytes(byte(o), dw)
}

func bytewordtobytes(b byte, qw uint64) []byte {
	return append([]byte{b}, wordfromint64(qw)...)
}

func opcodebytetobytes(o op.Opcode, dw uint64) []byte {
	return bytedwordtobytes(byte(o), dw)
}

func bytebytetobytes(b byte, qw uint64) []byte {
	return append([]byte{b}, bytefromint64(qw)...)
}
