package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const DEBUG_ADDR_START Address = 0x2_10_00
const DEBUG_ADDR_END Address = 0x2_10_10
const DEBUG_ADDR_ENABLE Address = 0x2_10_00

type DebuggerDevice struct {
	app     *tview.Application
	running bool
}

func (m *DebuggerDevice) Get(addr Address) byte {
	return 0x00 // write-only
}

func (m *DebuggerDevice) Put(addr Address, val byte) {
	if addr == DEBUG_ADDR_START {
		if val != 0 {
			go m.enable()
			return
		}
		m.disable()

	}
}

func (m *DebuggerDevice) enable() {
	m.app = tview.NewApplication()
	m.setupLayout()
	m.app.Run()
}
func (m *DebuggerDevice) disable() {
	m.app.Stop()
}

func (m *DebuggerDevice) Range() AddressRange {
	return AddressRange{DEBUG_ADDR_START, TTY_ADDR_END}
}

func (m *DebuggerDevice) setupLayout() {
	m.app.SetRoot(
		tview.NewFlex().
			AddItem(m.ramWatcher(), 0, 1, false).
			AddItem(m.romDumper(), 0, 1, false),
		true,
	)

}

func (m *DebuggerDevice) ramWatcher() tview.Primitive {
	view := tview.NewTextView()
	view.SetWordWrap(true).
		SetChangedFunc(func() { m.app.Draw() }).
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[cyan]RAM Watcher")
	go func() {
		for m.running {
			fmt.Fprint(view, memDump(Memory))
			time.Sleep(30 * time.Millisecond)
		}
	}()
	return view
}

func (m *DebuggerDevice) romDumper() tview.Primitive {
	view := tview.NewTextView()
	view.SetWordWrap(true).
		SetChangedFunc(func() { m.app.Draw() }).
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("[cyan]RAM Watcher")
	fmt.Fprint(view, romDump(Rom))
	time.Sleep(30 * time.Millisecond)
	return view
}

// Get a string representing the current state of the ram.
func memDump(ram *MemDevice) string {
	var sb strings.Builder
	for i := 0; i < int(MEM_SIZE_MAX); i++ {
		sb.WriteString("0x")
		sb.WriteString(strconv.Itoa(int(ram[i])))
		sb.WriteString(" ")
	}
	return sb.String()
}

// Get a string representing the content of the rom.
func romDump(ram *RomDevice) string {
	var sb strings.Builder
	for i := 0; i < int(ROM_SIZE_MAX); i++ {
		sb.WriteString("0x")
		sb.WriteString(strconv.Itoa(int(ram[i])))
		sb.WriteString(" ")
	}
	return sb.String()
}
func (d *DebuggerDevice) Name() string { return "Debugger" }

var Debugger Device = &DebuggerDevice{}

func init() {
	RegisterDevice(Debugger)
}
