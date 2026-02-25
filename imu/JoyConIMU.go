package imu

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

// NewJoyConIMU creates a virtual IMU device emulating a Nintendo Joy-Con accelerometer and gyroscope.
func NewJoyConIMU(isLeft bool) virtual_device.VirtualDevice {
	name := "Joy-Con (R) (IMU)"
	product := sdl.USB_PRODUCT_NINTENDO_SWITCH_JOYCON_RIGHT
	if isLeft {
		name = "Joy-Con (L) (IMU)"
		product = sdl.USB_PRODUCT_NINTENDO_SWITCH_JOYCON_LEFT
	}
	return virtual_device.
		NewVirtualDevice().
		WithBusType(linux.BUS_BLUETOOTH).
		WithVendor(sdl.USB_VENDOR_NINTENDO).
		WithProduct(product).
		WithVersion(0x8001).
		WithName(name).
		WithAbsAxes([]virtual_device.AbsAxis{
			// accelerometer that measures linear acceleration on three axes: X, Y, and Z
			virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32767, Value: 0, Max: 32767, Fuzz: 10, Resolution: 4096},
			virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32767, Value: 0, Max: 32767, Fuzz: 10, Resolution: 4096},
			virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: -32767, Value: 0, Max: 32767, Fuzz: 10, Resolution: 4096},

			// gyroscope that measures rotational velocity on three axes: RX, RY, and RZ
			virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: -32767000, Value: 0, Max: 32767000, Fuzz: 10, Resolution: 14247},
			virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: -32767000, Value: 0, Max: 32767000, Fuzz: 10, Resolution: 14247},
			virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: -32767000, Value: 0, Max: 32767000, Fuzz: 10, Resolution: 14247},
		}).
		WithMiscEvents([]linux.MiscEvent{linux.MSC_TIMESTAMP}).
		WithProperties([]linux.InputProp{linux.INPUT_PROP_ACCELEROMETER})
}
