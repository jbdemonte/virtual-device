package linux

// FROM https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h

const (
	/* ioctl */
	UI_DEV_CREATE  = 0x5501
	UI_DEV_DESTROY = 0x5502
)

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h#L63
const UINPUT_IOCTL_BASE = 'U'
const UINPUT_IOCTL_BASE_NUMERIC = 0x40045500 // _IOW(UINPUT_IOCTL_BASE, 0, int)

// https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h#L137
const (
	UI_SET_EVBIT   = UINPUT_IOCTL_BASE_NUMERIC + 100
	UI_SET_KEYBIT  = UINPUT_IOCTL_BASE_NUMERIC + 101
	UI_SET_RELBIT  = UINPUT_IOCTL_BASE_NUMERIC + 102
	UI_SET_ABSBIT  = UINPUT_IOCTL_BASE_NUMERIC + 103
	UI_SET_MSCBIT  = UINPUT_IOCTL_BASE_NUMERIC + 104
	UI_SET_LEDBIT  = UINPUT_IOCTL_BASE_NUMERIC + 105
	UI_SET_SNDBIT  = UINPUT_IOCTL_BASE_NUMERIC + 106
	UI_SET_FFBIT   = UINPUT_IOCTL_BASE_NUMERIC + 107
	UI_SET_PHYS    = UINPUT_IOCTL_BASE_NUMERIC + 108
	UI_SET_SWBIT   = UINPUT_IOCTL_BASE_NUMERIC + 109
	UI_SET_PROPBIT = UINPUT_IOCTL_BASE_NUMERIC + 110
)

const UINPUT_MAX_NAME_SIZE = 80

// UInputUserDev is the uinput_user_dev struct from linux/uinput.h.
type UInputUserDev struct {
	Name       [UINPUT_MAX_NAME_SIZE]byte
	ID         InputID
	EffectsMax uint32
	AbsMax     [ABS_CNT]int32
	AbsMin     [ABS_CNT]int32
	AbsFuzz    [ABS_CNT]int32
	AbsFlat    [ABS_CNT]int32
}

// UInputSetup is the uinput_setup struct used with UI_DEV_SETUP ioctl.
type UInputSetup struct {
	ID           InputID
	Name         [80]byte
	FFEffectsMax uint32
}

// UI_GET_SYSNAME returns the ioctl request code to get the sysfs device name.
func UI_GET_SYSNAME(len int) uintptr {
	return _IOC(_IOC_READ, UINPUT_IOCTL_BASE, 44, uintptr(len))
}
