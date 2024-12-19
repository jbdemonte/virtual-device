package virtual_device

import (
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/utils"
	"os"
)

type AbsAxis struct {
	Axis       linux.AbsoluteAxis
	Value      int32
	Min        int32
	Max        int32
	Flat       int32
	Fuzz       int32
	Resolution int32
}

type Repeat struct {
	delay  int32
	period int32
}

func (a AbsAxis) Denormalize(value float32) int32 {
	if value < -1 {
		value = -1
	} else if value > 1 {
		value = 1
	}
	result := float32(a.Min) + (value+1)*float32(a.Max-a.Min)/2
	return int32(result)
}

type Events struct {
	keys         []linux.Key
	buttons      []linux.Button
	absoluteAxes []AbsAxis
	relativeAxes []linux.RelativeAxis
	scanCode     bool
	repeat       *Repeat
	leds         []linux.Led
	properties   []linux.InputProp
}

type virtualDevice struct {
	fd           *os.File
	path         string
	eventPath    string
	mode         os.FileMode
	queueLen     int
	name         string
	id           linux.InputID
	events       Events
	isRegistered *utils.AtomicBool
	queue        chan *linux.InputEvent
}
