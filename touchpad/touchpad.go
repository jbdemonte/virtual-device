package touchpad

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

// https://www.kernel.org/doc/Documentation/input/multi-touch-protocol.txt

type TouchSlot struct {
	Slot     int
	X        float32
	Y        float32
	Pressure float32
}

type VirtualTouchpad interface {
	Register() error
	Unregister() error

	Touch(x, y, pressure float32)
	MultiTouch(touchSlots []TouchSlot) []TouchSlot

	ButtonPress(button linux.Button)
	ButtonRelease(button linux.Button)

	Click(btn linux.Button)
	DoubleClick(btn linux.Button)

	ClickLeft()
	ClickRight()

	DoubleClickLeft()
	DoubleClickRight()
}

type VirtualTouchpadFactory interface {
	WithDevice(device virtual_device.VirtualDevice) VirtualTouchpadFactory
	WithClickDelay(delay int) VirtualTouchpadFactory
	WithDoubleClickDelay(delay int) VirtualTouchpadFactory
	WithAxes(absoluteAxes []virtual_device.AbsAxis) VirtualTouchpadFactory
	WithButtons(buttons []linux.Button) VirtualTouchpadFactory
	WithProperties(properties []linux.InputProp) VirtualTouchpadFactory
	WithLegacyMultitouch() VirtualTouchpadFactory
	Create() VirtualTouchpad
}

func NewVirtualTouchpadFactory() VirtualTouchpadFactory {
	return &virtualTouchpadFactory{
		clickDelay:       -1,
		doubleClickDelay: -1,
	}
}

type virtualTouchpadFactory struct {
	device           virtual_device.VirtualDevice
	clickDelay       int
	doubleClickDelay int
	axes             []virtual_device.AbsAxis
	buttons          []linux.Button
	properties       []linux.InputProp
	protocolA        bool
}

func (f *virtualTouchpadFactory) WithDevice(device virtual_device.VirtualDevice) VirtualTouchpadFactory {
	f.device = device
	return f
}

func (f *virtualTouchpadFactory) WithClickDelay(delay int) VirtualTouchpadFactory {
	f.clickDelay = delay
	return f
}

func (f *virtualTouchpadFactory) WithDoubleClickDelay(delay int) VirtualTouchpadFactory {
	f.doubleClickDelay = delay
	return f
}

func (f *virtualTouchpadFactory) WithAxes(axes []virtual_device.AbsAxis) VirtualTouchpadFactory {
	f.axes = axes
	return f
}

func (f *virtualTouchpadFactory) WithButtons(buttons []linux.Button) VirtualTouchpadFactory {
	f.buttons = buttons
	return f
}

func (f *virtualTouchpadFactory) WithProperties(properties []linux.InputProp) VirtualTouchpadFactory {
	f.properties = properties
	return f
}

func (f *virtualTouchpadFactory) WithLegacyMultitouch() VirtualTouchpadFactory {
	f.protocolA = true
	return f
}

func (f *virtualTouchpadFactory) Create() VirtualTouchpad {
	clickDelay := f.clickDelay
	if clickDelay < 0 {
		clickDelay = 50
	}

	doubleClickDelay := f.doubleClickDelay
	if doubleClickDelay < 0 {
		doubleClickDelay = 250
	}

	f.device.WithAbsAxes(f.axes)
	f.device.WithButtons(f.buttons)
	f.device.WithProperties(f.properties)

	axes := map[linux.AbsoluteAxis]virtual_device.AbsAxis{}

	for _, axis := range f.axes {
		axes[axis.Axis] = axis
	}

	return &virtualTouchpad{
		device:           f.device,
		clickDelay:       clickDelay,
		doubleClickDelay: doubleClickDelay,
		axes:             axes,
		currentSlots:     map[int]bool{},
		protocolA:        f.protocolA,
	}
}

type virtualTouchpad struct {
	device           virtual_device.VirtualDevice
	clickDelay       int
	doubleClickDelay int
	axes             map[linux.AbsoluteAxis]virtual_device.AbsAxis
	currentSlots     map[int]bool
	fingerCount      int
	protocolA        bool
}

func (vt *virtualTouchpad) Register() error {
	return vt.device.Register()
}

func (vt *virtualTouchpad) Unregister() error {
	return vt.device.Unregister()
}

func (vt *virtualTouchpad) absDenormalized(axis linux.AbsoluteAxis, x float32) {
	axisAbs, exits := vt.axes[axis]
	if exits {
		vt.device.Abs(uint16(axis), axisAbs.Denormalize(x))
	}
}

func (vt *virtualTouchpad) Touch(x, y, pressure float32) {
	vt.absDenormalized(linux.ABS_X, x)
	vt.absDenormalized(linux.ABS_Y, y)
	vt.absDenormalized(linux.ABS_PRESSURE, pressure)
	vt.device.SyncReport()
}

