package virtual_device

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"virtual-input/linux"
)

// original function taken from: https://github.com/tianon/debian-golang-pty/blob/master/ioctl.go
func ioctl[A int | uint16 | linux.EventType](deviceFile *os.File, cmd uintptr, arg A) error {
	_, _, errorCode := syscall.Syscall(syscall.SYS_IOCTL, deviceFile.Fd(), cmd, uintptr(arg))
	if errorCode != 0 {
		return errorCode
	}
	return nil
}

func concatErrors(errs ...error) error {
	var errorMessages []string
	for _, err := range errs {
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}

	if len(errorMessages) == 0 {
		return nil
	}

	return errors.New(strings.Join(errorMessages, "; "))
}
