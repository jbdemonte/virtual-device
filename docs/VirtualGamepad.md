## VirtualGamepad Documentation

The `VirtualGamepad` interface provides functionality to simulate gamepad inputs, including button presses and analog stick movements, using a virtual device. 
The `VirtualGamepadFactory` is used to configure and create instances of `VirtualGamepad`.

---

### **VirtualGamepad**

The `VirtualGamepad` interface defines the core functionalities of a virtual gamepad.


#### **Methods**

| **Action**         | **Description**                                                                                      |
|--------------------|------------------------------------------------------------------------------------------------------|
| **Register**       | Registers the virtual gamepad device with the system.                                               |
| **Unregister**     | Unregisters the virtual gamepad device, releasing system resources.                                 |
| **Press**          | Simulates pressing a button on the gamepad.                                                         |
| **Release**        | Simulates releasing a button on the gamepad.                                                        |
| **MoveLeftStick**  | Moves the left analog stick to the specified X and Y coordinates (values between -1 and 1).         |
| **MoveLeftStickX** | Moves the left analog stick on the X-axis.                                                          |
| **MoveLeftStickY** | Moves the left analog stick on the Y-axis.                                                          |
| **MoveRightStick** | Moves the Right analog stick to the specified X and Y coordinates (values between -1 and 1).        |
| **MoveRightStickX**| Moves the right analog stick on the X-axis.                                                         |
| **MoveRightStickY**| Moves the right analog stick on the Y-axis.                                                         |


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

Here’s how to configure and use a `VirtualGamepad`:

```go
gamepad := virtual_device.NewVirtualGamepadFactory().
    WithDigital(MappingDigital{
        Buttons: MappingButtons{
            ButtonA: linux.BTN_SOUTH,
            ButtonB: linux.BTN_EAST,
            ButtonX: linux.BTN_WEST,
            ButtonY: linux.BTN_NORTH,
        },
    }).
    WithLeftStick(MappingStick{
        X: linux.ABS_X,
        Y: linux.ABS_Y,
    }).
    WithRightStick(MappingStick{
        X: linux.ABS_RX,
        Y: linux.ABS_RY,
    }).
	Create()

err := gamepad.Register()
if err != nil {
    log.Fatalf("Failed to register virtual gamepad: %v", err)
}
defer gamepad.Unregister()

gamepad.Press(ButtonA)
gamepad.Release(ButtonA)

gamepad.MoveLeftStick(0.5, -0.5)

gamepad.MoveRightStick(-1.0, 1.0)
```

This documentation outlines the essential steps for configuring, registering, and using a `VirtualGamepad` to simulate gamepad inputs and analog stick movements.