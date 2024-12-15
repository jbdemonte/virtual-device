package gamepad

import "github.com/jbdemonte/virtual-device/linux"

type MappingButtons map[Button]linux.Button

type MappingAxes map[Button]linux.AbsoluteAxis

type MappingHat map[Button]HatDirection

type MappingDigital map[Button][]InputEvent

type MappingStick struct {
	X linux.AbsoluteAxis
	Y linux.AbsoluteAxis
}

type MappingAnalog struct {
	Left  MappingStick
	Right MappingStick
}

type Mapping struct {
	Digital MappingDigital
	Analog  *MappingAnalog
}

type AxisConfig struct {
	Min int // Minimum value for the axis
	Max int // Maximum value for the axis
}

type Config struct {
	Axes map[linux.AbsoluteAxis]AxisConfig
}

type InputEvent interface{}

type MSCScanCode uint32

type HatEvent struct {
	axe   linux.AbsoluteAxis
	value int
}
