package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewXBox360() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_MICROSOFT).
				WithProduct(sdl.USB_PRODUCT_XBOX360_XUSB_CONTROLLER).
				WithVersion(0x107).
				WithName("Xbox 360 Wireless Receiver (XBOX) **"),
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

				ButtonUp:    []InputEvent{linux.BTN_TRIGGER_HAPPY3, HatEvent{Axis: linux.ABS_HAT0Y, Value: -1}},
				ButtonDown:  []InputEvent{linux.BTN_TRIGGER_HAPPY4, HatEvent{Axis: linux.ABS_HAT0Y, Value: 1}},
				ButtonLeft:  []InputEvent{linux.BTN_TRIGGER_HAPPY1, HatEvent{Axis: linux.ABS_HAT0X, Value: -1}},
				ButtonRight: []InputEvent{linux.BTN_TRIGGER_HAPPY2, HatEvent{Axis: linux.ABS_HAT0X, Value: 1}},

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
