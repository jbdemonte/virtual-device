package keyboard

import "github.com/jbdemonte/virtual-device/linux"

// Repeat holds the key repeat delay and period in milliseconds.
type Repeat struct {
	delay  int32
	period int32
}

// KeyMap maps runes to their key codes and required modifiers.
type KeyMap map[rune]struct {
	keyCode       linux.Key
	shiftRequired bool
	altGrRequired bool
}
