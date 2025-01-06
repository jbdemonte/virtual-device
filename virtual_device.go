package virtual_device

import (
	"errors"
	"fmt"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
	"github.com/jbdemonte/virtual-device/utils"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type VirtualDevice interface {
	WithPath(path string) VirtualDevice
	WithMode(mode os.FileMode) VirtualDevice
	WithQueueLen(queueLen int) VirtualDevice
	WithBusType(busType linux.BusType) VirtualDevice
	WithVendor(vendor sdl.Vendor) VirtualDevice
	WithProduct(product sdl.Product) VirtualDevice
	WithVersion(version uint16) VirtualDevice
	WithName(name string) VirtualDevice
	WithKeys(keys []linux.Key) VirtualDevice
	WithButtons(buttons []linux.Button) VirtualDevice
	WithScanCode() VirtualDevice
	WithAbsAxes(absoluteAxes []AbsAxis) VirtualDevice
	WithRelAxes(relativeAxes []linux.RelativeAxis) VirtualDevice
	WithRepeat(delay, period int32) VirtualDevice
	WithLEDs(leds []linux.Led) VirtualDevice
	WithProperties(properties []linux.InputProp) VirtualDevice
	WithMiscEvents(events []linux.MiscEvent) VirtualDevice

	Register() error
	Unregister() error

	Send(evType, code uint16, value int32)
	Sync(evType linux.SyncEvent)
	SyncReport()
	PressKey(key linux.Key)
	ReleaseKey(key linux.Key)
	PressButton(key linux.Button)
	ReleaseButton(key linux.Button)
	SendAbsoluteEvent(axis linux.AbsoluteAxis, value int32)
	SendRelativeEvent(axis linux.RelativeAxis, value int32)
	SendMiscEvent(event linux.MiscEvent, value int32)
	SetLed(led linux.Led, state bool)
}

func NewVirtualDevice() VirtualDevice {
	return &virtualDevice{
		path:         "/dev/uinput",
		mode:         0660,
		queueLen:     1024,
		isRegistered: &utils.AtomicBool{},
	}
}

func (vd *virtualDevice) WithPath(path string) VirtualDevice {
	vd.path = path
	return vd
}

func (vd *virtualDevice) WithMode(mode os.FileMode) VirtualDevice {
	vd.mode = mode
	return vd
}

func (vd *virtualDevice) WithQueueLen(queueLen int) VirtualDevice {
	vd.queueLen = queueLen
	return vd
}

func (vd *virtualDevice) WithBusType(busType linux.BusType) VirtualDevice {
	vd.id.BusType = busType
	return vd
}

func (vd *virtualDevice) WithVendor(vendor uint16) VirtualDevice {
	vd.id.Vendor = vendor
	return vd
}

func (vd *virtualDevice) WithProduct(product uint16) VirtualDevice {
	vd.id.Product = product
	return vd
}

func (vd *virtualDevice) WithVersion(version uint16) VirtualDevice {
	vd.id.Version = version
	return vd
}

func (vd *virtualDevice) WithName(name string) VirtualDevice {
	vd.name = name
	return vd
}

func (vd *virtualDevice) WithKeys(keys []linux.Key) VirtualDevice {
	vd.config.keys = keys
	return vd
}

func (vd *virtualDevice) WithButtons(buttons []linux.Button) VirtualDevice {
	vd.config.buttons = buttons
	return vd
}

func (vd *virtualDevice) WithScanCode() VirtualDevice {
	vd.config.scanCode = true
	return vd
}

func (vd *virtualDevice) WithAbsAxes(absoluteAxes []AbsAxis) VirtualDevice {
	vd.config.absoluteAxes = absoluteAxes
	return vd
}

func (vd *virtualDevice) WithRelAxes(relativeAxes []linux.RelativeAxis) VirtualDevice {
	vd.config.relativeAxes = relativeAxes
	return vd
}

func (vd *virtualDevice) WithRepeat(delay, period int32) VirtualDevice {
	vd.config.repeat = &Repeat{delay, period}
	return vd

}

func (vd *virtualDevice) WithLEDs(leds []linux.Led) VirtualDevice {
	vd.config.leds = leds
	return vd
}

func (vd *virtualDevice) WithProperties(properties []linux.InputProp) VirtualDevice {
	vd.config.properties = properties
	return vd
}

func (vd *virtualDevice) WithMiscEvents(miscEvents []linux.MiscEvent) VirtualDevice {
	vd.config.miscEvents = miscEvents
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

	steps := []func() error{
		vd.registerKeys,
		vd.registerScanCode,
		vd.registerAxes,
		vd.registerProperties,
		vd.registerMiscEvents,
		vd.registerLeds,
		vd.createDevice,
	}

	for _, step := range steps {
		if err := step(); err != nil {
			return vd.unregisterOnError(err)
		}
	}

	vd.pull()

	vd.isRegistered.Set(true)

	return nil
}

func (vd *virtualDevice) fetchEventPath() (string, error) {
	sysInputDir := "/sys/devices/virtual/input/"
	path := make([]byte, 65) // 64 bytes + null byte

	err := ioctl(vd.fd, linux.UI_GET_SYSNAME(64), uintptr(unsafe.Pointer(&path[0])))
	if err != nil {
		return "", fmt.Errorf("ioctl uiGetSysname failed: %v", err)
	}

	sysPath := sysInputDir + strings.TrimRight(string(path), "\x00")

	files, err := os.ReadDir(sysPath)
	if err != nil {
		return "", fmt.Errorf("unable to read directory %s: %v", sysPath, err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "event") {
			return fmt.Sprintf("/dev/input/%s", file.Name()), nil
		}
	}

	return "", fmt.Errorf("no event file found in %s", sysPath)
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

	setAbsResolution := false

	for _, event := range vd.config.absoluteAxes {
		uinputDev.AbsMin[event.Axis] = event.Min
		uinputDev.AbsMax[event.Axis] = event.Max
		uinputDev.AbsFlat[event.Axis] = event.Flat
		uinputDev.AbsFuzz[event.Axis] = event.Fuzz
		if event.Resolution > 0 {
			setAbsResolution = true
		}
	}

	_, err = vd.fd.Write((*[unsafe.Sizeof(uinputDev)]byte)(unsafe.Pointer(&uinputDev))[:])
	if err != nil {
		return fmt.Errorf("failed to write uidev struct to device file: %v", err)
	}

	err = ioctl(vd.fd, linux.UI_DEV_CREATE, uintptr(0))
	if err != nil {
		return fmt.Errorf("failed to create device: %v", err)
	}

	vd.eventPath, err = vd.fetchEventPath()
	if err != nil {
		return fmt.Errorf("fetchEventPath: %v", err)
	}

	err = utils.WaitForEventFile(vd.eventPath, 500*time.Millisecond)
	if err != nil {
		return fmt.Errorf("WaitForEventFile: %v", err)
	}

	if setAbsResolution {
		err = vd.setAbsResolution()
		if err != nil {
			return fmt.Errorf("setAbsResolution: %v", err)
		}
	}

	return nil
}

func (vd *virtualDevice) setAbsResolution() error {
	eventFile, err := os.Open(vd.eventPath)
	if err != nil {
		return fmt.Errorf("failed to open event file %s: %v", vd.eventPath, err)
	}
	defer eventFile.Close()

	for _, event := range vd.config.absoluteAxes {
		if event.Resolution > 0 {
			absInfo := linux.InputAbsInfo{
				Value:      event.Value,
				Minimum:    event.Min,
				Maximum:    event.Max,
				Fuzz:       event.Fuzz,
				Flat:       event.Flat,
				Resolution: event.Resolution,
			}

			err = ioctl(eventFile, linux.EVIOCSABS(event.Axis), uintptr(unsafe.Pointer(&absInfo)))
			if err != nil {
				return fmt.Errorf("failed to set EVIOCSABS(0x%x), InputAbsInfo: %v", event.Axis, err)
			}
		}
	}
	return nil
}

func (vd *virtualDevice) registerKeys() error {
	if (vd.config.keys == nil || len(vd.config.keys) == 0) && (vd.config.buttons == nil || len(vd.config.buttons) == 0) {
		return nil
	}

	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_KEY))
	if err != nil {
		return fmt.Errorf("failed to set UI_SET_EVBIT, EV_KEY: %v", err)
	}

	if vd.config.keys != nil {
		for _, key := range vd.config.keys {
			err = ioctl(vd.fd, linux.UI_SET_KEYBIT, uintptr(key))
			if err != nil {
				return fmt.Errorf("failed to register key 0x%x: %v", key, err)
			}
		}
	}

	if vd.config.buttons != nil {
		for _, button := range vd.config.buttons {
			err = ioctl(vd.fd, linux.UI_SET_KEYBIT, uintptr(button))
			if err != nil {
				return fmt.Errorf("failed to register button 0x%x: %v", button, err)
			}
		}
	}

	if vd.config.repeat != nil {
		err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_REP))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_EVBIT, EV_REP: %v", err)
		}
	}

	return nil
}

