package keyboard

import (
	"fmt"
	"time"

	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

// VirtualKeyboard is a high-level virtual keyboard input device.
type VirtualKeyboard interface {
	Register() error
	Type(content string)
	Unregister() error

	PressKey(key linux.Key)
	ReleaseKey(key linux.Key)
	TapKey(key linux.Key)
	SetLed(led linux.Led, state bool)
	SendMiscEvent(event linux.MiscEvent, value int32)
	SyncReport()

	Send(evType, code uint16, value int32)

	EventPath() string
}

// VirtualKeyboardFactory configures and creates VirtualKeyboard instances.
type VirtualKeyboardFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualKeyboardFactory
	WithTapDuration(duration time.Duration) VirtualKeyboardFactory
	WithKeys(keys []linux.Key) VirtualKeyboardFactory
	WithLEDs(leds []linux.Led) VirtualKeyboardFactory
	WithMiscEvents(events []linux.MiscEvent) VirtualKeyboardFactory
	WithRepeat(delay, period int32) VirtualKeyboardFactory
	WithKeyMap(keymap KeyMap) VirtualKeyboardFactory
	Create() VirtualKeyboard
}

// NewVirtualKeyboardFactory returns a new factory for building virtual keyboards.
func NewVirtualKeyboardFactory() VirtualKeyboardFactory {
	return &virtualKeyboardFactory{
		tapDuration: -1,
	}
}

type virtualKeyboardFactory struct {
	device      virtual_device.VirtualDevice
	keys        []linux.Key
	leds        []linux.Led
	miscEvents  []linux.MiscEvent
	repeat      *Repeat
	keymap      KeyMap
	tapDuration time.Duration
}

func (f *virtualKeyboardFactory) WithDevice(device virtual_device.VirtualDevice) VirtualKeyboardFactory {
	f.device = device
	return f
}

func (f *virtualKeyboardFactory) WithTapDuration(duration time.Duration) VirtualKeyboardFactory {
	f.tapDuration = duration
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

func (f *virtualKeyboardFactory) WithMiscEvents(events []linux.MiscEvent) VirtualKeyboardFactory {
	f.miscEvents = events
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
	tapDuration := f.tapDuration
	if tapDuration < 0 {
		tapDuration = 20 * time.Millisecond
	}

	vk := &virtualKeyboard{
		device:      f.device,
		keymap:      f.keymap,
		tapDuration: tapDuration,
	}
	if f.repeat != nil {
		vk.device.WithRepeat(f.repeat.delay, f.repeat.period)
	}
	vk.device.WithKeys(f.keys)
	vk.device.WithLEDs(f.leds)
	if len(f.miscEvents) > 0 {
		vk.device.WithMiscEvents(f.miscEvents)
	}
	return vk
}

type virtualKeyboard struct {
	device      virtual_device.VirtualDevice
	keymap      KeyMap
	tapDuration time.Duration
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

func (vk *virtualKeyboard) TapKey(key linux.Key) {
	vk.device.PressKey(key)
	vk.device.SyncReport()
	time.Sleep(vk.tapDuration)
	vk.device.ReleaseKey(key)
	vk.device.SyncReport()
}

func (vk *virtualKeyboard) SetLed(led linux.Led, state bool) {
	vk.device.SetLed(led, state)
}

func (vk *virtualKeyboard) SendMiscEvent(event linux.MiscEvent, value int32) {
	vk.device.SendMiscEvent(event, value)
}

func (vk *virtualKeyboard) SyncReport() {
	vk.device.SyncReport()
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

			time.Sleep(vk.tapDuration)

			vk.device.ReleaseKey(mapping.keyCode)

			if mapping.shiftRequired {
				vk.device.ReleaseKey(linux.KEY_LEFTSHIFT)
			}
			if mapping.altGrRequired {
				vk.device.ReleaseKey(linux.KEY_RIGHTALT)
			}

			vk.device.SyncReport()

			time.Sleep(vk.tapDuration)
		} else {
			fmt.Printf("Warning: Character '%c' is not mapped\n", char)
		}
	}
}

func (vk *virtualKeyboard) Send(evType, code uint16, value int32) {
	vk.device.Send(evType, code, value)
}

func (vk *virtualKeyboard) EventPath() string {
	return vk.device.EventPath()
}
