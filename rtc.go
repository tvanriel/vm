package vm

import "time"

const RTC_ADDR_START Address = 0x2_01_00
const RTC_ADDR_END Address = 0x2_01_09

const RTC_ADDR_FREEZE Address = 0x2_01_09

const RTC_ADDR_CLOCK_START Address = 0x2_01_00
const RTC_ADDR_CLOCK_END Address = 0x2_01_08

type ClockDevice struct {
	cur uint64 // Snapshotted timestamp
}

func (c *ClockDevice) Get(addr Address) byte {
	if addr == RTC_ADDR_FREEZE {
		return 0x00 // Write-only.
	}

	return qwordfromint64(c.cur)[addr-RTC_ADDR_CLOCK_START]

}
func (c *ClockDevice) Put(addr Address, val byte) {
	// update timestamp.
	if addr == RTC_ADDR_FREEZE {
		c.cur = uint64(time.Now().Unix())
	}
}
func (c *ClockDevice) Range() AddressRange {
	return AddressRange{RTC_ADDR_START, RTC_ADDR_END}
}

func (*ClockDevice) Name() string { return "Clock" }

var RealTimeClock Device = &ClockDevice{}

func init() {
	RegisterDevice(RealTimeClock)
}
