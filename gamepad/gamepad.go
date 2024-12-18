package gamepad

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

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

type VirtualGamepadFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualGamepadFactory
	WithDigital(mapping MappingDigital) VirtualGamepadFactory
	WithLeftStick(mapping MappingStick) VirtualGamepadFactory
	WithRightStick(mapping MappingStick) VirtualGamepadFactory
	Create() VirtualGamepad
}

type virtualGamepadFactory struct {
	device     virtual_device.VirtualDevice
	digital    MappingDigital
	leftStick  *MappingStick
	rightStick *MappingStick
}

func NewVirtualGamepadFactory() VirtualGamepadFactory {
	return &virtualGamepadFactory{}
}

func (f *virtualGamepadFactory) WithDevice(device virtual_device.VirtualDevice) VirtualGamepadFactory {
	f.device = device
	return f
}

func (f *virtualGamepadFactory) WithDigital(mapping MappingDigital) VirtualGamepadFactory {
	f.digital = mapping
	return f
}

func (f *virtualGamepadFactory) WithLeftStick(mapping MappingStick) VirtualGamepadFactory {
	f.leftStick = &mapping
	return f
}

func (f *virtualGamepadFactory) WithRightStick(mapping MappingStick) VirtualGamepadFactory {
	f.rightStick = &mapping
	return f
}

func (f *virtualGamepadFactory) Create() VirtualGamepad {
	vg := &virtualGamepad{
		device:     f.device,
		digital:    f.digital,
		leftStick:  f.leftStick,
		rightStick: f.rightStick,
	}

	vg.init()
	return vg
}

type virtualGamepad struct {
	device     virtual_device.VirtualDevice
	digital    MappingDigital
	leftStick  *MappingStick
	rightStick *MappingStick
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

	for _, events := range vg.digital {
		for _, event := range events {
			switch e := event.(type) {
			case linux.Button:
				buttons = append(buttons, e)
			case linux.Key:
				keys = append(keys, e)
			case MSCScanCode:
				vg.device.WithScanCode()
			case HatEvent:
				hatEvents = append(hatEvents, e)
			case virtual_device.AbsAxis:
				absoluteAxes = append(absoluteAxes, e)
			default:
				fmt.Println("Unknown event type")
			}
		}
	}

	if vg.leftStick != nil {
		absoluteAxes = append(
			absoluteAxes,
			vg.leftStick.X,
			vg.leftStick.Y,
		)
	}

	if vg.rightStick != nil {
		absoluteAxes = append(
			absoluteAxes,
			vg.rightStick.X,
			vg.rightStick.Y,
		)
	}

	if len(hatEvents) > 0 {
		absoluteAxes = append(absoluteAxes, convertHatToAbsAxis(hatEvents)...)
	}

	vg.device.WithButtons(buttons)
	vg.device.WithKeys(keys)
	vg.device.WithScanCode()
	vg.device.WithAbsAxes(absoluteAxes)
}

func (vg *virtualGamepad) Press(button Button) {
	events, exist := vg.digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			vg.device.KeyPress(uint16(e))
		case linux.Key:
			vg.device.KeyPress(uint16(e))
		case MSCScanCode:
			vg.device.ScanCode(int32(e))
		case HatEvent:
			vg.device.Abs(uint16(e.Axis), e.Value)
		case virtual_device.AbsAxis:
			vg.device.Abs(uint16(e.Axis), e.Max)
		default:
			fmt.Println("Unknown event type")
		}
	}
	vg.device.Sync()
}

func (vg *virtualGamepad) Release(button Button) {
	events, exist := vg.digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}
	for _, event := range events {
		switch e := event.(type) {
		case linux.Button:
			vg.device.KeyRelease(uint16(e))
		case linux.Key:
			vg.device.KeyRelease(uint16(e))
		case MSCScanCode:
			vg.device.ScanCode(int32(e))
		case HatEvent:
			vg.device.Abs(uint16(e.Axis), 0)
		case virtual_device.AbsAxis:
			vg.device.Abs(uint16(e.Axis), e.Min)
		default:
			fmt.Println("Unknown event type")
		}
	}
	vg.device.Sync()
}

func (vg *virtualGamepad) moveStick(stick *MappingStick, x, y float32) {
	vg.device.Abs(uint16(stick.X.Axis), stick.X.Denormalize(x))
	vg.device.Abs(uint16(stick.Y.Axis), stick.Y.Denormalize(y))
	vg.device.Sync()

}

func (vg *virtualGamepad) moveAxis(absAxis *virtual_device.AbsAxis, p float32) {
	vg.device.Abs(uint16(absAxis.Axis), absAxis.Denormalize(p))
	vg.device.Sync()
}

func (vg *virtualGamepad) MoveLeftStick(x, y float32) {
	if vg.leftStick != nil {
		vg.moveStick(vg.leftStick, x, y)
	}
}

func (vg *virtualGamepad) MoveLeftStickX(x float32) {
	if vg.leftStick != nil {
		vg.moveAxis(&vg.leftStick.X, x)
	}
}

func (vg *virtualGamepad) MoveLeftStickY(y float32) {
	if vg.leftStick != nil {
		vg.moveAxis(&vg.leftStick.Y, y)
	}
}

func (vg *virtualGamepad) MoveRightStick(x, y float32) {
	if vg.rightStick != nil {
		vg.moveStick(vg.rightStick, x, y)
	}
}

func (vg *virtualGamepad) MoveRightStickX(x float32) {
	if vg.rightStick != nil {
		vg.moveAxis(&vg.rightStick.X, x)
	}
}

func (vg *virtualGamepad) MoveRightStickY(y float32) {
	if vg.rightStick != nil {
		vg.moveAxis(&vg.rightStick.Y, y)
	}
}
