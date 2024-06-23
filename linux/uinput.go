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

const UINPUT_MAX_NAME_SIZE = 80

// uinput_user_dev from https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h#L223
type UInputUserDev struct {
	Name       [UINPUT_MAX_NAME_SIZE]byte
	ID         InputID
	EffectsMax uint32
	AbsMax     [ABS_CNT]int32
	AbsMin     [ABS_CNT]int32
	AbsFuzz    [ABS_CNT]int32
	AbsFlat    [ABS_CNT]int32
}
