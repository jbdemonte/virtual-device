package linux

import (
	"syscall"
	"unsafe"
)

// From https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h

// struct input_event
type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

const SizeofEvent = int(unsafe.Sizeof(InputEvent{}))

// Protocol version
const EV_VERSION = 0x010001

// struct input_id
type InputID struct {
	BusType BusType
	Vendor  uint16
	Product uint16
	Version uint16
}

// struct input_absinfo
type InputAbsInfo struct {
	Value      int32
	Minimum    int32
	Maximum    int32
	Fuzz       int32
	Flat       int32
	Resolution int32
}

// struct input_keymap_entry
type InputKeymapEntry struct {
	Flags    uint8
	Len      uint8
	Index    uint16
	Keycode  uint32
	Scancode [32]uint8
}

// struct input_mask
type InputMask struct {
	Type      uint32
	CodesSize uint32
	CodesPtr  uint64
}

/* get driver version */
// _IOR('E', 0x01, int)
const EVIOCGVERSION = 0x80044501

/* get device ID */
// _IOR('E', 0x02, struct input_id)
const EVIOCGID = 0x80084502

/* get repeat settings */
// _IOR('E', 0x03, unsigned int[2])
const EVIOCGREP = 0x80084503

/* set repeat settings */
// _IOW('E', 0x03, unsigned int[2])
const EVIOCSREP = 0x40084503

/* get keycode */
// _IOR('E', 0x04, unsigned int[2])
const EVIOCGKEYCODE = 0x80084504

// _IOR('E', 0x04, struct input_keymap_entry)
const EVIOCGKEYCODE_V2 = 0x80284504

// /* set keycode */
// _IOW('E', 0x04, unsigned int[2])
const EVIOCSKEYCODE = 0x40084504

// _IOW('E', 0x04, struct input_keymap_entry)
const EVIOCSKEYCODE_V2 = 0x40284504

/* get device name */
func EVIOCGNAME(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x06, uintptr(len))
}

/* get physical location */
func EVIOCGPHYS(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x07, uintptr(len))
}

/* get unique identifier */
func EVIOCGUNIQ(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x08, uintptr(len))
}

/* get device properties */
func EVIOCGPROP(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x09, uintptr(len))
}

func EVIOCGMTSLOTS(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x0a, uintptr(len))
}

/* get global key state */
func EVIOCGKEY(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x18, uintptr(len))
}

/* get all LEDs */
func EVIOCGLED(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x19, uintptr(len))
}

/* get all sounds status */
func EVIOCGSND(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x1a, uintptr(len))
}

/* get all switch states */
func EVIOCGSW(len int) uintptr {
	return _IOC(_IOC_READ, 'E', 0x1b, uintptr(len))
}

/* get event bits */
func EVIOCGBIT(ev, len int) uintptr {
	return _IOC(_IOC_READ, 'E', uintptr(0x20+ev), uintptr(len))
}

/* get abs value/limits */
func EVIOCGABS(abs AbsoluteAxis) uintptr {
	return _IOR('E', uintptr(0x40+abs), InputAbsInfo{})
}

/* set abs value/limits */
func EVIOCSABS(abs AbsoluteAxis) uintptr {
	return _IOW('E', uintptr(0xc0+abs), InputAbsInfo{})
}

/* send a force effect to a force feedback device */
func EVIOCSFF() uintptr {
	return _IOW('E', 0x80, FFEffect{})
}

/* Erase a force effect */
func EVIOCRMFF() uintptr {
	var i int32
	return _IOW('E', 0x81, i)
}

func EVIOCGEFFECTS() uintptr {
	var i int32
	return _IOR('E', 0x84, i) /* Report number of effects playable at the same time */
}

/* Grab/Release device */
func EVIOCGRAB() uintptr {
	var i int32
	return _IOW('E', 0x90, i)
}

/* Revoke device access */
func EVIOCREVOKE() uintptr {
	var i int32
	return _IOW('E', 0x91, i)
}

/* Get event-masks */
func EVIOCGMASK() uintptr {
	return _IOR('E', 0x92, InputMask{})
}

/* Set event-masks */
func EVIOCSMASK() uintptr {
	return _IOW('E', 0x93, InputMask{})
}

