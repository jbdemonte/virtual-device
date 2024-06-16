package linux

// From https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h

/*
 * Event types
 */

type EventType uint16

const (
	EV_SYN       EventType = 0x00
	EV_KEY       EventType = 0x01
	EV_REL       EventType = 0x02
	EV_ABS       EventType = 0x03
	EV_MSC       EventType = 0x04
	EV_SW        EventType = 0x05
	EV_LED       EventType = 0x11
	EV_SND       EventType = 0x12
	EV_REP       EventType = 0x14
	EV_FF        EventType = 0x15
	EV_PWR       EventType = 0x16
	EV_FF_STATUS EventType = 0x17
	EV_MAX       EventType = 0x1f
	EV_CNT       EventType = (EV_MAX + 1)
)
