package keyboard

import (
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
)

func NewLogitechG510() VirtualKeyboard {
	vd := virtual_device.NewVirtualDevice()

	vd.
		SetBusType(linux.BUS_USB).
		SetProduct(0xc22d).
		SetVendor(sdl.USB_VENDOR_LOGITECH).
		SetVersion(0x111).
		SetName("Logitech G510 Gaming Keyboard")

	config := Config{
		scanCode: true,
		repeat: &Repeat{
			delay:  250,
			period: 33,
		},
		leds: []linux.Led{
			linux.LED_NUML,
			linux.LED_CAPSL,
			linux.LED_SCROLLL,
			linux.LED_COMPOSE,
			linux.LED_KANA,
		},
		keys: []linux.Key{
			linux.KEY_ESC,
			linux.KEY_1, linux.KEY_2, linux.KEY_3, linux.KEY_4, linux.KEY_5, linux.KEY_6, linux.KEY_7, linux.KEY_8, linux.KEY_9, linux.KEY_0,
			linux.KEY_MINUS, linux.KEY_EQUAL, linux.KEY_BACKSPACE,

			linux.KEY_TAB,
			linux.KEY_Q, linux.KEY_W, linux.KEY_E, linux.KEY_R, linux.KEY_T, linux.KEY_Y, linux.KEY_U, linux.KEY_I, linux.KEY_O, linux.KEY_P,
			linux.KEY_LEFTBRACE, linux.KEY_RIGHTBRACE, linux.KEY_ENTER,

			linux.KEY_LEFTCTRL,
			linux.KEY_A, linux.KEY_S, linux.KEY_D, linux.KEY_F, linux.KEY_G, linux.KEY_H, linux.KEY_J, linux.KEY_K, linux.KEY_L,
			linux.KEY_SEMICOLON, linux.KEY_APOSTROPHE, linux.KEY_GRAVE,

			linux.KEY_LEFTSHIFT,
			linux.KEY_BACKSLASH, linux.KEY_Z, linux.KEY_X, linux.KEY_C, linux.KEY_V, linux.KEY_B, linux.KEY_N, linux.KEY_M,
			linux.KEY_COMMA, linux.KEY_DOT, linux.KEY_SLASH, linux.KEY_RIGHTSHIFT,

			linux.KEY_KPASTERISK,
			linux.KEY_LEFTALT, linux.KEY_SPACE, linux.KEY_RIGHTALT,

			linux.KEY_CAPSLOCK, linux.KEY_LEFTCTRL, linux.KEY_RIGHTCTRL,
			linux.KEY_LEFTMETA, linux.KEY_RIGHTMETA, linux.KEY_COMPOSE,

			linux.KEY_F1, linux.KEY_F2, linux.KEY_F3, linux.KEY_F4, linux.KEY_F5, linux.KEY_F6, linux.KEY_F7, linux.KEY_F8,
			linux.KEY_F9, linux.KEY_F10, linux.KEY_F11, linux.KEY_F12, linux.KEY_F13, linux.KEY_F14, linux.KEY_F15,
			linux.KEY_F16, linux.KEY_F17, linux.KEY_F18, linux.KEY_F19, linux.KEY_F20, linux.KEY_F21, linux.KEY_F22,
			linux.KEY_F23, linux.KEY_F24,

			linux.KEY_NUMLOCK, linux.KEY_KP0, linux.KEY_KP1, linux.KEY_KP2, linux.KEY_KP3, linux.KEY_KP4,
			linux.KEY_KP5, linux.KEY_KP6, linux.KEY_KP7, linux.KEY_KP8, linux.KEY_KP9,
			linux.KEY_KPMINUS, linux.KEY_KPPLUS, linux.KEY_KPEQUAL, linux.KEY_KPDOT, linux.KEY_KPSLASH,
			linux.KEY_KPASTERISK, linux.KEY_KPLEFTPAREN, linux.KEY_KPRIGHTPAREN, linux.KEY_KPCOMMA, linux.KEY_KPJPCOMMA,

			linux.KEY_HOME, linux.KEY_UP, linux.KEY_PAGEUP, linux.KEY_LEFT, linux.KEY_RIGHT,
			linux.KEY_END, linux.KEY_DOWN, linux.KEY_PAGEDOWN, linux.KEY_INSERT, linux.KEY_DELETE,
			linux.KEY_PAUSE, linux.KEY_SYSRQ, linux.KEY_SCROLLLOCK,

			linux.KEY_MUTE, linux.KEY_VOLUMEDOWN, linux.KEY_VOLUMEUP,
			linux.KEY_POWER, linux.KEY_SLEEP, linux.KEY_EJECTCD, linux.KEY_NEXTSONG, linux.KEY_PLAYPAUSE,
			linux.KEY_PREVIOUSSONG, linux.KEY_STOPCD,

			linux.KEY_ZENKAKUHANKAKU, linux.KEY_102ND, linux.KEY_RO,
			linux.KEY_KATAKANA, linux.KEY_HIRAGANA, linux.KEY_HENKAN,
			linux.KEY_KATAKANAHIRAGANA, linux.KEY_MUHENKAN, linux.KEY_HANGUEL, linux.KEY_HANJA, linux.KEY_YEN,

			linux.KEY_WWW, linux.KEY_BACK, linux.KEY_FORWARD, linux.KEY_REFRESH,
			linux.KEY_STOP, linux.KEY_EDIT, linux.KEY_CALC,

			linux.KEY_AGAIN, linux.KEY_PROPS, linux.KEY_UNDO, linux.KEY_FRONT,
			linux.KEY_COPY, linux.KEY_OPEN, linux.KEY_PASTE, linux.KEY_FIND,
			linux.KEY_CUT, linux.KEY_HELP, linux.KEY_SCREENLOCK, linux.KEY_UNKNOWN,
		},
	}

	return createVirtualKeyboard(vd, config)
}
