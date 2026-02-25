package testutil

import (
	"os"
	"sync"

	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

type Event struct {
	EvType uint16
	Code   uint16
	Value  int32
}

type MockDevice struct {
	mu     sync.Mutex
	events []Event

	Path         string
	Mode         os.FileMode
	QueueLen     int
	BusType      linux.BusType
	Vendor       uint16
	Product      uint16
	Version      uint16
	Name         string
	Keys         []linux.Key
	Buttons      []linux.Button
	AbsAxes      []virtual_device.AbsAxis
	RelAxes      []linux.RelativeAxis
	RepeatDelay  int32
	RepeatPeriod int32
	LEDs         []linux.Led
	Properties   []linux.InputProp
	MiscEvents   []linux.MiscEvent
}

func NewMockDevice() *MockDevice {
	return &MockDevice{}
}

func (m *MockDevice) Events() []Event {
	m.mu.Lock()
	defer m.mu.Unlock()
	cp := make([]Event, len(m.events))
	copy(cp, m.events)
	return cp
}

func (m *MockDevice) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = nil
}

func (m *MockDevice) record(evType, code uint16, value int32) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = append(m.events, Event{EvType: evType, Code: code, Value: value})
}

func (m *MockDevice) WithPath(path string) virtual_device.VirtualDevice {
	m.Path = path
	return m
}

func (m *MockDevice) WithMode(mode os.FileMode) virtual_device.VirtualDevice {
	m.Mode = mode
	return m
}

func (m *MockDevice) WithQueueLen(queueLen int) virtual_device.VirtualDevice {
	m.QueueLen = queueLen
	return m
}

func (m *MockDevice) WithBusType(busType linux.BusType) virtual_device.VirtualDevice {
	m.BusType = busType
	return m
}

func (m *MockDevice) WithVendor(vendor uint16) virtual_device.VirtualDevice {
	m.Vendor = vendor
	return m
}

func (m *MockDevice) WithProduct(product uint16) virtual_device.VirtualDevice {
	m.Product = product
	return m
}

func (m *MockDevice) WithVersion(version uint16) virtual_device.VirtualDevice {
	m.Version = version
	return m
}

func (m *MockDevice) WithName(name string) virtual_device.VirtualDevice {
	m.Name = name
	return m
}

func (m *MockDevice) WithKeys(keys []linux.Key) virtual_device.VirtualDevice {
	m.Keys = keys
	return m
}

func (m *MockDevice) WithButtons(buttons []linux.Button) virtual_device.VirtualDevice {
	m.Buttons = buttons
	return m
}

func (m *MockDevice) WithAbsAxes(absoluteAxes []virtual_device.AbsAxis) virtual_device.VirtualDevice {
	m.AbsAxes = absoluteAxes
	return m
}

func (m *MockDevice) WithRelAxes(relativeAxes []linux.RelativeAxis) virtual_device.VirtualDevice {
	m.RelAxes = relativeAxes
	return m
}

func (m *MockDevice) WithRepeat(delay, period int32) virtual_device.VirtualDevice {
	m.RepeatDelay = delay
	m.RepeatPeriod = period
	return m
}

func (m *MockDevice) WithLEDs(leds []linux.Led) virtual_device.VirtualDevice {
	m.LEDs = leds
	return m
}

func (m *MockDevice) WithProperties(properties []linux.InputProp) virtual_device.VirtualDevice {
	m.Properties = properties
	return m
}

func (m *MockDevice) WithMiscEvents(events []linux.MiscEvent) virtual_device.VirtualDevice {
	m.MiscEvents = events
	return m
}

func (m *MockDevice) Register() error   { return nil }
func (m *MockDevice) Unregister() error { return nil }

func (m *MockDevice) Send(evType, code uint16, value int32) {
	m.record(evType, code, value)
}

func (m *MockDevice) Sync(evType linux.SyncEvent) {
	m.Send(uint16(linux.EV_SYN), uint16(evType), 0)
}

func (m *MockDevice) SyncReport() {
	m.Sync(linux.SYN_REPORT)
}

func (m *MockDevice) PressKey(key linux.Key) {
	m.Send(uint16(linux.EV_KEY), uint16(key), 1)
}

func (m *MockDevice) ReleaseKey(key linux.Key) {
	m.Send(uint16(linux.EV_KEY), uint16(key), 0)
}

func (m *MockDevice) PressButton(button linux.Button) {
	m.Send(uint16(linux.EV_KEY), uint16(button), 1)
}

func (m *MockDevice) ReleaseButton(button linux.Button) {
	m.Send(uint16(linux.EV_KEY), uint16(button), 0)
}

func (m *MockDevice) SendAbsoluteEvent(axis linux.AbsoluteAxis, value int32) {
	m.Send(uint16(linux.EV_ABS), uint16(axis), value)
}

func (m *MockDevice) SendRelativeEvent(axis linux.RelativeAxis, value int32) {
	m.Send(uint16(linux.EV_REL), uint16(axis), value)
}

func (m *MockDevice) SendMiscEvent(event linux.MiscEvent, value int32) {
	m.Send(uint16(linux.EV_MSC), uint16(event), value)
}

func (m *MockDevice) SetLed(led linux.Led, state bool) {
	value := int32(0)
	if state {
		value = 1
	}
	m.Send(uint16(linux.EV_LED), uint16(led), value)
}

func (m *MockDevice) EventPath() string {
	return "/dev/input/event99"
}
