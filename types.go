package virtual_device

import (
	"context"
	"github.com/jbdemonte/virtual-device/linux"
	"os"
)

type Events struct {
	keys    []linux.Key
	buttons []linux.Button
}

type virtualDevice struct {
	fd     *os.File
	path   string
	mode   os.FileMode
	name   string
	id     linux.InputID
	events Events

	out    chan linux.InputEvent
	cancel context.CancelFunc
}
