package mouse

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

type VirtualMouse interface {
	Register() error
	Unregister() error
	Move(x, y int32)
	MoveX(x int32)
	MoveY(Y int32)
	Wheel(delta int32)
	HWheel(delta int32)
	ButtonDown(button linux.Button)
	ButtonUp(button linux.Button)
	ScrollUp()
	ScrollDown()
	ScrollLeft()
	ScrollRight()
	Click(btn linux.Button)
	DoubleClick(btn linux.Button)
	ClickLeft()
	ClickRight()
	ClickMiddle()
	DoubleClickLeft()
	DoubleClickRight()
	DoubleClickMiddle()
}

type virtualMouse struct {
	device virtual_device.VirtualDevice
	config Config
}

func createVirtualMouse(device virtual_device.VirtualDevice, config Config) VirtualMouse {
	if config.clickDelay == 0 {
		config.clickDelay = 50
	}
	if config.doubleClickDelay == 0 {
		config.doubleClickDelay = 250
	}
	vm := &virtualMouse{device, config}
	return vm
}

func (vm *virtualMouse) Register() error {
	return vm.device.Register()
}

func (vm *virtualMouse) Unregister() error {
	return vm.device.Unregister()
}

func (vm *virtualMouse) Move(x, y int32) {
	vm.device.SendRelativeEvent(uint16(linux.REL_X), x)
	vm.device.SendRelativeEvent(uint16(linux.REL_Y), y)
	vm.device.SendSync()
}

func (vm *virtualMouse) MoveX(x int32) {
	vm.device.SendRelativeEvent(uint16(linux.REL_X), x)
	vm.device.SendSync()
}

func (vm *virtualMouse) MoveY(y int32) {
	vm.device.SendRelativeEvent(uint16(linux.REL_Y), y)
	vm.device.SendSync()
}

func (vm *virtualMouse) Wheel(delta int32) {
	vm.device.SendRelativeEvent(uint16(linux.REL_WHEEL), delta)
	if vm.config.highResStep != 0 {
		vm.device.SendRelativeEvent(uint16(linux.REL_WHEEL_HI_RES), delta*vm.config.highResStep)
	}
	vm.device.SendSync()
}

func (vm *virtualMouse) HWheel(delta int32) {
	vm.device.SendRelativeEvent(uint16(linux.REL_HWHEEL), delta)
	if vm.config.highResHStep != 0 {
		vm.device.SendRelativeEvent(uint16(linux.REL_HWHEEL_HI_RES), delta*vm.config.highResHStep)
	}
	vm.device.SendSync()
}

func (vm *virtualMouse) ButtonDown(button linux.Button) {
	vm.device.KeyDown(uint16(button))
	vm.device.SendSync()
}

func (vm *virtualMouse) ButtonUp(button linux.Button) {
	vm.device.KeyUp(uint16(button))
	vm.device.SendSync()
}

func (vm *virtualMouse) ScrollUp() {
	vm.Wheel(1)
}

func (vm *virtualMouse) ScrollDown() {
	vm.Wheel(-1)
}

func (vm *virtualMouse) ScrollLeft() {
	vm.HWheel(-1)
}

func (vm *virtualMouse) ScrollRight() {
	vm.HWheel(1)
}

func (vm *virtualMouse) Click(btn linux.Button) {
	vm.ButtonDown(btn)
	time.Sleep(time.Millisecond * time.Duration(vm.config.clickDelay))
	vm.ButtonUp(btn)
}

func (vm *virtualMouse) DoubleClick(btn linux.Button) {
	vm.Click(btn)
	time.Sleep(time.Millisecond * time.Duration(vm.config.doubleClickDelay))
	vm.Click(btn)
}

func (vm *virtualMouse) ClickLeft() {
	vm.Click(linux.BTN_LEFT)
}

func (vm *virtualMouse) ClickRight() {
	vm.Click(linux.BTN_RIGHT)
}

func (vm *virtualMouse) ClickMiddle() {
	vm.Click(linux.BTN_MIDDLE)
}

func (vm *virtualMouse) DoubleClickLeft() {
	vm.DoubleClick(linux.BTN_LEFT)
}

func (vm *virtualMouse) DoubleClickRight() {
	vm.DoubleClick(linux.BTN_RIGHT)
}

func (vm *virtualMouse) DoubleClickMiddle() {
	vm.DoubleClick(linux.BTN_MIDDLE)
}
