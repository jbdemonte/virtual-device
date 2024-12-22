package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewSN30Pro() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_8BITDO).
				WithProduct(0x6001).
				WithVersion(0x111).
				WithName("8Bitdo SF30 Pro   8Bitdo SN30 Pro"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: []InputEvent{MSCScanCode(90002), linux.BTN_EAST},
				ButtonEast:  []InputEvent{MSCScanCode(90001), linux.BTN_SOUTH},
				ButtonNorth: []InputEvent{MSCScanCode(90004), linux.BTN_NORTH},
				ButtonWest:  []InputEvent{MSCScanCode(90005), linux.BTN_WEST},

				ButtonSelect: []InputEvent{MSCScanCode(0x9000b), linux.BTN_SELECT},
				ButtonStart:  []InputEvent{MSCScanCode(0x9000c), linux.BTN_START},
				ButtonMode:   []InputEvent{MSCScanCode(0x9000d), linux.BTN_MODE}, // Button under South button (B)

				ButtonFiller1: linux.BTN_C,         // ?
				ButtonFiller2: linux.BTN_Z,         // ?
				ButtonFiller3: linux.Button(0x13f), // ?

				ButtonUp:    HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
				ButtonDown:  HatEvent{Axis: linux.ABS_HAT0Y, Value: 1},
				ButtonLeft:  HatEvent{Axis: linux.ABS_HAT0X, Value: -1},
				ButtonRight: HatEvent{Axis: linux.ABS_HAT0X, Value: 1},

				ButtonL1: []InputEvent{MSCScanCode(90007), linux.BTN_TL},
				ButtonR1: []InputEvent{MSCScanCode(90008), linux.BTN_TR},

				ButtonL2: []InputEvent{MSCScanCode(90009), virtual_device.AbsAxis{Axis: linux.ABS_BRAKE, Min: 0, Value: 0, Max: 255, Flat: 15}, linux.BTN_TL2},
				ButtonR2: []InputEvent{MSCScanCode(0x9000a), virtual_device.AbsAxis{Axis: linux.ABS_GAS, Min: 0, Value: 0, Max: 255, Flat: 15}, linux.BTN_TR2},

				ButtonL3: []InputEvent{MSCScanCode(0x9000e), linux.BTN_THUMBL},
				ButtonR3: []InputEvent{MSCScanCode(0x9000f), linux.BTN_THUMBR},
			},
		).
		WithLeftStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: 0, Value: 0, Max: 255, Flat: 15},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: 0, Value: 0, Max: 255, Flat: 15},
			},
		).
		WithRightStick(
			MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 0, Value: 0, Max: 255, Flat: 15},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: 0, Value: 0, Max: 255, Flat: 15},
			},
		).
		Create()
}
