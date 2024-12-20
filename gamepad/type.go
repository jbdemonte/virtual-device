package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

type MappingDigital map[Button]InputEvent // single or array

type MappingStick struct {
	X virtual_device.AbsAxis
	Y virtual_device.AbsAxis
}

type InputEvent interface{}

type MSCScanCode uint32

type HatEvent struct {
	Axis  linux.AbsoluteAxis
	Value int32
}
