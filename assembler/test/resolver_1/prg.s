
#include "test/parser_1/include.s"

main:
    lda 1
    sta DISPLAY_ENABLE ;comment
    lda (1 + 1), P
    lda (1 - 1)
    pusha
    popa

label:
    test 1
    bne 123