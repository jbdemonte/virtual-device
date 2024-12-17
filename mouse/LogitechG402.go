package mouse

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewLogitechG402() VirtualMouse {
	vd := virtual_device.NewVirtualDevice()

	vd.
		SetBusType(linux.BUS_USB).
		SetProduct(0xc07e).
		SetVendor(sdl.USB_VENDOR_LOGITECH).
		SetVersion(0x111).
		SetName("Logitech Gaming Mouse G402")

	vd.SetEventButtons([]linux.Button{
		linux.BTN_LEFT,
		linux.BTN_RIGHT,
		linux.BTN_MIDDLE,
		linux.BTN_SIDE,
		linux.BTN_EXTRA,
		linux.BTN_FORWARD,
		linux.BTN_BACK,
		linux.BTN_TASK,

		linux.Button(280), // ?
		linux.Button(281), // ?
		linux.Button(282), // ?
		linux.Button(283), // ?
		linux.Button(284), // ?
		linux.Button(285), // ?
		linux.Button(286), // ?
		linux.Button(287), // ?
	})

	vd.ActivateScanCode()

	vd.SetEventRelativeAxes([]linux.RelativeAxis{
		linux.REL_X,
		linux.REL_Y,
		linux.REL_HWHEEL,
		linux.REL_WHEEL,
		linux.REL_WHEEL_HI_RES,
		linux.REL_HWHEEL_HI_RES,
	})

	config := Config{
		highResStep:  1,
		highResHStep: 120,
	}

	return createVirtualMouse(vd, config)
}
