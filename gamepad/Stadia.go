package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewStadia() VirtualGamepad {

	vd := virtual_device.NewVirtualDevice()

	vd.
		SetBusType(linux.BUS_USB).
		SetProduct(sdl.USB_PRODUCT_GOOGLE_STADIA_CONTROLLER).
		SetVendor(sdl.USB_VENDOR_GOOGLE).
		SetVersion(0x111).
		SetName("Google LLC Stadia Controller rev. A")

	mapping := Mapping{
		Digital: MappingDigital{
			ButtonSouth: []InputEvent{MSCScanCode(90001), linux.BTN_SOUTH},
			ButtonEast:  []InputEvent{MSCScanCode(90002), linux.BTN_EAST},
			ButtonNorth: []InputEvent{MSCScanCode(90005), linux.BTN_WEST},
			ButtonWest:  []InputEvent{MSCScanCode(90004), linux.BTN_NORTH},

			ButtonSelect: []InputEvent{MSCScanCode(0x9000b), linux.BTN_SELECT}, // button Option
			ButtonStart:  []InputEvent{MSCScanCode(0x9000c), linux.BTN_START},  // button Menu
			ButtonMode:   []InputEvent{MSCScanCode(0x9000d), linux.BTN_MODE},   // button Stadia

			ButtonFiller1: []InputEvent{MSCScanCode(90011), linux.BTN_TRIGGER_HAPPY1}, // Button Google Assistant
			ButtonFiller2: []InputEvent{MSCScanCode(90012), linux.BTN_TRIGGER_HAPPY2}, // Button Capture

			ButtonUp:    []InputEvent{HatEvent{Axis: linux.ABS_HAT0Y, Value: -1}},
			ButtonDown:  []InputEvent{HatEvent{Axis: linux.ABS_HAT0Y, Value: 1}},
			ButtonLeft:  []InputEvent{HatEvent{Axis: linux.ABS_HAT0X, Value: -1}},
			ButtonRight: []InputEvent{HatEvent{Axis: linux.ABS_HAT0X, Value: 1}},

			ButtonL1: []InputEvent{MSCScanCode(90007), linux.BTN_TL},
			ButtonR1: []InputEvent{MSCScanCode(90008), linux.BTN_TR},

			ButtonL2: []InputEvent{MSCScanCode(90014), linux.BTN_TRIGGER_HAPPY4, virtual_device.AbsAxis{Axis: linux.ABS_BRAKE, Min: 0, Value: 0, Max: 255, Flat: 15}},
			ButtonR2: []InputEvent{MSCScanCode(90013), linux.BTN_TRIGGER_HAPPY3, virtual_device.AbsAxis{Axis: linux.ABS_GAS, Min: 0, Value: 0, Max: 255, Flat: 15}},

			ButtonL3: []InputEvent{MSCScanCode(0x9000e), linux.BTN_THUMBL},
			ButtonR3: []InputEvent{MSCScanCode(0x9000f), linux.BTN_THUMBR},
		},
		Analog: &MappingAnalog{
			Left: MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: 1, Value: 128, Max: 255, Flat: 15},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: 1, Value: 128, Max: 255, Flat: 15},
			},
			Right: MappingStick{
				X: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 1, Value: 128, Max: 255, Flat: 15},
				Y: virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: 1, Value: 128, Max: 255, Flat: 15},
			},
		},
	}

	return createVirtualGamepad(vd, mapping)
}
