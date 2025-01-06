## **Example of use of a virtual gamepad**

All listed gamepad [factories](../../README.md#gamepads) works the same, for the purpose of the example, we'll simulate a XBOX 360 controller.

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/gamepad"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	g := gamepad.NewXBox360()

	err := g.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	fmt.Println("Register done")

	buttons := []gamepad.Button{
		gamepad.ButtonUp,
		gamepad.ButtonRight,
		gamepad.ButtonDown,
		gamepad.ButtonLeft,

		gamepad.ButtonNorth,
		gamepad.ButtonEast,
		gamepad.ButtonSouth,
		gamepad.ButtonWest,

		gamepad.ButtonL1,
		gamepad.ButtonR1,
		gamepad.ButtonL2,
		gamepad.ButtonR2,
		gamepad.ButtonL3,
		gamepad.ButtonR3,

		gamepad.ButtonSelect,
		gamepad.ButtonStart,
		gamepad.ButtonMode,
	}

	for i := 0; i < 10; i++ {

		for _, button := range buttons {
			fmt.Println("KeyDown...")
			g.Press(button)

			time.Sleep(1_000 * time.Millisecond)
			fmt.Println("KeyUp... ")
			g.Release(button)

		}

		g.MoveLeftStick(1, 1)
		time.Sleep(1_000 * time.Millisecond)
		g.MoveLeftStick(-1, -1)
		time.Sleep(1_000 * time.Millisecond)
		g.MoveLeftStick(0, 0)

		g.MoveRightStick(1, 1)
		time.Sleep(1_000 * time.Millisecond)
		g.MoveRightStick(-1, -1)
		time.Sleep(1_000 * time.Millisecond)
		g.MoveRightStick(0, 0)
	}

	err = g.Unregister()

	if err != nil {
		fmt.Printf("Failed to unregister the device: %s\n", err)
		return
	}
}
```