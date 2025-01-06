package keyboard

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func NewGenericKeyboard() VirtualKeyboard {
	keys := make([]linux.Key, 0)

	for key := linux.KEY_RESERVED + 1; key <= linux.KEY_MICMUTE; key++ {
		keys = append(keys, key)
	}

	return NewVirtualKeyboardFactory().
		WithDevice(
			virtual_device.NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(0xDEAD).
				WithProduct(0xBABE).
				WithVersion(0x01).
				WithName("Generic Keyboard"),
		).
		WithMiscEvents([]linux.MiscEvent{linux.MSC_SCAN}).
		WithRepeat(250, 33).
		WithLEDs(
			[]linux.Led{
				linux.LED_NUML,
				linux.LED_CAPSL,
			},
		).
		WithKeys(keys).
		Create()
}
