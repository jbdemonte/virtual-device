package mouse

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func NewGenericMouse() VirtualMouse {
	return NewVirtualMouseFactory().
		WithDevice(
			virtual_device.NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(0xDEAD).
				WithProduct(0xBEEF).
				WithVersion(0x01).
				WithName("Generic Mouse").
				WithButtons([]linux.Button{
					linux.BTN_LEFT,
					linux.BTN_RIGHT,
					linux.BTN_MIDDLE,
				}).
				WithRelAxes([]linux.RelativeAxis{
					linux.REL_X,
					linux.REL_Y,
					linux.REL_WHEEL,
				}),
		).
		Create()
}
