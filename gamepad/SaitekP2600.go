package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewSaitekP2600() VirtualGamepad {
	return NewVirtualGamepadFactory().
		WithDevice(
			virtual_device.
				NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(sdl.USB_VENDOR_SAITEK).
				WithProduct(0xff0d).
				WithVersion(0x110).
				WithName("Saitek PLC Saitek P2600 Rumble Force Pad"),
		).
		WithDigital(
			MappingDigital{
				ButtonSouth: []InputEvent{MSCScanCode(90002), linux.BTN_THUMB},
				ButtonEast:  []InputEvent{MSCScanCode(90003), linux.BTN_THUMB2},
				ButtonNorth: []InputEvent{MSCScanCode(90004), linux.BTN_TOP},
				ButtonWest:  []InputEvent{MSCScanCode(90001), linux.BTN_TRIGGER},

				ButtonSelect: []InputEvent{MSCScanCode(0x9000b), linux.BTN_BASE5}, // red button (FPS) above the DPAD
				ButtonStart:  []InputEvent{MSCScanCode(0x9000c), linux.BTN_BASE6}, // Btn "Analog"

				ButtonUp:    HatEvent{Axis: linux.ABS_HAT0Y, Value: -1},
				ButtonDown:  HatEvent{Axis: linux.ABS_HAT0Y, Value: 1},
				ButtonLeft:  HatEvent{Axis: linux.ABS_HAT0X, Value: -1},
				ButtonRight: HatEvent{Axis: linux.ABS_HAT0X, Value: 1},

				ButtonL1: []InputEvent{MSCScanCode(90005), linux.BTN_TOP2},
				ButtonR1: []InputEvent{MSCScanCode(90006), linux.BTN_PINKIE},

				ButtonL2: []InputEvent{MSCScanCode(90007), linux.BTN_BASE},
				ButtonR2: []InputEvent{MSCScanCode(90008), linux.BTN_BASE2},

				ButtonL3: []InputEvent{MSCScanCode(0x9000a), linux.BTN_BASE4}, // most right gray button
				ButtonR3: []InputEvent{MSCScanCode(90009), linux.BTN_BASE3},   // most right black button, under the ButtonL3
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
				X: virtual_device.AbsAxis{Axis: linux.ABS_RUDDER, Min: 0, Value: 0, Max: 255, Flat: 15},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_THROTTLE, Min: 0, Value: 0, Max: 255, Flat: 15},
			},
		).
		Create()
}
