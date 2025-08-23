## VirtualMouse Documentation

The `VirtualMouse` interface provides an easy way to simulate mouse inputs, including movement, scrolling, and button actions, using a virtual device. 
The `VirtualMouseFactory` is used to configure and create instances of `VirtualMouse`.

---

### **VirtualMouse**

The `VirtualMouse` interface defines the core functionalities of a virtual mouse.

#### **Methods**

| **Action**           | **Description**                                                                                   |
|----------------------|---------------------------------------------------------------------------------------------------|
| **Register**         | Registers the virtual mouse device with the system.                                              |
| **Unregister**       | Unregisters the virtual mouse device, releasing system resources.                                |
| **Move**             | Simulates moving the mouse cursor by a relative amount in both X and Y directions.               |
| **MoveX**            | Simulates moving the mouse cursor by a relative amount in the X direction.                       |
| **MoveY**            | Simulates moving the mouse cursor by a relative amount in the Y direction.                       |
| **ScrollVertical**   | Simulates vertical scrolling. Positive values scroll up; negative values scroll down.            |
| **ScrollHorizontal** | Simulates horizontal scrolling. Positive values scroll right; negative values scroll left.       |
| **ButtonPress**      | Simulates pressing a mouse button.                                                               |
| **ButtonRelease**    | Simulates releasing a mouse button.                                                              |
| **ScrollUp**         | Convenience method to simulate a vertical scroll up.                                             |
| **ScrollDown**       | Convenience method to simulate a vertical scroll down.                                           |
| **ScrollLeft**       | Convenience method to simulate a horizontal scroll left.                                         |
| **ScrollRight**      | Convenience method to simulate a horizontal scroll right.                                        |
| **Click**            | Simulates a single click of the specified button.                                                |
| **DoubleClick**      | Simulates a double click of the specified button.                                                |
| **ClickLeft**        | Convenience method to simulate a single left click.                                              |
| **ClickRight**       | Convenience method to simulate a single right click.                                             |
| **ClickMiddle**      | Convenience method to simulate a single middle click.                                            |
| **DoubleClickRight** | Convenience method to simulate a double left click.                                              |
| **DoubleClickRight** | Convenience method to simulate a double right click.                                             |
| **DoubleClickMiddle**| Convenience method to simulate a double middle click.                                            |
| **Send**             | Sends a raw input event of the specified type, code, and value.                                  |


##### **Custom Configuration**

The behavior of a virtual mouse depends directly on its **configuration**. You can define and configure axes (`REL_X`, `REL_Y` or any other) with custom ranges, resolutions, and properties to suit your specific requirements.

For more detail, look the code of [NewLogitechG402](../mouse/LogitechG402.go).

---

### **VirtualMouseFactory**

The `VirtualMouseFactory` is used to configure and create instances of `VirtualMouse`. 
It supports method chaining for easy setup.

#### **Methods**

| **Action**                    | **Description**                                                               |
|-------------------------------|-------------------------------------------------------------------------------|
| **WithDevice**                | Attaches an existing `VirtualDevice` to the mouse.                             |
| **WithClickDelay**            | Sets the delay between press and release for a single click.                   |
| **WithDoubleClickDelay**      | Sets the delay between two clicks for a double click.                          |
| **WithHighResStepVertical**   | Configures the step size for high-resolution vertical scrolling.              |
| **WithHighResStepHorizontal** | Configures the step size for high-resolution horizontal scrolling.         |
| **Create**                    | Creates an instance of `VirtualMouse` with the specified configuration.        |


---

### **Example Usage**

Here’s how to configure and use a `VirtualMouse`:

```go
package main

import (
	"fmt"
	"github.com/jbdemonte/virtual-device/mouse"
	"log"
)

func main() {
	m := mouse.NewVirtualMouseFactory().
		WithClickDelay(50).
		WithDoubleClickDelay(250).
		WithHighResStepVertical(120).
		WithHighResStepHorizontal(120).
		Create()

	err := m.Register()
	if err != nil {
		log.Fatalf("Failed to register virtual mouse: %v", err)
	}
	defer m.Unregister()

	m.Move(100, -50) // Move 100 units right and 50 units up

	// Simulate clicks
	m.ClickLeft()        // Single left click
	m.DoubleClickRight() // Double right click

	// Simulate scrolling
	m.ScrollUp()
	m.ScrollHorizontal(-1) // Scroll left
}
```

This documentation outlines the essential steps for configuring, registering, and using a `VirtualMouse` to simulate mouse inputs and control related features.
