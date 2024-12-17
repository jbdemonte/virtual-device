package keyboard

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

type VirtualKeyboard interface {
	Register() error
	Unregister() error
	KeyDown(key linux.Key)
	KeyPress(key linux.Key)
	KeyUp(key linux.Key)
	SwitchLed(led linux.Led, state bool)
}

type virtualKeyboard struct {
	device virtual_device.VirtualDevice
	config Config
}

func createVirtualKeyboard(device virtual_device.VirtualDevice, config Config) VirtualKeyboard {
	vk := &virtualKeyboard{device, config}
	vk.init()
	return vk
}

func (vk *virtualKeyboard) init() {
	if vk.config.scanCode {
		vk.device.ActivateScanCode()
	}
	if vk.config.repeat != nil {
		vk.device.SetRepeat(vk.config.repeat.delay, vk.config.repeat.period)
	}
	vk.device.SetEventKeys(vk.config.keys)
	vk.device.SetLeds(vk.config.leds)
}

func (vk *virtualKeyboard) Register() error {
	return vk.device.Register()
}

func (vk *virtualKeyboard) Unregister() error {
	return vk.device.Unregister()
}

func (vk *virtualKeyboard) KeyDown(key linux.Key) {
	vk.device.KeyDown(uint16(key))
}

func (vk *virtualKeyboard) KeyPress(key linux.Key) {
	vk.device.KeyPress(uint16(key))
}

func (vk *virtualKeyboard) KeyUp(key linux.Key) {
	vk.device.KeyUp(uint16(key))
}

func (vk *virtualKeyboard) SwitchLed(led linux.Led, state bool) {
	vk.device.SwitchLed(led, state)
}
