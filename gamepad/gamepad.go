package gamepad

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func createVirtualGamepad(device virtual_device.VirtualDevice, mapping Mapping) VirtualGamepad {
	vg := &virtualGamepad{
		device:  device,
		mapping: mapping,
	}

	vg.setEvents()
	return vg
}

type VirtualGamepad interface {
	Register() error
	Unregister() error
	Press(button Button)
	Release(button Button)
	MoveLeftStick(x, y float32)
}

type virtualGamepad struct {
	device  virtual_device.VirtualDevice
	mapping Mapping
	config  Config
}

func (v *virtualGamepad) Register() error {
	return v.device.Register()
}

func (v *virtualGamepad) Unregister() error {
	return v.device.Unregister()
}

func (v *virtualGamepad) setEvents() {
	buttons := make([]linux.Button, 0)
	keys := make([]linux.Key, 0)
	absoluteAxes := make([]linux.AbsoluteAxis, 0)
	scanCodes := make([]uint32, 0)

	for _, events := range v.mapping.Digital {
		for _, event := range events {
			switch e := event.(type) {
			case linux.Button:
				buttons = append(buttons, e)
			case linux.Key:
				keys = append(keys, e)
			case MSCScanCode:
				fmt.Printf("scanCodes %d\n", e)
				scanCodes = append(scanCodes, uint32(e))
			case HatEvent:
				absoluteAxes = append(absoluteAxes, e.axe)
			case linux.AbsoluteAxis:
				absoluteAxes = append(absoluteAxes, e)
			default:
				fmt.Println("Unknown event type")
			}
		}
	}

	if v.mapping.Analog != nil {
		absoluteAxes = append(
			absoluteAxes,
			v.mapping.Analog.Left.X,
			v.mapping.Analog.Left.Y,
			v.mapping.Analog.Right.X,
			v.mapping.Analog.Right.Y,
		)
	}

	v.device.SetEventButtons(buttons)
	v.device.SetEventKeys(keys)
	v.device.SetEventScanCode(scanCodes)

	// todo absoluteAxes may contains duplicate due to Hat, need to check if it's a problem
	v.device.SetEventAbsoluteAxes(absoluteAxes)
}

func (v *virtualGamepad) Press(button Button) {
	events, exist := v.mapping.Digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			v.device.KeyDown(uint16(e))
		case linux.Key:
			v.device.KeyDown(uint16(e))
		case MSCScanCode:
			v.device.SendScanCode(int32(e))
		case HatEvent:
			v.device.SendAbsoluteEvent(uint16(e.axe), int32(e.value))
		default:
			fmt.Println("Unknown event type")
		}
	}
	v.device.SendSync()
}

func (v *virtualGamepad) Release(button Button) {
	events, exist := v.mapping.Digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			v.device.KeyUp(uint16(e))
		case linux.Key:
			v.device.KeyUp(uint16(e))
		case MSCScanCode:
			v.device.SendScanCode(int32(e))
		case HatEvent:
			v.device.SendAbsoluteEvent(uint16(e.axe), 0)
		default:
			fmt.Println("Unknown event type")
		}
	}
	v.device.SendSync()
}

// MoveLeftStick x & y are normalized values
func (v *virtualGamepad) MoveLeftStick(x, y float32) {
	if v.mapping.Analog == nil {
		return
	}
	// todo - replace 32767 by the config one if defined
	v.device.SendAbsoluteEvent(uint16(v.mapping.Analog.Left.X), int32(x*32767))
	v.device.SendAbsoluteEvent(uint16(v.mapping.Analog.Left.Y), int32(y*32767))
	v.device.SendSync()
}
