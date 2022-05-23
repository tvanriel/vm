package main

import "fmt"

const TTY_ADDR_START Address = 0x2_03_00
const TTY_ADDR_END Address = 0x2_03_10
const TTY_ADDR_BUF_APP Address = 0x2_03_01
const TTY_ADDR_BUF_FLUSH Address = 0x2_03_02

type TTYDevice struct {
	buf   [1024]byte
	index int
}

func (m *TTYDevice) Get(addr Address) byte {
	return 0x00 // write-only
}

func (m *TTYDevice) Put(addr Address, val byte) {
	if addr == TTY_ADDR_BUF_APP {
		m.buf[m.index] = val
	}
	if addr == TTY_ADDR_BUF_FLUSH {
		fmt.Print(string(m.buf[:m.index]))
	}
}

func (m *TTYDevice) Range() AddressRange {
	return AddressRange{TTY_ADDR_START, TTY_ADDR_END}
}

func (*TTYDevice) Name() string { return "TTY" }

var Tty1 Device = &TTYDevice{}

func init() {
	RegisterDevice(Tty1)
}