/* Set clockid to be used for timestamps */
func EVIOCSCLOCKID() uintptr {
	var i int32
	return _IOW('E', 0xa0, i)
}

const (
	ID_BUS     = 0
	ID_VENDOR  = 1
	ID_PRODUCT = 2
	ID_VERSION = 3
)

type BusType uint16

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

type MtToolType uint16

const (
	// MT_TOOL types
	MT_TOOL_FINGER MtToolType = 0x00
	MT_TOOL_PEN    MtToolType = 0x01
	MT_TOOL_PALM   MtToolType = 0x02
	MT_TOOL_DIAL   MtToolType = 0x0a
	MT_TOOL_MAX    MtToolType = 0x0f
)

type FFStatusType uint16

const (
	// Force feedback statuses
	FF_STATUS_STOPPED FFStatusType = 0x00
	FF_STATUS_PLAYING FFStatusType = 0x01
	FF_STATUS_MAX     FFStatusType = 0x01
)

// struct ff_replay
type FFReplay struct {
	Length uint16
	Delay  uint16
}

// struct ff_trigger
type FFTrigger struct {
	Button   uint16
	Interval uint16
}

// struct ff_envelope
type FFEnvelope struct {
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

// struct ff_constant_effect
type FFConstantEffect struct {
	Level    int16
	Envelope FFEnvelope
}

// struct  ff_ramp_effect
type FFRampEffect struct {
	StartLevel int16
	EndLevel   int16
	Envelope   FFEnvelope
}

// struct ff_condition_effect
type FFConditionEffect struct {
	RightSaturation uint16
	LeftSaturation  uint16
	RightCoeff      int16
	LeftCoeff       int16
	Deadband        uint16
	Center          int16
}

// struct ff_periodic_effect
type FFPeriodicEffect struct {
	Waveform   uint16
	Period     uint16
	Magnitude  int16
	Offset     int16
	Phase      uint16
	Envelope   FFEnvelope
	CustomLen  uint32
	CustomData *int16 // Pointer to user data; in Go may need handling differently
}

// struct ff_rumble_effect
type FFRumbleEffect struct {
	StrongMagnitude uint16
	WeakMagnitude   uint16
}

// struct ff_effect
// In C, the union U overlaps all members in the same memory.
// We emulate this with a byte array sized to the largest member ([2]FFConditionEffect = 24 bytes).
type FFEffect struct {
	Type      uint16
	ID        int16
	Direction uint16
	Trigger   FFTrigger
	Replay    FFReplay
	U         [24]byte
}

type FFEffectType uint16

const (
	// FF effect types
	FF_RUMBLE   FFEffectType = 0x50
	FF_PERIODIC FFEffectType = 0x51
	FF_CONSTANT FFEffectType = 0x52
	FF_SPRING   FFEffectType = 0x53
	FF_FRICTION FFEffectType = 0x54
	FF_DAMPER   FFEffectType = 0x55
	FF_INERTIA  FFEffectType = 0x56
	FF_RAMP     FFEffectType = 0x57

	FF_EFFECT_MIN FFEffectType = FF_RUMBLE
	FF_EFFECT_MAX FFEffectType = FF_RAMP
)

type FFPeriodicEffectType uint16

const (

	// FF periodic effect types
	FF_SQUARE   FFPeriodicEffectType = 0x58
	FF_TRIANGLE FFPeriodicEffectType = 0x59
	FF_SINE     FFPeriodicEffectType = 0x5a
	FF_SAW_UP   FFPeriodicEffectType = 0x5b
	FF_SAW_DOWN FFPeriodicEffectType = 0x5c
	FF_CUSTOM   FFPeriodicEffectType = 0x5d

	FF_WAVEFORM_MIN FFPeriodicEffectType = FF_SQUARE
	FF_WAVEFORM_MAX FFPeriodicEffectType = FF_CUSTOM

	FF_GAIN       = 0x60
	FF_AUTOCENTER = 0x61

	FF_MAX_EFFECTS = FF_GAIN

	FF_MAX = 0x7f
	FF_CNT = (FF_MAX + 1)
)
