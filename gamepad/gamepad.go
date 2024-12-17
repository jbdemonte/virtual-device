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

	vg.init()
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

func (vg *virtualGamepad) Register() error {
	return vg.device.Register()
}

func (vg *virtualGamepad) Unregister() error {
	return vg.device.Unregister()
}

func (vg *virtualGamepad) init() {
	buttons := make([]linux.Button, 0)
	keys := make([]linux.Key, 0)
	absoluteAxes := make([]virtual_device.AbsAxis, 0)
	hatEvents := make([]HatEvent, 0)

	for _, events := range vg.mapping.Digital {
		for _, event := range events {
			switch e := event.(type) {
			case linux.Button:
				buttons = append(buttons, e)
			case linux.Key:
				keys = append(keys, e)
			case MSCScanCode:
				vg.device.ActivateScanCode()
			case HatEvent:
				hatEvents = append(hatEvents, e)
			case virtual_device.AbsAxis:
				absoluteAxes = append(absoluteAxes, e)
			default:
				fmt.Println("Unknown event type")
			}
		}
	}

	if vg.mapping.Analog != nil {
		absoluteAxes = append(
			absoluteAxes,
			vg.mapping.Analog.Left.X,
			vg.mapping.Analog.Left.Y,
			vg.mapping.Analog.Right.X,
			vg.mapping.Analog.Right.Y,
		)
	}

	if len(hatEvents) > 0 {
		absoluteAxes = append(absoluteAxes, convertHatToAbsAxis(hatEvents)...)
	}

	vg.device.SetEventButtons(buttons)
	vg.device.SetEventKeys(keys)
	vg.device.ActivateScanCode()
	vg.device.SetEventAbsoluteAxes(absoluteAxes)
}

func (vg *virtualGamepad) Press(button Button) {
	events, exist := vg.mapping.Digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			vg.device.KeyDown(uint16(e))
		case linux.Key:
			vg.device.KeyDown(uint16(e))
		case MSCScanCode:
			vg.device.SendScanCode(int32(e))
		case HatEvent:
			vg.device.SendAbsoluteEvent(uint16(e.Axis), e.Value)
		case virtual_device.AbsAxis:
			vg.device.SendAbsoluteEvent(uint16(e.Axis), e.Max)
		default:
			fmt.Println("Unknown event type")
		}
	}
	vg.device.SendSync()
}

func (vg *virtualGamepad) Release(button Button) {
	events, exist := vg.mapping.Digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			vg.device.KeyUp(uint16(e))
		case linux.Key:
			vg.device.KeyUp(uint16(e))
		case MSCScanCode:
			vg.device.SendScanCode(int32(e))
		case HatEvent:
			vg.device.SendAbsoluteEvent(uint16(e.Axis), 0)
		case virtual_device.AbsAxis:
			vg.device.SendAbsoluteEvent(uint16(e.Axis), e.Min)
		default:
			fmt.Println("Unknown event type")
		}
	}
	vg.device.SendSync()
}

func (vg *virtualGamepad) moveStick(stick *MappingStick, x, y float32) {
	vg.device.SendAbsoluteEvent(uint16(stick.X.Axis), stick.X.Denormalize(x))
	vg.device.SendAbsoluteEvent(uint16(stick.Y.Axis), stick.Y.Denormalize(y))
	vg.device.SendSync()

}

func (vg *virtualGamepad) moveAxis(absAxis *virtual_device.AbsAxis, p float32) {
	vg.device.SendAbsoluteEvent(uint16(absAxis.Axis), absAxis.Denormalize(p))
	vg.device.SendSync()
}

func (vg *virtualGamepad) MoveLeftStick(x, y float32) {
	if vg.mapping.Analog != nil {
		vg.moveStick(&vg.mapping.Analog.Left, x, y)
	}
}

func (vg *virtualGamepad) MoveLeftStickX(x float32) {
	if vg.mapping.Analog != nil {
		vg.moveAxis(&vg.mapping.Analog.Left.X, x)
	}
}

func (vg *virtualGamepad) MoveLeftStickY(y float32) {
	if vg.mapping.Analog != nil {
		vg.moveAxis(&vg.mapping.Analog.Left.Y, y)
	}
}

func (vg *virtualGamepad) MoveRightStick(x, y float32) {
	if vg.mapping.Analog != nil {
		vg.moveStick(&vg.mapping.Analog.Right, x, y)
	}
}

func (vg *virtualGamepad) MoveRightStickX(x float32) {
	if vg.mapping.Analog != nil {
		vg.moveAxis(&vg.mapping.Analog.Right.X, x)
	}
}

func (vg *virtualGamepad) MoveRightStickY(y float32) {
	if vg.mapping.Analog != nil {
		vg.moveAxis(&vg.mapping.Analog.Right.Y, y)
	}
}
