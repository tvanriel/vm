package operations

type Opcode byte

const (
	OP_NOOP Opcode = iota
	OP_STOP        // Halts the execution

	// Integer math operations
	OP_ADD_A_TO_B   // A += B
	OP_ADD_ABS_A    // A += (arg)
	OP_SUB_A_FROM_B // A -= B
	OP_SUB_ABS_A    // A -= (arg)
	OP_MUL_A_BY_B   // A *= B
	OP_MUL_ABS_A    // A *= (arg)
	OP_DIV_A_BY_B   // A /= B
	OP_DIV_ABS_A    // A /= (arg)

	// Bitwise operations.

	// Rotate-left.  Leftmost bit sets the C flag.
	OP_ROL_A // A = A << 1
	OP_ROL_B // B = B << 1
	OP_ROL_X // X = X << 1
	OP_ROL_Y // Y = Y << 1

	OP_XOR_A_B // A^=B
	OP_XOR_A_X // A^=X

	// Floating point math operations
	OP_FADD_A_TO_B // A += B
	OP_FADD_ABS_A  // A += (arg)
	OP_FSUB_A_TO_B // A -= B
	OP_FSUB_ABS_A  // A -= (arg)
	OP_FMUL_A_BY_B // A *= B
	OP_FMUL_ABS_A  // A *= (arg)
	OP_FDIV_A_BY_B // A /= B
	OP_FDIV_ABS_A  // A /= (arg)

	// Memory operations

	// Load address
	OP_LEA_BYTE_A // A[:1] = Mem[A]
	OP_LEA_BYTE_B // B[:1] = Mem[B]
	OP_LEA_BYTE_X // X[:1] = Mem[X]
	OP_LEA_BYTE_Y // Y[:1] = Mem[Y]

	OP_LEA_WORD_A // A[:2] = Mem[A:A+2]
	OP_LEA_WORD_B // B[:2] = Mem[B:B+2]
	OP_LEA_WORD_X // X[:2] = Mem[X:X+2]
	OP_LEA_WORD_Y // Y[:2] = Mem[Y:Y+2]

	OP_LEA_DWORD_A // A[:4] = Mem[A:A+4]
	OP_LEA_DWORD_B // B[:4] = Mem[B:B+4]
	OP_LEA_DWORD_X // X[:4] = Mem[X:X+4]
	OP_LEA_DWORD_Y // Y[:4] = Mem[Y:Y+4]

	OP_LEA_QWORD_A // A = Mem[A:A+8]
	OP_LEA_QWORD_B // B = Mem[B:B+8]
	OP_LEA_QWORD_X // X = Mem[X:X+8]
	OP_LEA_QWORD_Y // Y = Mem[Y:Y+8]

	// SET-operations - sets register = address at PC
	OP_SET_BYTE_A // A[:1] = Mem[PC+1]
	OP_SET_BYTE_B // B[:1] = Mem[PC+1]
	OP_SET_BYTE_X // X[:1] = Mem[PC+1]
	OP_SET_BYTE_Y // Y[:1] = Mem[PC+1]

	OP_SET_WORD_A // A[:2] = Mem[PC+1:PC+2]
	OP_SET_WORD_B // B[:2] = Mem[PC+1:PC+2]
	OP_SET_WORD_X // X[:2] = Mem[PC+1:PC+2]
	OP_SET_WORD_Y // Y[:2] = Mem[PC+1:PC+2]

	OP_SET_DWORD_A // A[:4] = Mem[PC+1:PC+4]
	OP_SET_DWORD_B // B[:4] = Mem[PC+1:PC+4]
	OP_SET_DWORD_X // X[:4] = Mem[PC+1:PC+4]
	OP_SET_DWORD_Y // Y[:4] = Mem[PC+1:PC+4]

	OP_SET_QWORD_A // A = Mem[PC+1:PC+8]
	OP_SET_QWORD_B // B = Mem[PC+1:PC+8]
	OP_SET_QWORD_X // X = Mem[PC+1:PC+8]
	OP_SET_QWORD_Y // Y = Mem[PC+1:PC+8]

	// Store address
	OP_STR_BYTE  // Mem[B] = A[:1]
	OP_STR_WORD  // Mem[B:B+2] = A[:2]
	OP_STR_DWORD // Mem[B:B+4] = A[:4]
	OP_STR_QWORD // Mem[B:B+8] = A

	// Jumps
	OP_JSR_ABS // Jump to (arg)
	OP_JSR_IND // Jump to address contained in (arg)

	OP_RET // Return

	// Branches
	OP_BPS       // Branch if A & 1 == 0
	OP_BNG       // Branch if A & 1 != 0
	OP_BCC       // Branch if C flag is not set
	OP_BCS       // Branch if C flag is set
	OP_BNE_A_B   // Branch if A != B
	OP_BEQ_A_B   // Branch if A == B
	OP_BNE_A_ABS // Branch if A != (arg)
	OP_BEQ_A_ABS // Branch if A == (arg)

	// Register transfer operations
	OP_TR_A_TO_B // B = A
	OP_TR_A_TO_X // X = A
	OP_TR_A_TO_Y // Y = A

	OP_TR_B_TO_A // A = B
	OP_TR_B_TO_X // X = B
	OP_TR_B_TO_Y // Y = B

	OP_TR_X_TO_A // A = X
	OP_TR_X_TO_B // B = X
	OP_TR_X_TO_Y // Y = X

	OP_TR_Y_TO_A // A = Y
	OP_TR_Y_TO_B // B = Y
	OP_TR_Y_TO_X // X = Y

	// Register decrement operations
	OP_DEC_A // A--
	OP_DEC_B // B--
	OP_DEC_X // X--
	OP_DEC_Y // Y--

	// Register increment operations
	OP_INC_A // A++
	OP_INC_B // B++
	OP_INC_X // X++
	OP_INC_Y // Y++

	// Stack operations
	OP_PUSH_BYTE_A // Push val of A on the stack
	OP_PUSH_BYTE_B // Push val of B on the stack
	OP_PUSH_BYTE_X // Push val of X on the stack
	OP_PUSH_BYTE_Y // Push val of Y on the stack

	OP_POP_BYTE_A // Pop top of stack onto A
	OP_POP_BYTE_B // Pop top of stack onto B
	OP_POP_BYTE_X // Pop top of stack onto X
	OP_POP_BYTE_Y // Pop top of stack onto Y

	OP_PUSH_WORD_A // Push val of A on the stack
	OP_PUSH_WORD_B // Push val of B on the stack
	OP_PUSH_WORD_X // Push val of X on the stack
	OP_PUSH_WORD_Y // Push val of Y on the stack

	OP_POP_WORD_A // Pop top of stack onto A
	OP_POP_WORD_B // Pop top of stack onto B
	OP_POP_WORD_X // Pop top of stack onto X
	OP_POP_WORD_Y // Pop top of stack onto Y

	OP_PUSH_DWORD_A // Push val of A on the stack
	OP_PUSH_DWORD_B // Push val of B on the stack
	OP_PUSH_DWORD_X // Push val of X on the stack
	OP_PUSH_DWORD_Y // Push val of Y on the stack

	OP_POP_DWORD_A // Pop top of stack onto A
	OP_POP_DWORD_B // Pop top of stack onto B
	OP_POP_DWORD_X // Pop top of stack onto X
	OP_POP_DWORD_Y // Pop top of stack onto Y

	OP_PUSH_QWORD_A // Push val of A on the stack
	OP_PUSH_QWORD_B // Push val of B on the stack
	OP_PUSH_QWORD_X // Push val of X on the stack
	OP_PUSH_QWORD_Y // Push val of Y on the stack

	OP_POP_QWORD_A // Pop top of stack onto A
	OP_POP_QWORD_B // Pop top of stack onto B
	OP_POP_QWORD_X // Pop top of stack onto X
	OP_POP_QWORD_Y // Pop top of stack onto Y

	OP_HALT
)
