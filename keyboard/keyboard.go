package keyboard

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

type VirtualKeyboard interface {
	Register() error
	Type(content string)
	Unregister() error

	PressKey(key linux.Key)
	ReleaseKey(key linux.Key)
	SetLed(led linux.Led, state bool)
}

type VirtualKeyboardFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualKeyboardFactory
	WithScanCode() VirtualKeyboardFactory
	WithKeys(keys []linux.Key) VirtualKeyboardFactory
	WithLEDs(leds []linux.Led) VirtualKeyboardFactory
	WithRepeat(delay, period int32) VirtualKeyboardFactory
	WithKeyMap(keymap KeyMap) VirtualKeyboardFactory
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
	keymap   KeyMap
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

func (f *virtualKeyboardFactory) WithKeyMap(keymap KeyMap) VirtualKeyboardFactory {
	f.keymap = keymap
	return f
}

func (f *virtualKeyboardFactory) Create() VirtualKeyboard {
	vk := &virtualKeyboard{
		device: f.device,
		keymap: f.keymap,
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
	keymap KeyMap
}

func (vk *virtualKeyboard) Register() error {
	return vk.device.Register()
}

func (vk *virtualKeyboard) Unregister() error {
	return vk.device.Unregister()
}

func (vk *virtualKeyboard) PressKey(key linux.Key) {
	vk.device.PressKey(key)
}

func (vk *virtualKeyboard) ReleaseKey(key linux.Key) {
	vk.device.ReleaseKey(key)
}

func (vk *virtualKeyboard) SetLed(led linux.Led, state bool) {
	vk.device.SetLed(led, state)
}

func (vk *virtualKeyboard) Type(content string) {
	if vk.keymap == nil {
		vk.keymap = getKeymap()
	}
	for _, char := range content {
		if mapping, ok := vk.keymap[char]; ok {
			if mapping.shiftRequired {
				vk.device.PressKey(linux.KEY_LEFTSHIFT)
			}
			if mapping.altGrRequired {
				vk.device.PressKey(linux.KEY_RIGHTALT)
			}

			vk.device.PressKey(mapping.keyCode)

			vk.device.SyncReport()

			time.Sleep(20 * time.Millisecond)

			vk.device.ReleaseKey(mapping.keyCode)

			if mapping.shiftRequired {
				vk.device.ReleaseKey(linux.KEY_LEFTSHIFT)
			}
			if mapping.altGrRequired {
				vk.device.ReleaseKey(linux.KEY_RIGHTALT)
			}

			vk.device.SyncReport()

			time.Sleep(30 * time.Millisecond)
		} else {
			fmt.Printf("Warning: Character '%c' is not mapped\n", char)
		}
	}
}
