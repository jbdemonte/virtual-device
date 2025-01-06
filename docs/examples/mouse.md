## **Example of use of a virtual mouse**

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/mouse"
	"github.com/jbdemonte/virtual-device/linux"
	"time"
)

func main() {
	m := mouse.NewGenericMouse()

	err := m.Register()

	if err != nil {
		fmt.Printf("Failed to register the device: %s\n", err)
		return
	}

	fmt.Println("move down")
	m.MoveY(100)
	time.Sleep(1_000 * time.Millisecond)

	fmt.Println("move left")
	m.MoveX(-100)
	time.Sleep(1_000 * time.Millisecond)

	fmt.Println("move up")
	m.MoveY(-100)
	time.Sleep(1_000 * time.Millisecond)

	fmt.Println("move right")
	m.MoveX(100)
	time.Sleep(1_000 * time.Millisecond)

	fmt.Println("double click left")
	m.DoubleClickLeft()
	time.Sleep(1_000 * time.Millisecond)

	fmt.Println("right click")
	m.ClickRight()
	time.Sleep(1_000 * time.Millisecond)

	m.ScrollUp()
}
```