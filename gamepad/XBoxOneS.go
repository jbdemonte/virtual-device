package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewXBoxOneS() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_MICROSOFT).
				WithProduct(sdl.USB_PRODUCT_XBOX_ONE_S).
				WithVersion(0x408).
				WithName("Microsoft X-Box One S pad"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: linux.BTN_SOUTH,
				ButtonEast:  linux.BTN_EAST,
				ButtonNorth: linux.BTN_WEST,
				ButtonWest:  linux.BTN_NORTH,

				ButtonSelect: linux.BTN_SELECT,
				ButtonStart:  linux.BTN_START,
				ButtonMode:   linux.BTN_MODE, // button XBox

				ButtonUp:    HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
				ButtonDown:  HatEvent{Axis: linux.ABS_HAT0Y, Value: 1},
				ButtonLeft:  HatEvent{Axis: linux.ABS_HAT0X, Value: -1},
				ButtonRight: HatEvent{Axis: linux.ABS_HAT0X, Value: 1},

				ButtonL1: linux.BTN_TL,
				ButtonR1: linux.BTN_TR,

				ButtonL2: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 0, Value: 0, Max: 1023},
				ButtonR2: virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: 0, Value: 0, Max: 1023},

				ButtonL3: linux.BTN_THUMBL,
				ButtonR3: linux.BTN_THUMBR,
			},
		).
		WithLeftStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
			},
		).
		WithRightStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
			},
		).
		Create()
}
