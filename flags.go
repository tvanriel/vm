package main

type Flag int64

const (
	FLAG_ZERO Flag = iota << 1
	FLAG_CARRY
	FLAG_HALT
)