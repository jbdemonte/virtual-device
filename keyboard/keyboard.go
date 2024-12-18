package keyboard

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

type VirtualKeyboard interface {
	Register() error
	Unregister() error

	KeyPress(key linux.Key)
	KeyRelease(key linux.Key)
	Led(led linux.Led, state bool)
}

type VirtualKeyboardFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualKeyboardFactory
	WithScanCode() VirtualKeyboardFactory
	WithKeys(keys []linux.Key) VirtualKeyboardFactory
	WithLEDs(leds []linux.Led) VirtualKeyboardFactory
	WithRepeat(delay, period int32) VirtualKeyboardFactory
	Create() VirtualKeyboard
}

func NewVirtualKeyboardFactory() VirtualKeyboardFactory {
	return &virtualKeyboardFactory{}
}

type virtualKeyboardFactory struct {
	device   virtual_device.VirtualDevice
	scanCode bool
	keys     []linux.Key
	leds     []linux.Led
	repeat   *Repeat
}

func (f *virtualKeyboardFactory) WithDevice(device virtual_device.VirtualDevice) VirtualKeyboardFactory {
	f.device = device
	return f
}

func (f *virtualKeyboardFactory) WithScanCode() VirtualKeyboardFactory {
	f.scanCode = true
	return f
}

func (f *virtualKeyboardFactory) WithKeys(keys []linux.Key) VirtualKeyboardFactory {
	f.keys = keys
	return f
}

func (f *virtualKeyboardFactory) WithLEDs(leds []linux.Led) VirtualKeyboardFactory {
	f.leds = leds
	return f
}

func (f *virtualKeyboardFactory) WithRepeat(delay, period int32) VirtualKeyboardFactory {
	f.repeat = &Repeat{delay, period}
	return f
}

func (f *virtualKeyboardFactory) Create() VirtualKeyboard {
	vk := &virtualKeyboard{
		device: f.device,
	}
	if f.scanCode {
		vk.device.WithScanCode()
	}
	if f.repeat != nil {
		vk.device.WithRepeat(f.repeat.delay, f.repeat.period)
	}
	vk.device.WithKeys(f.keys)
	vk.device.WithLEDs(f.leds)
	return vk
}

type virtualKeyboard struct {
	device virtual_device.VirtualDevice
}

func (vk *virtualKeyboard) Register() error {
	return vk.device.Register()
}

func (vk *virtualKeyboard) Unregister() error {
	return vk.device.Unregister()
}

func (vk *virtualKeyboard) KeyPress(key linux.Key) {
	vk.device.KeyPress(uint16(key))
}

func (vk *virtualKeyboard) KeyRelease(key linux.Key) {
	vk.device.KeyRelease(uint16(key))
}

func (vk *virtualKeyboard) Led(led linux.Led, state bool) {
	vk.device.Led(led, state)
}
