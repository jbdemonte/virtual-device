package virtual_device

import (
	"errors"
	"fmt"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
	"github.com/jbdemonte/virtual-device/utils"
	"os"
	"syscall"
	"time"
	"unsafe"
)

type VirtualDevice interface {
	SetPath(path string) VirtualDevice
	SetMode(mode os.FileMode) VirtualDevice
	SetQueueLength(queueLength int) VirtualDevice
	SetBusType(busType linux.BusType) VirtualDevice
	SetVendor(vendor sdl.Vendor) VirtualDevice
	SetProduct(product sdl.Product) VirtualDevice
	SetVersion(version uint16) VirtualDevice
	SetName(name string) VirtualDevice
	SetEventKeys(keys []linux.Key) VirtualDevice
	SetEventButtons(buttons []linux.Button) VirtualDevice
	ActivateScanCode() VirtualDevice
	SetEventAbsoluteAxes(absoluteAxes []AbsAxis) VirtualDevice
	Register() error
	Unregister() error
	Send(evType, code uint16, value int32)
	SendSync()
	KeyPress(key uint16)
	KeyDown(key uint16)
	KeyUp(key uint16)
	SendAbsoluteEvent(absCode uint16, value int32)
	SendScanCode(value int32)
}

func NewVirtualDevice() VirtualDevice {
	return &virtualDevice{
		path:         "/dev/uinput",
		mode:         0660,
		queueLen:     1024,
		isRegistered: &utils.AtomicBool{},
	}
}

func (vd *virtualDevice) SetPath(path string) VirtualDevice {
	vd.path = path
	return vd
}

func (vd *virtualDevice) SetMode(mode os.FileMode) VirtualDevice {
	vd.mode = mode
	return vd
}

func (vd *virtualDevice) SetQueueLength(queueLen int) VirtualDevice {
	vd.queueLen = queueLen
	return vd
}

func (vd *virtualDevice) SetBusType(busType linux.BusType) VirtualDevice {
	vd.id.BusType = busType
	return vd
}

func (vd *virtualDevice) SetVendor(vendor uint16) VirtualDevice {
	vd.id.Vendor = vendor
	return vd
}

func (vd *virtualDevice) SetProduct(product uint16) VirtualDevice {
	vd.id.Product = product
	return vd
}

func (vd *virtualDevice) SetVersion(version uint16) VirtualDevice {
	vd.id.Version = version
	return vd
}

func (vd *virtualDevice) SetName(name string) VirtualDevice {
	vd.name = name
	return vd
}

func (vd *virtualDevice) SetEventKeys(keys []linux.Key) VirtualDevice {
	vd.events.keys = keys
	return vd
}

func (vd *virtualDevice) SetEventButtons(buttons []linux.Button) VirtualDevice {
	vd.events.buttons = buttons
	return vd
}

func (vd *virtualDevice) ActivateScanCode() VirtualDevice {
	vd.events.scanCode = true
	return vd
}

func (vd *virtualDevice) SetEventAbsoluteAxes(absoluteAxes []AbsAxis) VirtualDevice {
	vd.events.absoluteAxes = absoluteAxes
	return vd
}

func (vd *virtualDevice) Register() error {
	if vd.isRegistered.Get() {
		return nil
	}
	fd, err := os.OpenFile(vd.path, syscall.O_WRONLY|syscall.O_NONBLOCK, vd.mode)
	if err != nil {
		return errors.New("could not open device file")
	}

	vd.fd = fd

	err = vd.registerEvents()
	if err != nil {
		return vd.unregisterOnError(err)
	}

	err = vd.createDevice()
	if err != nil {
		return err
	}

	vd.pull()
	vd.isRegistered.Set(true)

	return nil
}

func (vd *virtualDevice) createDevice() (err error) {
	var fixedSizeName [linux.UINPUT_MAX_NAME_SIZE]byte
	copy(fixedSizeName[:], vd.name)
	if len(vd.name) < len(fixedSizeName) {
		fixedSizeName[len(vd.name)] = 0
	}

	var uinputDev linux.UInputUserDev
	copy(uinputDev.Name[:], fixedSizeName[:])
	uinputDev.ID = vd.id

	for _, event := range vd.events.absoluteAxes {
		uinputDev.AbsMin[event.Axis] = event.Min
		uinputDev.AbsMax[event.Axis] = event.Max
		uinputDev.AbsFlat[event.Axis] = event.Flat
		uinputDev.AbsFuzz[event.Axis] = event.Fuzz
	}

	_, err = vd.fd.Write((*[unsafe.Sizeof(uinputDev)]byte)(unsafe.Pointer(&uinputDev))[:])
	if err != nil {
		return fmt.Errorf("failed to write uidev struct to device file: %v", err)
	}

	err = ioctl(vd.fd, linux.UI_DEV_CREATE, uintptr(0))
	if err != nil {
		return fmt.Errorf("failed to create device: %v", err)
	}

	time.Sleep(time.Millisecond * 200)

	return nil
}

