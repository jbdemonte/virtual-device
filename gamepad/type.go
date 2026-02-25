package gamepad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

// MappingDigital maps gamepad buttons to their input events.
type MappingDigital map[Button]InputEvent // single or array

// MappingStick defines the X and Y axes for an analog stick.
type MappingStick struct {
	X virtual_device.AbsAxis
	Y virtual_device.AbsAxis
}

// InputEvent represents a gamepad input event (button, key, axis, or scan code).
type InputEvent interface{}

// MSCScanCode is a scan code value sent as an MSC_SCAN event.
type MSCScanCode uint32

// HatEvent represents a D-pad hat switch event with axis and direction.
type HatEvent struct {
	Axis  linux.AbsoluteAxis
	Value int32
}
