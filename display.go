package vm

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

const DISPLAY_ADDR_START Address = 0x2_02_00
const DISPLAY_ADDR_END Address = 0x2_02_10

const DISPLAY_ADDR_X1 Address = 0x2_01_00
const DISPLAY_ADDR_X2 Address = 0x2_01_01

const DISPLAY_ADDR_Y1 Address = 0x2_01_02
const DISPLAY_ADDR_Y2 Address = 0x2_01_03

const DISPLAY_ADDR_R Address = 0x2_01_04
const DISPLAY_ADDR_G Address = 0x2_01_05
const DISPLAY_ADDR_B Address = 0x2_01_06

// Set-To-Write
const DISPLAY_ADDR_W Address = 0x2_01_07
const DISPLAY_ADDR_ENABLE Address = 0x2_01_08

const DISPLAY_SIZE_X = 1280
const DISPLAY_SIZE_Y = 720

type Monitor struct {
	cur_r, cur_g, cur_b byte
	cur_x1, cur_x2      byte
	cur_y1, cur_y2      byte

	running bool
	window  *sdl.Window
	surface *sdl.Surface
}

func (c *Monitor) Get(addr Address) byte {
	if addr == DISPLAY_ADDR_W {
		return 0x00 // Write-only.
	}

	if addr == DISPLAY_ADDR_ENABLE {
		if c.running {
			return 0x01
		}
		return 0x00
	}

	return 0x00
}
func (c *Monitor) Put(addr Address, val byte) {
	// update timestamp.
	if addr == DISPLAY_ADDR_W {
		c.putPixel()
	}
	if addr == DISPLAY_ADDR_ENABLE {
		if val != 0 {
			go c.enable()
			return
		}
		c.disable()
	}
}

func (c *Monitor) enable() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var err error
	c.window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		DISPLAY_SIZE_X, DISPLAY_SIZE_Y, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer c.window.Destroy()

	c.surface, err = c.window.GetSurface()
	if err != nil {
		panic(err)
	}

	// Blank out the screen.
	c.surface.FillRect(nil, 0)

	c.window.UpdateSurface()

	c.running = true
	for c.running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				c.running = false
			}
		}
	}
}

func (c *Monitor) putPixel() {

	color := color.RGBA{
		R: c.cur_r,
		G: c.cur_g,
		B: c.cur_b,
		A: byte(255),
	}

	xPos := int((int(c.cur_x1) << 8) + int(c.cur_x2))
	yPos := int((int(c.cur_y1) << 8) + int(c.cur_y2))

	c.surface.Set(xPos, yPos, color)
}

func (c *Monitor) disable() {
	c.running = false
}

func (c *Monitor) Range() AddressRange {
	return AddressRange{DISPLAY_ADDR_START, DISPLAY_ADDR_END}
}
func (*Monitor) Name() string { return "Monitor" }

var Display Device = &Monitor{}

func init() {
	RegisterDevice(Display)
}