func (vd *virtualDevice) registerEvents() error {
	if (vd.events.keys == nil || len(vd.events.keys) == 0) && (vd.events.buttons == nil || len(vd.events.buttons) == 0) {
		return nil
	}

	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_KEY))
	if err != nil {
		return fmt.Errorf("failed to set UI_SET_EVBIT, EV_KEY: %v", err)
	}

	if vd.events.keys != nil {
		for _, key := range vd.events.keys {
			err = ioctl(vd.fd, linux.UI_SET_KEYBIT, uintptr(key))
			if err != nil {
				return fmt.Errorf("failed to register key 0x%x: %v", key, err)
			}
		}
	}

	if vd.events.buttons != nil {
		for _, button := range vd.events.buttons {
			err = ioctl(vd.fd, linux.UI_SET_KEYBIT, uintptr(button))
			if err != nil {
				return fmt.Errorf("failed to register button 0x%x: %v", button, err)
			}
		}
	}

	if vd.events.scanCode {
		err = ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_MSC))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_EVBIT, EV_MSC: %v", err)
		}

		err = ioctl(vd.fd, linux.UI_SET_MSCBIT, uintptr(linux.MSC_SCAN))
		if err != nil {
			return fmt.Errorf("failed to register MSC_SCAN: %v", err)
		}
	}

	if vd.events.absoluteAxes != nil {
		err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_ABS))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_EVBIT, EV_ABS: %v", err)
		}
		for _, event := range vd.events.absoluteAxes {
			err = ioctl(vd.fd, linux.UI_SET_ABSBIT, uintptr(event.Axis))
			if err != nil {
				return fmt.Errorf("failed to register axe 0x%x: %v", event.Axis, err)
			}
		}
	}

	return nil
}

func (vd *virtualDevice) pull() {
	vd.queue = make(chan *linux.InputEvent, vd.queueLen)

	go func() {
		for event := range vd.queue {
			err := vd.writeEvent(event)
			if err != nil {
				fmt.Printf("failed to write event: %v", err)
			}
		}
	}()
}

func (vd *virtualDevice) writeEvent(event *linux.InputEvent) error {
	buf := (*[unsafe.Sizeof(*event)]byte)(unsafe.Pointer(event))[:]
	n, err := vd.fd.Write(buf)
	if err != nil {
		return err
	}
	if n < linux.SizeofEvent {
		fmt.Fprintf(os.Stderr, "poll outbox: short write\n")
	}
	return nil
}

func (vd *virtualDevice) unregisterOnError(err error) error {
	uErr := vd.Unregister()
	return concatErrors(err, uErr)
}

func (vd *virtualDevice) closeQueue() {
	if vd.queue == nil {
		return
	}
	// wait for the queue to be flushed
	for len(vd.queue) > 0 {
		time.Sleep(time.Millisecond)
	}
	close(vd.queue)
	vd.queue = nil
}

func (vd *virtualDevice) releaseDevice() error {
	if vd.fd == nil {
		return nil
	}
	err := ioctl(vd.fd, linux.UI_DEV_DESTROY, 0)
	if err != nil {
		err = fmt.Errorf("failed to unregister the device: %v", err)
	}
	return err
}

func (vd *virtualDevice) closeDevice() error {
	if vd.fd == nil {
		return nil
	}
	err := vd.fd.Close()
	if err != nil {
		err = fmt.Errorf("failed to close the device: %v", err)
	}
	vd.fd = nil
	return err
}

func (vd *virtualDevice) Unregister() (err error) {
	vd.isRegistered.Set(false)

	vd.closeQueue()

	return concatErrors(
		vd.releaseDevice(),
		vd.closeDevice(),
	)
}

// Send an event to the device.
func (vd *virtualDevice) Send(evType, code uint16, value int32) {
	vd.queue <- &linux.InputEvent{
		Type:  evType,
		Code:  code,
		Value: value,
	}
}

func (vd *virtualDevice) SendSync() {
	vd.Send(uint16(linux.EV_SYN), uint16(linux.SYN_REPORT), 0)
}

func (vd *virtualDevice) KeyPress(key uint16) {
	vd.KeyDown(key)
	time.Sleep(time.Millisecond * 100)
	vd.KeyUp(key)
}

func (vd *virtualDevice) KeyDown(key uint16) {
	vd.Send(uint16(linux.EV_KEY), key, 1)
}

func (vd *virtualDevice) KeyUp(key uint16) {
	vd.Send(uint16(linux.EV_KEY), key, 0)
}

func (vd *virtualDevice) SendAbsoluteEvent(absCode uint16, value int32) {
	vd.Send(uint16(linux.EV_ABS), absCode, value)
}

func (vd *virtualDevice) SendScanCode(value int32) {
	vd.Send(uint16(linux.EV_MSC), uint16(linux.MSC_SCAN), value)
}