func (vt *virtualTouchpad) assignSlotIfNeeded(touchSlots []TouchSlot) []TouchSlot {
	reserved := make(map[int]bool)
	for _, ts := range touchSlots {
		reserved[ts.Slot] = true
	}
	used := make(map[int]bool)
	for i, _ := range touchSlots {
		slot := touchSlots[i].Slot
		if used[slot] {
			for reserved[slot] || used[slot] {
				slot = slot + 1
			}
		}
		used[slot] = true
		touchSlots[i].Slot = slot
	}
	return touchSlots
}

// https://www.kernel.org/doc/Documentation/input/multi-touch-protocol.txt
func (vt *virtualTouchpad) multiTouchA(touchSlots []TouchSlot) []TouchSlot {
	for _, ts := range touchSlots {
		vt.absDenormalized(linux.ABS_MT_POSITION_X, ts.X)
		vt.absDenormalized(linux.ABS_MT_POSITION_Y, ts.Y)
		vt.device.Sync(linux.SYN_MT_REPORT)
	}
	if len(touchSlots) == 0 {
		vt.device.Sync(linux.SYN_MT_REPORT)
	}
	vt.device.SyncReport()
	return touchSlots
}

func (vt *virtualTouchpad) multiTouchB(touchSlots []TouchSlot) []TouchSlot {
	watcher := vt.createFingerCountWatcher()
	touchSlots = vt.assignSlotIfNeeded(touchSlots)

	for _, ts := range touchSlots {
		vt.device.Abs(uint16(linux.ABS_MT_SLOT), int32(ts.Slot))
		if ts.Pressure == 0 {
			// release the Slot
			vt.currentSlots[ts.Slot] = false
			vt.device.Abs(uint16(linux.ABS_MT_TRACKING_ID), int32(-1))
			vt.fingerCount = vt.fingerCount - 1
		} else if vt.currentSlots[ts.Slot] == false {
			// lock the Slot
			vt.currentSlots[ts.Slot] = true
			vt.device.Abs(uint16(linux.ABS_MT_TRACKING_ID), int32(ts.Slot))
			vt.fingerCount = vt.fingerCount + 1
		}
		if ts.Pressure > 0 {
			vt.absDenormalized(linux.ABS_MT_POSITION_X, ts.X)
			vt.absDenormalized(linux.ABS_MT_POSITION_Y, ts.Y)
		}
		vt.absDenormalized(linux.ABS_PRESSURE, ts.Pressure)
	}

	watcher()
	vt.device.SyncReport()
	return touchSlots
}

func (vt *virtualTouchpad) toggleFingerCount(count int, value bool) {
	if count == 0 {
		return
	}
	buttons := []linux.Button{
		linux.BTN_TOOL_FINGER, linux.BTN_TOOL_DOUBLETAP,
		linux.BTN_TOOL_TRIPLETAP, linux.BTN_TOOL_QUADTAP,
		linux.BTN_TOOL_QUINTTAP,
	}
	button := buttons[count-1]
	if value {
		vt.device.KeyPress(uint16(button))
	} else {
		vt.device.KeyRelease(uint16(button))
	}
}

func (vt *virtualTouchpad) createFingerCountWatcher() func() {
	previousCount := vt.fingerCount
	return func() {
		if previousCount != vt.fingerCount {
			vt.toggleFingerCount(previousCount, false)
			vt.toggleFingerCount(vt.fingerCount, true)
		}
	}
}

func (vt *virtualTouchpad) MultiTouch(touchSlots []TouchSlot) []TouchSlot {
	if vt.protocolA {
		return vt.multiTouchA(touchSlots)
	}
	return vt.multiTouchB(touchSlots)
}

func (vt *virtualTouchpad) ButtonPress(button linux.Button) {
	vt.device.KeyPress(uint16(button))
	vt.device.SyncReport()
}

func (vt *virtualTouchpad) ButtonRelease(button linux.Button) {
	vt.device.KeyRelease(uint16(button))
	vt.device.SyncReport()
}

func (vt *virtualTouchpad) Click(btn linux.Button) {
	vt.ButtonPress(btn)
	time.Sleep(time.Millisecond * time.Duration(vt.clickDelay))
	vt.ButtonRelease(btn)
}

func (vt *virtualTouchpad) DoubleClick(btn linux.Button) {
	vt.Click(btn)
	time.Sleep(time.Millisecond * time.Duration(vt.doubleClickDelay))
	vt.Click(btn)
}

func (vt *virtualTouchpad) ClickLeft() {
	vt.Click(linux.BTN_LEFT)
}

func (vt *virtualTouchpad) ClickRight() {
	vt.Click(linux.BTN_RIGHT)
}

func (vt *virtualTouchpad) DoubleClickLeft() {
	vt.DoubleClick(linux.BTN_LEFT)
}

func (vt *virtualTouchpad) DoubleClickRight() {
	vt.DoubleClick(linux.BTN_RIGHT)
}
