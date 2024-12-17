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
	MoveLeftStickX(x float32)
	MoveLeftStickY(y float32)
	MoveRightStick(x, y float32)
	MoveRightStickX(x float32)
	MoveRightStickY(y float32)
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
	absoluteAxes := make([]virtual_device.AbsAxis, 0)
	hatEvents := make([]HatEvent, 0)

	for _, events := range v.mapping.Digital {
		for _, event := range events {
			switch e := event.(type) {
			case linux.Button:
				buttons = append(buttons, e)
			case linux.Key:
				keys = append(keys, e)
			case MSCScanCode:
				v.device.ActivateScanCode()
			case HatEvent:
				hatEvents = append(hatEvents, e)
			case virtual_device.AbsAxis:
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

	if len(hatEvents) > 0 {
		absoluteAxes = append(absoluteAxes, convertHatToAbsAxis(hatEvents)...)
	}

	v.device.SetEventButtons(buttons)
	v.device.SetEventKeys(keys)
	v.device.ActivateScanCode()
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
			v.device.SendAbsoluteEvent(uint16(e.Axis), e.Value)
		case virtual_device.AbsAxis:
			v.device.SendAbsoluteEvent(uint16(e.Axis), e.Max)
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
			v.device.SendAbsoluteEvent(uint16(e.Axis), 0)
		case virtual_device.AbsAxis:
			v.device.SendAbsoluteEvent(uint16(e.Axis), e.Min)
		default:
			fmt.Println("Unknown event type")
		}
	}
	v.device.SendSync()
}

func (v *virtualGamepad) moveStick(stick *MappingStick, x, y float32) {
	v.device.SendAbsoluteEvent(uint16(stick.X.Axis), stick.X.Denormalize(x))
	v.device.SendAbsoluteEvent(uint16(stick.Y.Axis), stick.Y.Denormalize(y))
	v.device.SendSync()

}

func (v *virtualGamepad) moveAxis(absAxis *virtual_device.AbsAxis, p float32) {
	v.device.SendAbsoluteEvent(uint16(absAxis.Axis), absAxis.Denormalize(p))
	v.device.SendSync()
}

func (v *virtualGamepad) MoveLeftStick(x, y float32) {
	if v.mapping.Analog != nil {
		v.moveStick(&v.mapping.Analog.Left, x, y)
	}
}

func (v *virtualGamepad) MoveLeftStickX(x float32) {
	if v.mapping.Analog != nil {
		v.moveAxis(&v.mapping.Analog.Left.X, x)
	}
}

func (v *virtualGamepad) MoveLeftStickY(y float32) {
	if v.mapping.Analog != nil {
		v.moveAxis(&v.mapping.Analog.Left.Y, y)
	}
}

func (v *virtualGamepad) MoveRightStick(x, y float32) {
	if v.mapping.Analog != nil {
		v.moveStick(&v.mapping.Analog.Right, x, y)
	}
}

func (v *virtualGamepad) MoveRightStickX(x float32) {
	if v.mapping.Analog != nil {
		v.moveAxis(&v.mapping.Analog.Right.X, x)
	}
}

func (v *virtualGamepad) MoveRightStickY(y float32) {
	if v.mapping.Analog != nil {
		v.moveAxis(&v.mapping.Analog.Right.Y, y)
	}
}
