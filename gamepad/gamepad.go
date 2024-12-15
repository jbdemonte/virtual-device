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

	vg.setEventButtons()
	vg.setEventAbsoluteAxes()

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

func (v *virtualGamepad) setEventButtons() {
	buttons := make([]linux.Button, 0)

	if v.mapping.Digital.Buttons != nil {
		for _, button := range v.mapping.Digital.Buttons {
			buttons = append(buttons, button)
		}
	}
	if len(buttons) > 0 {
		v.device.SetEventButtons(buttons)
	}
}

func (v *virtualGamepad) setEventAbsoluteAxes() {
	absoluteAxes := make([]linux.AbsoluteAxis, 0)

	if v.mapping.Digital.Hat != nil {
		for _, hat := range v.mapping.Digital.Hat {
			if hat == HatUp || hat == HatDown {
				absoluteAxes = append(absoluteAxes, linux.ABS_HAT0Y)
			}
			if hat == HatLeft || hat == HatRight {
				absoluteAxes = append(absoluteAxes, linux.ABS_HAT0X)
			}
		}
	}

	if v.mapping.Digital.Axes != nil {
		for _, axis := range v.mapping.Digital.Axes {
			absoluteAxes = append(absoluteAxes, axis)
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

	if len(absoluteAxes) > 0 {
		// todo absoluteAxes may contains duplicate due to Hat, need to check if it's a problem
		v.device.SetEventAbsoluteAxes(absoluteAxes)
	}
}

// todo - handle Axes / Hat when used as a button

func (v *virtualGamepad) Press(button Button) {
	b, exist := v.mapping.Digital.Buttons[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	v.device.KeyDown(uint16(b))
}

func (v *virtualGamepad) Release(button Button) {
	b, exist := v.mapping.Digital.Buttons[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	v.device.KeyUp(uint16(b))
}

// x & y are normalized values
func (v *virtualGamepad) MoveLeftStick(x, y float32) {
	if v.mapping.Analog == nil {
		return
	}
	// todo - replace 32767 by the config one if defined
	v.device.SendStickAxisEvent(uint16(v.mapping.Analog.Left.X), int32(x*32767))
	v.device.SendStickAxisEvent(uint16(v.mapping.Analog.Left.Y), int32(y*32767))
}
