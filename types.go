package virtual_input

import (
	"os"
	"virtual-input/linux"
)

type ID struct {
	BusType linux.BusType
	Vendor  uint16
	Product uint16
	Version uint16
}

type UserInput struct {
	fd *os.File

	syspath string

	id         ID
	name       string
	path       string
	effectsMax uint32

	events map[linux.EventType]bool
}
