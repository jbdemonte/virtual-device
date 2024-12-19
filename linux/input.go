package linux

import (
	"syscall"
	"unsafe"
)

// From https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h

/*
 * The event structure itself
 */
type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

const SizeofEvent = int(unsafe.Sizeof(InputEvent{}))

type BusType uint16

// input_id from https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h#L59
type InputID struct {
	BusType BusType
	Vendor  uint16
	Product uint16
	Version uint16
}

const (
	BUS_PCI       BusType = 0x01
	BUS_ISAPNP    BusType = 0x02
	BUS_USB       BusType = 0x03
	BUS_HIL       BusType = 0x04
	BUS_BLUETOOTH BusType = 0x05
	BUS_VIRTUAL   BusType = 0x06

	BUS_ISA         BusType = 0x10
	BUS_I8042       BusType = 0x11
	BUS_XTKBD       BusType = 0x12
	BUS_RS232       BusType = 0x13
	BUS_GAMEPORT    BusType = 0x14
	BUS_PARPORT     BusType = 0x15
	BUS_AMIGA       BusType = 0x16
	BUS_ADB         BusType = 0x17
	BUS_I2C         BusType = 0x18
	BUS_HOST        BusType = 0x19
	BUS_GSC         BusType = 0x1A
	BUS_ATARI       BusType = 0x1B
	BUS_SPI         BusType = 0x1C
	BUS_RMI         BusType = 0x1D
	BUS_CEC         BusType = 0x1E
	BUS_INTEL_ISHTP BusType = 0x1F
	BUS_AMD_SFH     BusType = 0x20
)

// input_absinfo from https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h#L93
type InputAbsInfo struct {
	Value      int32
	Minimum    int32
	Maximum    int32
	Fuzz       int32
	Flat       int32
	Resolution int32
}

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h#L180

func EVIOCSABS(abs AbsoluteAxis) uintptr {
	return _IOW('E', uintptr(0xc0+abs), InputAbsInfo{})
}
