package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewJoyConR() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_BLUETOOTH).
				WithVendor(sdl.USB_VENDOR_NINTENDO).
				WithProduct(sdl.USB_PRODUCT_NINTENDO_SWITCH_JOYCON_RIGHT).
				WithVersion(0x8001).
				WithName("Joy-Con (R)"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: linux.BTN_SOUTH,
				ButtonEast:  linux.BTN_EAST,
				ButtonNorth: linux.BTN_NORTH,
				ButtonWest:  linux.BTN_WEST,

				ButtonStart: linux.BTN_START, // Plus
				ButtonMode:  linux.BTN_MODE,  // Home

				ButtonR1: linux.BTN_TR,
				ButtonR2: linux.BTN_TR2,

				ButtonL1: linux.BTN_TL,  // SL
				ButtonL2: linux.BTN_TL2, // SR

				ButtonR3: linux.BTN_THUMBR,
			},
		).
		WithRightStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: -32767, Value: 0, Max: 32767, Flat: 500, Fuzz: 250},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: -32767, Value: 0, Max: 32767, Flat: 500, Fuzz: 250},
			},
		).
		Create()
}
