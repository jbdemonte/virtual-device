package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewJoyConL() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_BLUETOOTH).
				WithVendor(sdl.USB_VENDOR_NINTENDO).
				WithProduct(sdl.USB_PRODUCT_NINTENDO_SWITCH_JOYCON_LEFT).
				WithVersion(0x8001).
				WithName("Joy-Con (L)"),
		).
		WithDigital(
			MappingDigital{
				ButtonUp:    linux.BTN_DPAD_UP,
				ButtonRight: linux.BTN_DPAD_RIGHT,
				ButtonDown:  linux.BTN_DPAD_DOWN,
				ButtonLeft:  linux.BTN_DPAD_LEFT,

				ButtonSelect:  linux.BTN_SELECT, // Minus
				ButtonFiller1: linux.BTN_Z,      // Capture

				ButtonL1: linux.BTN_TL,
				ButtonL2: linux.BTN_TL2,

				ButtonR1: linux.BTN_TR,  // SL
				ButtonR2: linux.BTN_TR2, // SR

				ButtonL3: linux.BTN_THUMBL,
			},
		).
		WithLeftStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32767, Value: 0, Max: 32767, Flat: 500, Fuzz: 250},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32767, Value: 0, Max: 32767, Flat: 500, Fuzz: 250},
			},
		).
		Create()
}
