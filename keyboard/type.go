package keyboard

import "github.com/jbdemonte/virtual-device/linux"

type Repeat struct {
	delay  int32
	period int32
}

type Config struct {
	scanCode bool
	keys     []linux.Key
	leds     []linux.Led
	repeat   *Repeat
}
