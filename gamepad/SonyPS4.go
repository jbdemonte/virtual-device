package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewSonyPS4() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_SONY).
				WithProduct(sdl.USB_PRODUCT_SONY_DS4_SLIM).
				WithVersion(0x8111).
				WithName("Sony Interactive Entertainment Wireless Controller"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: linux.BTN_SOUTH,
				ButtonEast:  linux.BTN_EAST,
				ButtonNorth: linux.BTN_NORTH,
				ButtonWest:  linux.BTN_WEST,

				ButtonSelect: linux.BTN_SELECT,
				ButtonStart:  linux.BTN_START,
				ButtonMode:   linux.BTN_MODE, // Button Playstation

				ButtonUp:    HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
				ButtonDown:  HatEvent{Axis: linux.ABS_HAT0Y, Value: 1},
				ButtonLeft:  HatEvent{Axis: linux.ABS_HAT0X, Value: -1},
				ButtonRight: HatEvent{Axis: linux.ABS_HAT0X, Value: 1},

				ButtonL1: linux.BTN_TL,
				ButtonR1: linux.BTN_TR,

				ButtonL2: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 0, Value: 0, Max: 255},
				ButtonR2: virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: 0, Value: 0, Max: 255},

				ButtonL3: linux.BTN_THUMBL,
				ButtonR3: linux.BTN_THUMBR,
			},
		).
		WithLeftStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: 0, Value: 0, Max: 255},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: 0, Value: 0, Max: 255},
			},
		).
		WithRightStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: 0, Value: 0, Max: 255},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: 0, Value: 0, Max: 255},
			},
		).
		Create()
}
