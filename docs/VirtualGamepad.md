## VirtualGamepad Documentation

The `VirtualGamepad` interface provides functionality to simulate gamepad inputs, including button presses and analog stick movements, using a virtual device. 
The `VirtualGamepadFactory` is used to configure and create instances of `VirtualGamepad`.

---

### **VirtualGamepad**

The `VirtualGamepad` interface defines the core functionalities of a virtual gamepad.


#### **Methods**

| **Action**          | **Description**                                                                              |
|---------------------|----------------------------------------------------------------------------------------------|
| **Register**        | Registers the virtual gamepad device with the system.                                        |
| **Unregister**      | Unregisters the virtual gamepad device, releasing system resources.                          |
| **Press**           | Simulates pressing a button on the gamepad.                                                  |
| **Release**         | Simulates releasing a button on the gamepad.                                                 |
| **MoveLeftStick**   | Moves the left analog stick to the specified X and Y coordinates (values between -1 and 1).  |
| **MoveLeftStickX**  | Moves the left analog stick on the X-axis.                                                   |
| **MoveLeftStickY**  | Moves the left analog stick on the Y-axis.                                                   |
| **MoveRightStick**  | Moves the Right analog stick to the specified X and Y coordinates (values between -1 and 1). |
| **MoveRightStickX** | Moves the right analog stick on the X-axis.                                                  |
| **MoveRightStickY** | Moves the right analog stick on the Y-axis.                                                  |
| **Send**            | Sends a raw input event of the specified type, code, and value.                              |

#### **Standardized Gamepad Input Handling**

