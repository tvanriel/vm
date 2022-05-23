; a comment here
.org 0x10000
    lda 0xDEADBEEFDEADBEEF
    ldb 0xC0FFEEC0FFEEC0FF
    ldx 0b0101010101010101
    ldy 012345671234567123
    jmp label
label:


.org 0x1fff6
    jmp 0x10000
