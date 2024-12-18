package mouse

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

type VirtualMouse interface {
	Register() error
	Unregister() error

	Move(deltaX, deltaY int32)
	MoveX(deltaX int32)
	MoveY(deltaY int32)

	ScrollVertical(delta int32)
	ScrollHorizontal(delta int32)

	ButtonPress(button linux.Button)
	ButtonRelease(button linux.Button)

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

type VirtualMouseFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualMouseFactory
	WithClickDelay(delay int) VirtualMouseFactory
	WithDoubleClickDelay(delay int) VirtualMouseFactory
	WithHighResStepVertical(step int32) VirtualMouseFactory
	WithHighResStepHorizontal(step int32) VirtualMouseFactory
	Create() VirtualMouse
}

func NewVirtualMouseFactory() VirtualMouseFactory {
	return &virtualMouseFactory{
		clickDelay:       -1,
		doubleClickDelay: -1,
	}
}

type virtualMouseFactory struct {
	device                virtual_device.VirtualDevice
	highResStepVertical   int32
	highResStepHorizontal int32
	clickDelay            int
	doubleClickDelay      int
}

func (f *virtualMouseFactory) WithDevice(device virtual_device.VirtualDevice) VirtualMouseFactory {
	f.device = device
	return f
}

func (f *virtualMouseFactory) WithClickDelay(delay int) VirtualMouseFactory {
	f.clickDelay = delay
	return f
}

func (f *virtualMouseFactory) WithDoubleClickDelay(delay int) VirtualMouseFactory {
	f.doubleClickDelay = delay
	return f
}

func (f *virtualMouseFactory) WithHighResStepVertical(step int32) VirtualMouseFactory {
	f.highResStepVertical = step
	return f
}

func (f *virtualMouseFactory) WithHighResStepHorizontal(step int32) VirtualMouseFactory {
	f.highResStepHorizontal = step
	return f
}

func (f *virtualMouseFactory) Create() VirtualMouse {
	clickDelay := f.clickDelay
	if clickDelay < 0 {
		clickDelay = 50
	}

	doubleClickDelay := f.doubleClickDelay
	if doubleClickDelay < 0 {
		doubleClickDelay = 250
	}
	return &virtualMouse{
		device:                f.device,
		clickDelay:            clickDelay,
		doubleClickDelay:      doubleClickDelay,
		highResStepVertical:   f.highResStepVertical,
		highResStepHorizontal: f.highResStepHorizontal,
	}
}

type virtualMouse struct {
	device                virtual_device.VirtualDevice
	highResStepVertical   int32
	highResStepHorizontal int32
	clickDelay            int
	doubleClickDelay      int
}

func (vm *virtualMouse) Register() error {
	return vm.device.Register()
}

func (vm *virtualMouse) Unregister() error {
	return vm.device.Unregister()
}

func (vm *virtualMouse) Move(deltaX, deltaY int32) {
	vm.device.Rel(uint16(linux.REL_X), deltaX)
	vm.device.Rel(uint16(linux.REL_Y), deltaY)
	vm.device.Sync()
}

func (vm *virtualMouse) MoveX(deltaX int32) {
	vm.device.Rel(uint16(linux.REL_X), deltaX)
	vm.device.Sync()
}

func (vm *virtualMouse) MoveY(deltaY int32) {
	vm.device.Rel(uint16(linux.REL_Y), deltaY)
	vm.device.Sync()
}

func (vm *virtualMouse) ScrollVertical(delta int32) {
	vm.device.Rel(uint16(linux.REL_WHEEL), delta)
	if vm.highResStepVertical != 0 {
		vm.device.Rel(uint16(linux.REL_WHEEL_HI_RES), delta*vm.highResStepVertical)
	}
	vm.device.Sync()
}

func (vm *virtualMouse) ScrollHorizontal(delta int32) {
	vm.device.Rel(uint16(linux.REL_HWHEEL), delta)
	if vm.highResStepHorizontal != 0 {
		vm.device.Rel(uint16(linux.REL_HWHEEL_HI_RES), delta*vm.highResStepHorizontal)
	}
	vm.device.Sync()
}

func (vm *virtualMouse) ButtonPress(button linux.Button) {
	vm.device.KeyPress(uint16(button))
	vm.device.Sync()
}

func (vm *virtualMouse) ButtonRelease(button linux.Button) {
	vm.device.KeyRelease(uint16(button))
	vm.device.Sync()
}

func (vm *virtualMouse) ScrollUp() {
	vm.ScrollVertical(1)
}

func (vm *virtualMouse) ScrollDown() {
	vm.ScrollVertical(-1)
}

func (vm *virtualMouse) ScrollLeft() {
	vm.ScrollHorizontal(-1)
}

func (vm *virtualMouse) ScrollRight() {
	vm.ScrollHorizontal(1)
}

func (vm *virtualMouse) Click(btn linux.Button) {
	vm.ButtonPress(btn)
	time.Sleep(time.Millisecond * time.Duration(vm.clickDelay))
	vm.ButtonRelease(btn)
}

func (vm *virtualMouse) DoubleClick(btn linux.Button) {
	vm.Click(btn)
	time.Sleep(time.Millisecond * time.Duration(vm.doubleClickDelay))
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
