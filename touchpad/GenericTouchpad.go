package touchpad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func NewGenericTouchpad() VirtualTouchpad {
	return NewVirtualTouchpadFactory().
		WithDevice(
			virtual_device.NewVirtualDevice().
				WithBusType(linux.BUS_USB).
				WithVendor(0x02).
				WithProduct(0x07).
				WithVersion(0x01).
				WithName("SynPS/2 Synaptics TouchPad"),
		).
		WithAxes([]virtual_device.AbsAxis{
			{Axis: linux.ABS_X, Min: 1472, Value: 1472, Max: 5472, Resolution: 40},
			{Axis: linux.ABS_Y, Min: 1408, Value: 1408, Max: 4448, Resolution: 40},
			{Axis: linux.ABS_PRESSURE, Min: 0, Value: 0, Max: 255},
			{Axis: linux.ABS_MT_SLOT, Min: 0, Value: 0, Max: 4},
			{Axis: linux.ABS_MT_POSITION_X, Min: 1472, Value: 0, Max: 5472, Resolution: 40},
			{Axis: linux.ABS_MT_POSITION_Y, Min: 1408, Value: 1408, Max: 4448, Resolution: 40},
			{Axis: linux.ABS_MT_TRACKING_ID, Min: 0, Value: 0, Max: 65535},
		}).
		WithButtons([]linux.Button{
			linux.BTN_LEFT,
			linux.BTN_RIGHT,
			linux.BTN_TOOL_FINGER,
			linux.BTN_TOUCH,
			linux.BTN_TOOL_DOUBLETAP,
			linux.BTN_TOOL_TRIPLETAP,
		}).
		WithProperties([]linux.InputProp{
			linux.INPUT_PROP_POINTER, linux.INPUT_PROP_BUTTONPAD,
		}).
		Create()

}
