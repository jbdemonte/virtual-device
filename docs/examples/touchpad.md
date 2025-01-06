## **Example of use of a virtual touchpad**

### Simple touch

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/touchpad"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	tp := touchpad.NewGenericTouchpad()

	err := tp.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	time.Sleep(7_000 * time.Millisecond)

	fmt.Println("touch center")
	tp.Touch(0.5, 0.5, 0.5)
	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("move")
	tp.Touch(0.8, 0.8, 0.5)
	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("full press")
	tp.Touch(0.8, 0.8, 1)
	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("release")
	tp.Touch(0.8, 0.8, 0)
}
```

### Multi touch

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/touchpad"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	tp := touchpad.NewGenericTouchpad()

	err := tp.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	time.Sleep(7_000 * time.Millisecond)

	fmt.Println("2 fingers")
	slots := tp.MultiTouch([]touchpad.TouchSlot{
		{X: 0, Y: 0, Pressure: 0.5},
		{X: 0.2, Y: 0.2, Pressure: 0.5},
	})

	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("move #0")
	slots[0].X = 0.8
	slots[0].Pressure = 1
	slots = tp.MultiTouch(slots)

	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("add one finger")
	slots = append(slots, touchpad.TouchSlot{
		X:        0.3,
		Y:        0.4,
		Pressure: 0.2,
	})
	slots = tp.MultiTouch(slots)

	time.Sleep(2_000 * time.Millisecond)
	fmt.Println("release all fingers")
	slots[0].Pressure = 0
	slots[1].Pressure = 0
	slots[2].Pressure = 0
	slots = tp.MultiTouch(slots)
}
```