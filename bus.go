package vm

import (
	"fmt"
	"sort"
	"sync"
)

type Address uint64

type AddressRange [2]Address

func (r AddressRange) IsInRange(addr Address) bool {
	lowerLimit := addr >= r[0]
	upperLimit := addr <= r[1]
	return upperLimit && lowerLimit

}

type Device interface {
	Range() AddressRange
	Put(addr Address, val byte)
	Get(addr Address) byte
	Name() string
}

var devicesMutex sync.Mutex
var devices []Device

func GetFromAddressSpace(addr Address) byte {
	fmt.Printf("Bus: GET 0x%x", addr)

	for _, device := range devices {
		if device.Range().IsInRange(addr) {
			fmt.Printf(" -> %s\n", device.Name())
			return device.Get(addr)
		}
	}
	fmt.Printf(" -> NULL\n")

	// NULL-Device
	return 0x00
}

func PutInAddressSpace(addr Address, val byte) {
	for _, device := range devices {
		if device.Range().IsInRange(addr) {
			device.Put(addr, val)
		}
	}

}

func RegisterDevice(device Device) {
	devicesMutex.Lock()
	devices = append(devices, device)
	sort.Slice(devices, func(i, j int) bool {
		return devices[i].Range()[0] < devices[j].Range()[0]
	})
	devicesMutex.Unlock()
}
