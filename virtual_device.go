package virtual_device

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/jbdemonte/virtual-device/linux"
	"os"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

type VirtualDevice interface {
	SetPath(path string) VirtualDevice
	SetMode(mode os.FileMode) VirtualDevice
	SetBusType(busType linux.BusType) VirtualDevice
	SetVendorID(vendorID uint16) VirtualDevice
	SetProductID(productID uint16) VirtualDevice
	SetVersion(version uint16) VirtualDevice
	SetName(name string) VirtualDevice
	SetEventKeys(keys []linux.Key) VirtualDevice
	SetEventButtons(buttons []linux.Button) VirtualDevice
	Register() error
	Close() error
	Send(evType, code uint16, value int32) error
	SendSync() error
	KeyPress(key int) error
	KeyDown(key int) error
	KeyUp(key int) error
}

func NewVirtualDevice() VirtualDevice {
	return &virtualDevice{
		path: "/dev/uinput",
		mode: 0660,
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

func (vd *virtualDevice) SetBusType(busType linux.BusType) VirtualDevice {
	vd.id.BusType = busType
	return vd
}

func (vd *virtualDevice) SetVendorID(vendorID uint16) VirtualDevice {
	vd.id.Vendor = vendorID
	return vd
}

func (vd *virtualDevice) SetProductID(productID uint16) VirtualDevice {
	vd.id.Product = productID
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

func (vd *virtualDevice) Register() error {
	fd, err := os.OpenFile(vd.path, syscall.O_WRONLY|syscall.O_NONBLOCK, vd.mode)
	if err != nil {
		return errors.New("could not open device file")
	}

	vd.fd = fd

	err = vd.registerEventKeysAndButtons()
	if err != nil {
		return vd.unregisterOnError(err)
	}

	return vd.createDevice()
}

func (vd *virtualDevice) createDevice() (err error) {
	var fixedSizeName [linux.UINPUT_MAX_NAME_SIZE]byte
	copy(fixedSizeName[:], vd.name)

	fmt.Println("Creating virtual device")

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, linux.UInputUserDev{
		Name: fixedSizeName,
		ID:   vd.id,
	})
	if err != nil {
		return fmt.Errorf("failed to write user device buffer: %v", err)
	}

	_, err = vd.fd.Write(buf.Bytes())
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

func (vd *virtualDevice) registerEventKeysAndButtons() error {
	if (vd.events.keys == nil || len(vd.events.keys) == 0) && (vd.events.buttons == nil || len(vd.events.buttons) == 0) {
		return nil
	}

	err := ioctl(vd.fd, linux.UI_SET_EVBIT, uintptr(linux.EV_KEY))
	if err != nil {
		return fmt.Errorf("failed to set EvBit: %v", err)
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

	return nil
}

func (vd *virtualDevice) unregisterOnError(err error) error {
	uErr := vd.Close()
	return concatErrors(err, uErr)
}

func (vd *virtualDevice) Close() (err error) {
	defer func() {
		if vd.fd != nil {
			cErr := vd.fd.Close()
			if cErr != nil {
				cErr = fmt.Errorf("failed to close the device: %v", cErr)
				err = concatErrors(err, cErr)
			}
			vd.fd = nil
		}
	}()
	err = ioctl(vd.fd, linux.UI_DEV_DESTROY, 0)
	if err != nil {
		err = fmt.Errorf("failed to unregister the device: %v", err)
	}
	if vd.cancel != nil {
		vd.cancel()
	}
	return nil
}

// Send sends an event to the device.
func (vd *virtualDevice) Send(evType, code uint16, value int32) error {
	var once sync.Once
	once.Do(func() {
		var ctxt context.Context
		ctxt, vd.cancel = context.WithCancel(context.Background())
		vd.out = make(chan linux.InputEvent, 1)
		go func() {
			defer close(vd.out)
			var event linux.InputEvent
			for {
				select {
				case event = <-vd.out:
					//buf := (*(*[1<<27 - 1]byte)(unsafe.Pointer(&event)))[:linux.SizeofEvent]

					// chat GPT (todo a verifier)
					// _, err := fd.Write((*[unsafe.Sizeof(event)]byte)(unsafe.Pointer(&event))[:])
					buf := (*[unsafe.Sizeof(event)]byte)(unsafe.Pointer(&event))[:]

					n, err := vd.fd.Write(buf)
					if err != nil {
						return
					}
					if n < linux.SizeofEvent {
						fmt.Fprintf(os.Stderr, "poll outbox: short write\n")
					}

				case <-ctxt.Done():
					break
				}
			}
		}()
	})
	vd.out <- linux.InputEvent{
		Type:  evType,
		Code:  code,
		Value: value,
	}
	return nil
}

func (vd *virtualDevice) SendSync() error {
	return vd.Send(uint16(linux.EV_SYN), uint16(linux.SYN_REPORT), 0)
}

func (vd *virtualDevice) KeyPress(key int) error {
	err := vd.KeyDown(key)
	if err != nil {
		return err
	}
	err = vd.KeyUp(key)
	if err != nil {
		return err
	}
	return nil
}

func (vd *virtualDevice) KeyDown(key int) error {
	return vd.Send(uint16(linux.EV_KEY), uint16(key), 1)
}

func (vd *virtualDevice) KeyUp(key int) error {
	return vd.Send(uint16(linux.EV_KEY), uint16(key), 0)
}
