package virtual_device

import (
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/utils"
	"os"
)

type Events struct {
	keys         []linux.Key
	buttons      []linux.Button
	absoluteAxes []linux.AbsoluteAxis
}

type virtualDevice struct {
	fd           *os.File
	path         string
	mode         os.FileMode
	queueLen     int
	name         string
	id           linux.InputID
	events       Events
	isRegistered *utils.AtomicBool
	queue        chan *linux.InputEvent
}
