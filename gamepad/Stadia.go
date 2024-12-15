package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewStadia() VirtualGamepad {
	mapping := Mapping{
		Digital: MappingDigital{
			Buttons: MappingButtons{
				ButtonSouth: linux.BTN_SOUTH,
				ButtonEast:  linux.BTN_EAST,
				ButtonNorth: linux.BTN_WEST,
				ButtonWest:  linux.BTN_NORTH,

				ButtonL1: linux.BTN_TL,
				ButtonR1: linux.BTN_TR,

				ButtonL2: linux.BTN_TRIGGER_HAPPY4,
				ButtonR2: linux.BTN_TRIGGER_HAPPY3,

				ButtonSelect: linux.BTN_SELECT,
				ButtonStart:  linux.BTN_START,
				ButtonMode:   linux.BTN_MODE, // button stadia

				ButtonFiller1: linux.BTN_TRIGGER_HAPPY1, // Assistant
				ButtonFiller2: linux.BTN_TRIGGER_HAPPY2, // Capture

				ButtonL3: linux.BTN_THUMBL,
				ButtonR3: linux.BTN_THUMBR,

				ButtonUp:    linux.BTN_DPAD_UP,
				ButtonDown:  linux.BTN_DPAD_DOWN,
				ButtonLeft:  linux.BTN_DPAD_LEFT,
				ButtonRight: linux.BTN_DPAD_RIGHT,
			},
			Axes: MappingAxes{
				ButtonL2: linux.ABS_BRAKE,
				ButtonR2: linux.ABS_GAS,
			},
			Hat: MappingHat{
				ButtonUp:    HatUp,
				ButtonDown:  HatDown,
				ButtonLeft:  HatLeft,
				ButtonRight: HatRight,
			},
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
