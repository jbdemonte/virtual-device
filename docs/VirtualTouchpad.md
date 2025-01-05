## VirtualTouchpad Documentation

The `VirtualTouchpad` interface provides functionality to simulate touch inputs, multitouch gestures, and button actions using a virtual device. 
The `VirtualTouchpadFactory` is used to configure and create instances of `VirtualTouchpad`.

---

### **VirtualTouchpad**

The `VirtualTouchpad` interface defines the core functionalities of a virtual touchpad.

It supports 3 normes (see [multi-touch-protocol](https://www.kernel.org/doc/Documentation/input/multi-touch-protocol.txt)) : 
- the legacy single touch protocol by using the `Touch`
- the multitouch Protocol A by activating it with `WithLegacyMultitouch` using the `Multitoutch` function.
- the multitouch Protocol B sing the `Multitoutch` function.

_Coordinates X and Y are between -1 and 1 (normalized)_  
_Pressure is between 0 and 1_

#### **Methods**

| **Action**           | **Description**                                                                                                   |
|----------------------|-------------------------------------------------------------------------------------------------------------------|
| **Register**         | Registers the virtual touchpad device with the system.                                                            |
| **Unregister**       | Unregisters the virtual touchpad device, releasing system resources.                                              |
| **Touch**            | Simulates a single touch event at the specified coordinates with pressure.                                        |
| **MultiTouch**       | Simulates multitouch events using multiple touch slots. Returns the updated touch slots after the gesture.<br>_The slot IDs are automatically managed, so reuse the previous result on the next call._ |
| **PressButton**      | Simulates pressing a touchpad button.                                                                             |
| **ReleaseButton**    | Simulates releasing a touchpad button.                                                                            |
| **Click**            | Simulates a single click of the specified button.                                                                 |
| **DoubleClick**      | Simulates a double click of the specified button.                                                                 |
| **ClickLeft**        | Convenience method to simulate a single left click.                                                               |
| **ClickRight**       | Convenience method to simulate a single right click.                                                              |
| **DoubleClickLeft**  | Convenience method to simulate a double left click.                                                               |
| **DoubleClickRight** | Convenience method to simulate a double right click.                                                             |


#### **Touchpad Handling**

Touchpads use **absolute axes** to represent the position of touch points. Each touch point has two main axes: **X** (horizontal) and **Y** (vertical). 
Additionally, touchpads often provide a **pressure** value that indicates how firmly the user is pressing on the surface. 
The ranges for these values can vary depending on the hardware.

To simplify handling these variations, this package **normalizes** touch coordinates and pressure values to standard ranges. 
This ensures consistent behavior across different touchpad devices.

##### **Normalization**

Normalized values follow these conventions:
- **X and Y Axes**:
  - **`-1.0`**: Fully left or fully up.
  - **`0.0`**: Centered.
  - **`1.0`**: Fully right or fully down.

- **Pressure**:
  - **`0.0`**: No pressure (touch lifted or minimal contact).
  - **`1.0`**: Maximum pressure (fully pressed).

This standardization allows you to work with consistent touch input, regardless of the underlying device.

##### **Custom Configuration**

The behavior of a virtual touchpad depends directly on its **configuration**. You can define and configure axes (`ABS_X`, `ABS_Y` or `ABS_PRESSURE`) with custom ranges, resolutions, and properties to suit your specific requirements.

For more detail, look the code of [NewGenericTouchpad](../touchpad/GenericTouchpad.go).

---

### **VirtualTouchpadFactory**

The `VirtualTouchpadFactory` is used to configure and create instances of `VirtualTouchpad`. It supports method chaining for easy setup.

#### **Methods**

| **Action**               | **Description**                                                                                 |
|--------------------------|-------------------------------------------------------------------------------------------------|
| **WithDevice**           | Attaches an existing `VirtualDevice` to the touchpad.                                           |
| **WithClickDelay**       | Sets the delay between press and release for a single click.                                    |
| **WithDoubleClickDelay** | Sets the delay between two clicks for a double click.                                           |
| **WithAxes**             | Configures the absolute axes supported by the touchpad (e.g., X, Y coordinates).                |
| **WithButtons**          | Configures the buttons supported by the touchpad.                                              |
| **WithProperties**       | Configures the properties of the touchpad (e.g., multitouch support).                           |
| **WithLegacyMultitouch** | Enables legacy multitouch protocol (Protocol A).                                                |
| **Create**               | Creates an instance of `VirtualTouchpad` with the specified configuration.                      |


---

### **Example Usage**

Here’s how to configure and use a `VirtualTouchpad`:

```go
package main

import (
  "fmt"
  virtual_device "github.com/jbdemonte/virtual-device"
  "github.com/jbdemonte/virtual-device/linux"
  "github.com/jbdemonte/virtual-device/touchpad"
  "log"
  "time"
)

func main() {
    tp := touchpad.NewVirtualTouchpadFactory().
            WithDevice(
              virtual_device.NewVirtualDevice().
                WithBusType(linux.BUS_USB).
                WithVendor(0x02).
                WithProduct(0x07).
                WithVersion(0x01).
                WithName("Synaptics TouchPad"),
            ).
            WithAxes([]virtual_device.AbsAxis{
              {Axis: linux.ABS_X, Min: 1472, Value: 1472, Max: 5472, Resolution: 40},
              {Axis: linux.ABS_Y, Min: 1408, Value: 1408, Max: 4448, Resolution: 40},
              {Axis: linux.ABS_PRESSURE, Min: 0, Value: 0, Max: 255, IsUnidirectional: true},
              {Axis: linux.ABS_MT_SLOT, Min: 0, Value: 0, Max: 4},
              {Axis: linux.ABS_MT_POSITION_X, Min: 1472, Value: 0, Max: 5472, Resolution: 40},
              {Axis: linux.ABS_MT_POSITION_Y, Min: 1408, Value: 1408, Max: 4448, Resolution: 40},
              {Axis: linux.ABS_MT_TRACKING_ID, Min: 0, Value: 0, Max: 65535},
            }).
            WithButtons([]linux.Button{
              linux.BTN_LEFT,
              linux.BTN_RIGHT,
              linux.BTN_TOOL_FINGER,
              linux.BTN_TOUCH,
              linux.BTN_TOOL_DOUBLETAP,
              linux.BTN_TOOL_TRIPLETAP,
            }).
            WithProperties([]linux.InputProp{
              linux.INPUT_PROP_POINTER, linux.INPUT_PROP_BUTTONPAD,
            }).
      Create() 
    
    err := tp.Register()
    if err != nil {
      log.Fatalf("Failed to register virtual touchpad: %v", err)
    }
    defer tp.Unregister()
    
    
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
        {X: -0.3, Y:-0.4, Pressure: 0.2},
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

This documentation outlines the essential steps for configuring, registering, and using a `VirtualTouchpad` to simulate touch inputs and gestures.