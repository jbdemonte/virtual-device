package linux

// FROM https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h

const (
	/* ioctl */
	UI_DEV_CREATE  = 0x5501
	UI_DEV_DESTROY = 0x5502
)

const (
	UI_SET_EVBIT  = 0x40045564
	UI_SET_KEYBIT = 0x40045565
	UI_SET_MSCBIT = 0x40045566
)
