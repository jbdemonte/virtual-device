package virtual_device

import (
	"os"
	"virtual-input/linux"
)

type ID struct {
	busType linux.BusType
	vendor  uint16
	product uint16
	version uint16
}

type Events struct {
	keys []uint16
}

type virtualDevice struct {
	fd     *os.File
	path   string
	mode   os.FileMode
	name   string
	id     ID
	events Events
}
