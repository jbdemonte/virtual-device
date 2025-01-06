## **Example of use of a virtual keyboard**

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/keyboard"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	kb := keyboard.NewLogitechG510()

	err := kb.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	// Turns on the Num Lock LED
	kb.SetLed(linux.LED_NUML, true)

	// Presses the key mapped to 'A' based on the keyboard layout (qwerty, azerty, etc.)
	kb.PressKey(linux.KEY_A)
	kb.SyncReport()
	time.Sleep(200 * time.Millisecond)
	kb.ReleaseKey(linux.KEY_A)
	kb.SyncReport()

	// Press then release the 'Z' based on the keyboard layout (qwerty, azerty, etc.)
	kb.TapKey(linux.KEY_Z)

	// Attempts to detect the keyboard layout and sends the appropriate key events accordingly
	kb.Type("\n\n")
	kb.Type("Hello World!")
	kb.Type("\n\n")
	kb.Type("MASTER, I AM HERE TO SERVE YOU.\n123.67.88.1")
	kb.Type("\n\n")
}
```