func (vd *virtualDevice) registerScanCode() error {
	if !vd.config.scanCode {
		return nil
	}
	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_MSC))
	if err != nil {
		return fmt.Errorf("failed to set UI_SET_EVBIT, EV_MSC: %v", err)
	}

	err = ioctl(vd.fd, linux.UI_SET_MSCBIT, uintptr(linux.MSC_SCAN))
	if err != nil {
		return fmt.Errorf("failed to register MSC_SCAN: %v", err)
	}
	return nil
}

func (vd *virtualDevice) registerAxes() error {
	if vd.config.absoluteAxes != nil {
		err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_ABS))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_EVBIT, EV_ABS: %v", err)
		}
		for _, event := range vd.config.absoluteAxes {
			err = ioctl(vd.fd, linux.UI_SET_ABSBIT, uintptr(event.Axis))
			if err != nil {
				return fmt.Errorf("failed to register absolute axe 0x%x: %v", event.Axis, err)
			}
		}
	}

	if vd.config.relativeAxes != nil {
		err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_REL))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_EVBIT, EV_REL: %v", err)
		}
		for _, axis := range vd.config.relativeAxes {
			err = ioctl(vd.fd, linux.UI_SET_RELBIT, uintptr(axis))
			if err != nil {
				return fmt.Errorf("failed to register relative axe 0x%x: %v", axis, err)
			}
		}
	}
	return nil
}

