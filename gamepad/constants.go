package gamepad

// https://www.kernel.org/doc/Documentation/input/gamepad.txt

type Button int

const (
	ButtonUp Button = iota + 1
	ButtonRight
	ButtonDown
	ButtonLeft

	ButtonNorth
	ButtonEast
	ButtonSouth
	ButtonWest

	ButtonL1
	ButtonR1

	ButtonL2
	ButtonR2

	ButtonL3
	ButtonR3

	ButtonSelect
	ButtonStart
	ButtonMode

	ButtonFiller1
	ButtonFiller2
	ButtonFiller3
	ButtonFiller4
)