To ensure consistent behavior across different virtual gamepad implementations, this package standardizes gamepad inputs using predefined constants that align with the **Debian gamepad API** specification. These constants are derived from the [Linux Gamepad Documentation](https://www.kernel.org/doc/Documentation/input/gamepad.txt) and cover all common gamepad buttons.

![Linux Gamepad API](Linux%20Gamepad%20API.png "Linux Gamepad API")

---

##### **Button Constants**

The `gamepad.Button` type defines a set of constants representing standard gamepad buttons:

| **Constant**      | **Description**              |
|-------------------|------------------------------|
| `ButtonUp`        | D-pad up                     |
| `ButtonRight`     | D-pad right                  |
| `ButtonDown`      | D-pad down                   |
| `ButtonLeft`      | D-pad left                   |
| `ButtonNorth`     | Top face button              |
| `ButtonEast`      | Right face button            |
| `ButtonSouth`     | Bottom face button           |
| `ButtonWest`      | Left face button             |
| `ButtonL1`        | Left shoulder button         |
| `ButtonR1`        | Right shoulder button        |
| `ButtonL2`        | Left trigger button          |
| `ButtonR2`        | Right trigger button         |
| `ButtonL3`        | Left stick button (pressed)  |
| `ButtonR3`        | Right stick button (pressed) |
| `ButtonSelect`    | Select button                |
| `ButtonStart`     | Start button                 |
| `ButtonMode`      | Mode or system button        |
| `ButtonFiller1`   | Custom button 1              |
| `ButtonFiller2`   | Custom button 2              |
| `ButtonFiller3`   | Custom button 3              |
| `ButtonFiller4`   | Custom button 4              |


##### Why Standardization Matters

By standardizing the gamepad inputs:

- **Cross-Platform Consistency**: Ensures that virtual gamepads behave consistently across different systems and platforms.
- **Compatibility**: Aligns with the widely used Debian gamepad API, ensuring better support for applications that rely on standard gamepad behavior. 
- **Simplified Development**: Developers can rely on a consistent set of button and direction constants, making


---

#### **Gamepad Stick Handling**

Gamepads use **absolute axes** to represent the position of their analog sticks. Each stick has two axes: **X** (horizontal) and **Y** (vertical).
These axes usually provide raw integer values, which can vary depending on the hardware.

For example:
- A stick might return values ranging from `-32768` to `32767` (common for many controllers).
- Another device might use a different range, such as `0` to `1024`.

To simplify handling these values, this package **normalizes** the stick positions to a standard range of `-1.0` to `1.0`. This eliminates the need to handle device-specific ranges in your application.

##### **Normalization**

Normalized stick coordinates follow this convention:
- **`-1.0`**: Fully left or fully up.
- **`0.0`**: Centered.
- **`1.0`**: Fully right or fully down.

This standardization allows developers to write consistent logic for stick input, regardless of the underlying hardware.

##### **Custom Configuration**

The behavior of a virtual gamepad depends directly on its **configuration**. You can define and configure axes (`ABS_X`, `ABS_Y` or any other) with custom ranges, resolutions, and properties to suit your specific requirements.

For more detail, look the code of [NewXBox360](../gamepad/XBox360.go).

---

### **VirtualGamepadFactory**

The `VirtualGamepadFactory` is used to configure and create instances of `VirtualGamepad`. It supports method chaining for easy setup.

#### **Methods**

| **Action**         | **Description**                                                                       |
|--------------------|---------------------------------------------------------------------------------------|
| **WithDevice**     | Attaches an existing `VirtualDevice` to the gamepad.                                  |
| **WithDigital**    | Configures the digital button mappings for the gamepad.                               |
| **WithLeftStick**  | Configures the analog mappings for the left stick.                                    |
| **WithRightStick** | Configures the analog mappings for the right stick.                                   |
| **Create**         | Creates an instance of `VirtualGamepad` with the specified configuration.             |


---

### **Example Usage**

Here’s how to configure and use a custom `VirtualGamepad` (based on the predefined [XBox360](../gamepad/XBox360.go)):
```go
package main

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/gamepad"
	"github.com/jbdemonte/virtual-device/linux"
	"github.com/jbdemonte/virtual-device/sdl"
	"log"
)

func main() {
    g := gamepad.NewVirtualGamepadFactory().
            WithDevice(
                virtual_device.
                    NewVirtualDevice().
                    WithBusType(linux.BUS_USB).
                    WithVendor(sdl.USB_VENDOR_MICROSOFT).
                    WithProduct(sdl.USB_PRODUCT_XBOX360_XUSB_CONTROLLER).
                    WithVersion(0x107).
                    WithName("Xbox 360 Wireless Receiver (XBOX)"),
            ).
            WithDigital(
                gamepad.MappingDigital{
                    gamepad.ButtonSouth: linux.BTN_SOUTH,
                    gamepad.ButtonEast:  linux.BTN_EAST,
                    gamepad.ButtonNorth: linux.BTN_WEST,
                    gamepad.ButtonWest:  linux.BTN_NORTH,
    
                    gamepad.ButtonSelect: linux.BTN_SELECT,
                    gamepad.ButtonStart:  linux.BTN_START,
                    gamepad.ButtonMode:   linux.BTN_MODE,
    
                    gamepad.ButtonUp:    []gamepad.InputEvent{linux.BTN_TRIGGER_HAPPY3, gamepad.HatEvent{Axis: linux.ABS_HAT0Y, Value: -1}},
                    gamepad.ButtonDown:  []gamepad.InputEvent{linux.BTN_TRIGGER_HAPPY4, gamepad.HatEvent{Axis: linux.ABS_HAT0Y, Value: 1}},
                    gamepad.ButtonLeft:  []gamepad.InputEvent{linux.BTN_TRIGGER_HAPPY1, gamepad.HatEvent{Axis: linux.ABS_HAT0X, Value: -1}},
                    gamepad.ButtonRight: []gamepad.InputEvent{linux.BTN_TRIGGER_HAPPY2, gamepad.HatEvent{Axis: linux.ABS_HAT0X, Value: 1}},
    
                    gamepad.ButtonL1: linux.BTN_TL,
                    gamepad.ButtonR1: linux.BTN_TR,
    
                    gamepad.ButtonL2: virtual_device.AbsAxis{Axis: linux.ABS_Z, Min: 0, Value: 0, Max: 255},
                    gamepad.ButtonR2: virtual_device.AbsAxis{Axis: linux.ABS_RZ, Min: 0, Value: 0, Max: 255},
    
                    gamepad.ButtonL3: linux.BTN_THUMBL,
                    gamepad.ButtonR3: linux.BTN_THUMBR,
                },
            ).
            WithLeftStick(
                gamepad.MappingStick{
                    X: virtual_device.AbsAxis{Axis: linux.ABS_X, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
                    Y: virtual_device.AbsAxis{Axis: linux.ABS_Y, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
                },
            ).
            WithRightStick(
                gamepad.MappingStick{
                    X: virtual_device.AbsAxis{Axis: linux.ABS_RX, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
                    Y: virtual_device.AbsAxis{Axis: linux.ABS_RY, Min: -32768, Value: 0, Max: 32767, Flat: 128, Fuzz: 16},
                },
            ).
            Create()
    
    err := g.Register()
    if err != nil {
		log.Fatalf("Failed to register virtual gamepad: %v", err)
    }
    defer g.Unregister()
    
    g.Press(gamepad.ButtonSouth)
    g.Release(gamepad.ButtonSouth)
    
    g.MoveLeftStick(0.5, -0.5)
    
    g.MoveRightStick(-1.0, 1.0)

    g.Send(uint16(linux.EV_KEY), uint16(gamepad.ButtonSouth), 1)
	
}
```
This documentation outlines the essential steps for configuring, registering, and using a `VirtualGamepad` to simulate gamepad inputs and analog stick movements.