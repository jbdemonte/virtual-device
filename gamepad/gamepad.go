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

	Send(evType, code uint16, value int32)
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

	withScanCode := false

	var init func(event InputEvent)

	init = func(event InputEvent) {
		switch e := event.(type) {
		case []InputEvent:
			for _, item := range e {
				init(item)
			}
		case linux.Button:
			buttons = append(buttons, e)
		case linux.Key:
			keys = append(keys, e)
		case MSCScanCode:
			withScanCode = true
		case HatEvent:
			hatEvents = append(hatEvents, e)
		case virtual_device.AbsAxis:
			absoluteAxes = append(absoluteAxes, e)
		default:
			fmt.Println("Unknown event type")
		}
	}

	for _, event := range vg.digital {
		init(event)
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
	vg.device.WithAbsAxes(absoluteAxes)

	if withScanCode {
		vg.device.WithMiscEvents([]linux.MiscEvent{linux.MSC_SCAN})
	}
}

func (vg *virtualGamepad) Press(button Button) {
	var press func(event InputEvent)

	press = func(event InputEvent) {
		switch e := event.(type) {
		case []InputEvent:
			for _, item := range e {
				press(item)
			}
		case linux.Button:
			vg.device.PressButton(e)
		case linux.Key:
			vg.device.PressKey(e)
		case MSCScanCode:
			vg.device.SendMiscEvent(linux.MSC_SCAN, int32(e))
		case HatEvent:
			vg.device.SendAbsoluteEvent(e.Axis, e.Value)
		case virtual_device.AbsAxis:
			vg.device.SendAbsoluteEvent(e.Axis, e.Max)
		default:
			fmt.Println("Unknown event type")
		}
	}

	event, exist := vg.digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}

	press(event)
	vg.device.SyncReport()
}

func (vg *virtualGamepad) Release(button Button) {
	var release func(event InputEvent)

	release = func(event InputEvent) {
		switch e := event.(type) {
		case []InputEvent:
			for _, item := range e {
				release(item)
			}
		case linux.Button:
			vg.device.ReleaseButton(e)
		case linux.Key:
			vg.device.ReleaseKey(e)
		case MSCScanCode:
			vg.device.SendMiscEvent(linux.MSC_SCAN, int32(e))
		case HatEvent:
			vg.device.SendAbsoluteEvent(e.Axis, 0)
		case virtual_device.AbsAxis:
			vg.device.SendAbsoluteEvent(e.Axis, e.Min)
		default:
			fmt.Println("Unknown event type")
		}
	}

	event, exist := vg.digital[button]
	if !exist {
		fmt.Printf("button not assigned (0x%x)\n", button)
	}

	release(event)
	vg.device.SyncReport()
}

func (vg *virtualGamepad) moveStick(stick *MappingStick, x, y float32) {
	vg.device.SendAbsoluteEvent(stick.X.Axis, stick.X.Denormalize(x))
	vg.device.SendAbsoluteEvent(stick.Y.Axis, stick.Y.Denormalize(y))
	vg.device.SyncReport()

}

func (vg *virtualGamepad) moveAxis(absAxis *virtual_device.AbsAxis, p float32) {
	vg.device.SendAbsoluteEvent(absAxis.Axis, absAxis.Denormalize(p))
	vg.device.SyncReport()
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

func (vg *virtualGamepad) Send(evType, code uint16, value int32) {
	vg.device.Send(evType, code, value)
}
