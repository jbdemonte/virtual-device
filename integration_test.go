//go:build integration

package virtual_device

import (
	"bytes"
	"encoding/binary"
	"os"
	"testing"
	"time"
	"unsafe"

	"github.com/jbdemonte/virtual-device/linux"
)

func TestIntegration_DeviceLifecycle(t *testing.T) {
	vd := NewVirtualDevice().
		WithName("test-lifecycle").
		WithKeys([]linux.Key{linux.KEY_A})

	if err := vd.Register(); err != nil {
		t.Fatalf("Register: %v", err)
	}
	defer vd.Unregister()

	path := vd.EventPath()
	if path == "" {
		t.Fatal("EventPath is empty after Register")
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("event file %s does not exist: %v", path, err)
	}

	if err := vd.Unregister(); err != nil {
		t.Fatalf("Unregister: %v", err)
	}
}

func TestIntegration_SendAndReadKeyEvent(t *testing.T) {
	vd := NewVirtualDevice().
		WithName("test-key-event").
		WithKeys([]linux.Key{linux.KEY_A})

	if err := vd.Register(); err != nil {
		t.Fatalf("Register: %v", err)
	}
	defer vd.Unregister()

	path := vd.EventPath()

	// Open event file for reading
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open event file: %v", err)
	}
	defer f.Close()

	// Send key press + sync
	vd.PressKey(linux.KEY_A)
	vd.SyncReport()

	// Give events time to be written
	time.Sleep(50 * time.Millisecond)

	events := readEvents(t, f)
	if len(events) == 0 {
		t.Fatal("no events read from event file")
	}

	foundKeyPress := false
	for _, ev := range events {
		if ev.Type == uint16(linux.EV_KEY) && ev.Code == uint16(linux.KEY_A) && ev.Value == 1 {
			foundKeyPress = true
		}
	}
	if !foundKeyPress {
		t.Errorf("KEY_A press not found in events: %+v", events)
	}
}

func TestIntegration_SendBeforeRegister(t *testing.T) {
	vd := NewVirtualDevice().
		WithName("test-no-panic").
		WithKeys([]linux.Key{linux.KEY_A})

	// Should not panic
	vd.PressKey(linux.KEY_A)
	vd.SyncReport()
	vd.Send(uint16(linux.EV_KEY), uint16(linux.KEY_A), 0)
}

func TestIntegration_DoubleRegister(t *testing.T) {
	vd := NewVirtualDevice().
		WithName("test-double-register").
		WithKeys([]linux.Key{linux.KEY_A})

	if err := vd.Register(); err != nil {
		t.Fatalf("first Register: %v", err)
	}
	defer vd.Unregister()

	// Second register should be a no-op (returns nil)
	if err := vd.Register(); err != nil {
		t.Fatalf("second Register should succeed: %v", err)
	}
}

func readEvents(t *testing.T, f *os.File) []linux.InputEvent {
	t.Helper()

	// Set a read deadline so we don't block forever
	f.SetReadDeadline(time.Now().Add(time.Second))

	eventSize := int(unsafe.Sizeof(linux.InputEvent{}))
	buf := make([]byte, eventSize*64)

	n, _ := f.Read(buf)
	if n == 0 {
		return nil
	}

	reader := bytes.NewReader(buf[:n])
	var events []linux.InputEvent
	for reader.Len() >= eventSize {
		var ev linux.InputEvent
		if err := binary.Read(reader, binary.LittleEndian, &ev); err != nil {
			break
		}
		events = append(events, ev)
	}
	return events
}
