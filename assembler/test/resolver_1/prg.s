
#include "test/resolver_1/include.s"

main:
    lda 1
    sta DISPLAY_ENABLE ;comment
    lda (1 + 1)
    lda (1 - 1)
    pusha
    popa
    jmp label

label:
    test 1
    bne 123
