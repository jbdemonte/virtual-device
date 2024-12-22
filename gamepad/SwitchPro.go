package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewSwitchPro() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_NINTENDO).
				WithProduct(sdl.USB_PRODUCT_NINTENDO_SWITCH_PRO).
				WithVersion(0x8111).
				WithName("Nintendo Co., Ltd. Pro Controller"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: linux.BTN_SOUTH,
				ButtonEast:  linux.BTN_EAST,
				ButtonNorth: linux.BTN_NORTH,
				ButtonWest:  linux.BTN_WEST,

				ButtonSelect:  linux.BTN_SELECT, // Button -
				ButtonStart:   linux.BTN_START,  // Button +
				ButtonMode:    linux.BTN_MODE,   // Button Home
				ButtonFiller1: linux.BTN_Z,      // Button Square, under -

				ButtonUp:    HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
				ButtonDown:  HatEvent{Axis: linux.ABS_HAT0Y, Value: 1},
				ButtonLeft:  HatEvent{Axis: linux.ABS_HAT0X, Value: -1},
				ButtonRight: HatEvent{Axis: linux.ABS_HAT0X, Value: 1},

				ButtonL1: linux.BTN_TL,
				ButtonR1: linux.BTN_TR,

				ButtonL2: linux.BTN_TL2,
				ButtonR2: linux.BTN_TR2,

				ButtonL3: linux.BTN_THUMBL,
				ButtonR3: linux.BTN_THUMBR,
			},
		).
		WithLeftStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32767, Value: 0, Max: 32767, Fuzz: 250, Flat: 500},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32767, Value: 0, Max: 32767, Fuzz: 250, Flat: 500},
			},
		).
		WithRightStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: -32767, Value: 0, Max: 32767, Fuzz: 250, Flat: 500},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: -32767, Value: 0, Max: 32767, Fuzz: 250, Flat: 500},
			},
		).
		Create()
}
