package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewStadia() VirtualGamepad {
	mapping := Mapping{
		Digital: MappingDigital{
			ButtonSouth: []InputEvent{linux.BTN_SOUTH},
			ButtonEast:  []InputEvent{linux.BTN_EAST},
			ButtonNorth: []InputEvent{MSCScanCode(90005), linux.BTN_WEST},
			ButtonWest:  []InputEvent{linux.BTN_NORTH},

			ButtonL1: []InputEvent{linux.BTN_TL},
			ButtonR1: []InputEvent{linux.BTN_TR},

			ButtonSelect: []InputEvent{linux.BTN_SELECT},
			ButtonStart:  []InputEvent{linux.BTN_START},
			ButtonMode:   []InputEvent{linux.BTN_MODE}, // button stadia

			ButtonFiller1: []InputEvent{linux.BTN_TRIGGER_HAPPY1}, // Assistant
			ButtonFiller2: []InputEvent{linux.BTN_TRIGGER_HAPPY2}, // Capture

			ButtonL3: []InputEvent{linux.BTN_THUMBL},
			ButtonR3: []InputEvent{linux.BTN_THUMBR},

			ButtonUp:    []InputEvent{linux.BTN_DPAD_UP, HatEvent{axe: linux.ABS_HAT0Y, value: -1}},
			ButtonDown:  []InputEvent{linux.BTN_DPAD_DOWN, HatEvent{axe: linux.ABS_HAT0Y, value: 1}},
			ButtonLeft:  []InputEvent{linux.BTN_DPAD_LEFT, HatEvent{axe: linux.ABS_HAT0X, value: -1}},
			ButtonRight: []InputEvent{linux.BTN_DPAD_RIGHT, HatEvent{axe: linux.ABS_HAT0X, value: 1}},
			ButtonL2:    []InputEvent{linux.BTN_TRIGGER_HAPPY4, linux.ABS_BRAKE},
			ButtonR2:    []InputEvent{linux.BTN_TRIGGER_HAPPY3, linux.ABS_GAS},
		},
		Analog: &MappingAnalog{
			Left: MappingStick{
				X: linux.ABS_X,
				Y: linux.ABS_Y,
			},
			Right: MappingStick{
				X: linux.ABS_Z,
				Y: linux.ABS_RZ,
			},
		},
	}

	vd := virtual_device.NewVirtualDevice()

	vd.
		SetBusType(linux.BUS_USB).
		SetProduct(sdl.USB_PRODUCT_GOOGLE_STADIA_CONTROLLER).
		SetVendor(sdl.USB_VENDOR_GOOGLE).
		SetVersion(0x111).
		SetName("Google LLC Stadia Controller rev. A")

	return createVirtualGamepad(vd, mapping)
}
