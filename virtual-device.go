package virtual_device

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"virtual-input/linux"
)

type VirtualDevice interface {
	SetPath(path string) VirtualDevice
	SetMode(mode os.FileMode) VirtualDevice
	SetVendorID(vendorID uint16) VirtualDevice
	SetProductID(productID uint16) VirtualDevice
	SetBusType(busType linux.BusType) VirtualDevice
	SetEventKeys(keys []uint16) VirtualDevice
	Register() error
}

func NewVirtualInput() VirtualDevice {
	return &virtualDevice{
		path: "/dev/uinput",
		mode: 0660,
	}
}

func (vd virtualDevice) SetPath(path string) VirtualDevice {
	vd.path = path
	return vd
}

func (vd virtualDevice) SetMode(mode os.FileMode) VirtualDevice {
	vd.mode = mode
	return vd
}

func (vd virtualDevice) SetVendorID(vendorID uint16) VirtualDevice {
	vd.id.vendor = vendorID
	return vd
}

func (vd virtualDevice) SetProductID(productID uint16) VirtualDevice {
	vd.id.product = productID
	return vd
}

func (vd virtualDevice) SetBusType(busType linux.BusType) VirtualDevice {
	vd.id.busType = busType
	return vd
}

func (vd virtualDevice) SetEventKeys(keys []uint16) VirtualDevice {
	vd.events.keys = keys
	return vd
}

func (vd virtualDevice) Register() error {
	fd, err := os.OpenFile(vd.path, syscall.O_WRONLY|syscall.O_NONBLOCK, vd.mode)
	if err != nil {
		return errors.New("could not open device file")
	}

	vd.fd = fd

	err = vd.registerEventKeys()
	if err != nil {
		return vd.unregisterOnError(err)
	}

	return nil
}

func (vd virtualDevice) registerEventKeys() error {
	if vd.events.keys == nil || len(vd.events.keys) == 0 {
		return nil
	}

	err := ioctl(vd.fd, linux.UI_SET_EVBIT, linux.EV_KEY)
	if err != nil {
		return fmt.Errorf("failed to set evBit%v", err)
	}

	for _, key := range vd.events.keys {

		err = ioctl(vd.fd, linux.UI_SET_KEYBIT, key)
		if err != nil {
			return fmt.Errorf("failed to register key %d: %v", key, err)
		}
	}

	return nil
}

func (vd virtualDevice) unregisterOnError(err error) error {
	uErr := vd.Unregister()
	return concatErrors(err, uErr)
}

func (vd virtualDevice) Unregister() (err error) {
	defer func() {
		cErr := vd.fd.Close()
		if cErr != nil {
			cErr = fmt.Errorf("failed to close the device: %v", cErr)
			err = concatErrors(err, cErr)
		}
	}()
	err = ioctl(vd.fd, linux.UI_DEV_DESTROY, 0)
	if err != nil {
		err = fmt.Errorf("failed to unregister the device: %v", err)
	}
	return nil
}
