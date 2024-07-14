package virtual_device

import (
	"context"
	"github.com/jbdemonte/virtual-device/linux"
	"os"
	"sync"
)

type Events struct {
	keys    []linux.Key
	buttons []linux.Button
}

type virtualDevice struct {
	fd       *os.File
	path     string
	mode     os.FileMode
	queueLen int
	name     string
	id       linux.InputID
	events   Events

	out    chan *linux.InputEvent
	cancel context.CancelFunc
	mu     sync.Mutex
}
