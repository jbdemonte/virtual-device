package virtual_device

import (
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/utils"
	"os"
)

// AbsAxis describes an absolute axis with its range and properties.
type AbsAxis struct {
	Axis             linux.AbsoluteAxis
	Value            int32
	Min              int32
	Max              int32
	Flat             int32
	Fuzz             int32
	Resolution       int32
	IsUnidirectional bool
}

// Repeat holds the key repeat delay and period in milliseconds.
type Repeat struct {
	delay  int32
	period int32
}

func (a AbsAxis) denormalizeUniDirectional(value float32) int32 {
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}
	return int32(float32(a.Min) + (value)*float32(a.Max-a.Min))
}

func (a AbsAxis) denormalizeBiDirectional(value float32) int32 {
	if value < -1 {
		value = -1
	}
	if value > 1 {
		value = 1
	}
	return int32(float32(a.Min) + (value+1)*float32(a.Max-a.Min)/2)
}

func (a AbsAxis) Denormalize(value float32) int32 {
	if a.IsUnidirectional {
		return a.denormalizeUniDirectional(value)
	}
	return a.denormalizeBiDirectional(value)
}

// Config holds the input capabilities for a virtual device.
type Config struct {
	keys         []linux.Key
	buttons      []linux.Button
	absoluteAxes []AbsAxis
	relativeAxes []linux.RelativeAxis
	repeat       *Repeat
	leds         []linux.Led
	properties   []linux.InputProp
	miscEvents   []linux.MiscEvent
}

type virtualDevice struct {
	fd           *os.File
	path         string
	eventPath    string
	mode         os.FileMode
	queueLen     int
	name         string
	id           linux.InputID
	config       Config
	isRegistered *utils.AtomicBool
	queue        chan *linux.InputEvent
	pullDone     chan struct{}
}