func (vd *virtualDevice) registerProperties() error {
	if len(vd.config.properties) == 0 {
		return nil
	}
	for _, prop := range vd.config.properties {
		err := ioctl(vd.fd, linux.UI_SET_PROPBIT, uintptr(prop))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_PROPBIT, 0x%x: %v", prop, err)
		}
	}
	return nil
}

func (vd *virtualDevice) registerMiscEvents() error {
	if len(vd.config.miscEvents) == 0 {
		return nil
	}
	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_MSC))
	if err != nil {
		return fmt.Errorf("failed to set UI_SET_EVBIT, EV_MSC: %v", err)
	}
	for _, event := range vd.config.miscEvents {
		err := ioctl(vd.fd, linux.UI_SET_MSCBIT, uintptr(event))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_MSCBIT, 0x%x: %v", event, err)
		}
	}
	return nil
}

func (vd *virtualDevice) registerLeds() error {
	if len(vd.config.leds) == 0 {
		return nil
	}
	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_LED))
	if err != nil {
		return fmt.Errorf("failed to set UI_SET_EVBIT, EV_LED: %v", err)
	}
	for _, led := range vd.config.leds {
		err := ioctl(vd.fd, linux.UI_SET_LEDBIT, uintptr(led))
		if err != nil {
			return fmt.Errorf("failed to set UI_SET_LEDBIT, 0x%x: %v", led, err)
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

	if vd.config.repeat != nil {
		vd.Send(uint16(linux.EV_MSC), uint16(linux.REP_DELAY), vd.config.repeat.delay)
		vd.Send(uint16(linux.EV_MSC), uint16(linux.REP_PERIOD), vd.config.repeat.period)
		vd.SyncReport()
	}
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

func (vd *virtualDevice) Sync(evType linux.SyncEvent) {
	vd.Send(uint16(linux.EV_SYN), uint16(evType), 0)
}

func (vd *virtualDevice) SyncReport() {
	vd.Sync(linux.SYN_REPORT)
}

func (vd *virtualDevice) PressKey(key linux.Key) {
	vd.Send(uint16(linux.EV_KEY), uint16(key), 1)
}

func (vd *virtualDevice) ReleaseKey(key linux.Key) {
	vd.Send(uint16(linux.EV_KEY), uint16(key), 0)
}

func (vd *virtualDevice) PressButton(button linux.Button) {
	vd.Send(uint16(linux.EV_KEY), uint16(button), 1)
}

func (vd *virtualDevice) ReleaseButton(button linux.Button) {
	vd.Send(uint16(linux.EV_KEY), uint16(button), 0)
}

func (vd *virtualDevice) SendAbsoluteEvent(axis linux.AbsoluteAxis, value int32) {
	vd.Send(uint16(linux.EV_ABS), uint16(axis), value)
}

func (vd *virtualDevice) SendRelativeEvent(axis linux.RelativeAxis, value int32) {
	vd.Send(uint16(linux.EV_REL), uint16(axis), value)
}

func (vd *virtualDevice) SendMiscEvent(event linux.MiscEvent, value int32) {
	vd.Send(uint16(linux.EV_MSC), uint16(event), value)
}

func (vd *virtualDevice) SetLed(led linux.Led, state bool) {
	value := int32(0)
	if state {
		value = 1
	}
	vd.Send(uint16(linux.EV_LED), uint16(led), value)
}
