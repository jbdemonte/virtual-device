# Virtual Device

This package provides pure go functions to create virtual devices on Linux using [input](https://www.kernel.org/doc/html/latest/input/index.html) and [uinput](https://www.kernel.org/doc/html/latest/input/uinput.html) through [ioctl](https://en.wikipedia.org/wiki/Ioctl).

## Installation

```sh
$ go get -u github.com/jbdemonte/virtual-device
```

## **Documentation - Quick access**

- Base Class: [VirtualDevice](./docs/VirtualDevice.md)
- Helper Class: [VirtualKeyboard](./docs/VirtualKeyboard.md)
- Helper Class: [VirtualMouse](./docs/VirtualMouse.md)
- Helper Class: [VirtualTouchpad](./docs/VirtualTouchpad.md)
- Helper Class: [VirtualGamepad](./docs/VirtualGamepad.md)
- Method Used: [How to create a new virtual device profile](./docs/Creation.md)

## Permission Issues

Reading events from input devices or creating virtual `uinput` devices requires `$USER` to have the appropriate system-level permissions. 
This can be accomplished by adding `$USER` to a group with read/write access to `/dev/input/event*` and `uinput` block devices.

Example on Debian based OS:
```sh
sudo usermod -aG input $USER
```

## Overview

The **`virtual-device`** package provides a flexible framework for creating and managing virtual input devices on Linux using the uinput interface. It is designed to simplify the process of creating virtual devices like keyboards, mice, and gamepads, while offering a clean and extensible API.

### **Core Architecture**

At the heart of the package is the **`VirtualDevice`** base class. This class provides the fundamental functionality required to create, configure, and manage virtual devices. It serves as the foundation upon which the helper classes are built.

#### Base Class: [`VirtualDevice`](./docs/VirtualDevice.md)

The **`VirtualDevice`** class encapsulates the low-level interactions with uinput, including:
- Device creation and configuration (e.g., setting event types and capabilities).
- Sending events such as key presses, mouse movements, and button clicks.
- Handling synchronization (`EV_SYN`) to ensure event sequences are properly reported.

This class is generic and can be used directly for custom virtual devices, but it requires detailed knowledge of the underlying input subsystem.

#### Helper Classes

To simplify common use cases, the package provides **helper classes** built on top of **`VirtualDevice`**, each tailored for specific device types:

1. **[`VirtualKeyboard`](./docs/VirtualKeyboard.md)**:
   - A high-level interface for creating and managing virtual keyboards.
   - Includes helper methods for sending key presses, key releases, and full key strokes.
   - Example: `keyboard.Type("Hello World!")`.

2. **[`VirtualMouse`](./docs/VirtualMouse.md)**:
   - Designed for creating virtual mice or pointing devices.
   - Provides methods for moving the cursor, scrolling, and simulating mouse button actions.
   - Example: `mouse.ClickLeft()` or `mouse.Move(0.5, 1)`.

3. **[`VirtualGamepad`](./docs/VirtualGamepad.md)**:
   - Tailored for creating virtual game controllers.
   - Supports axis movements, button presses, and handling force feedback effects.
   - Example: `gamepad.Press(gamepad.ButtonUp)` or `gamepad.MoveLeftStick(0.5, 1)`.

4. **[`VirtualTouchpad`](./docs/VirtualTouchpad.md)**:
   - A high-level interface for creating and managing virtual touchpads.
   - Includes methods for simulating multitouch gestures, individual finger movements, and tap actions.
   - Supports both Protocol A and Protocol B for multitouch devices.
   - Example: `	slots := tp.MultiTouch([]touchpad.TouchSlot{{X: 0, Y: 0, Pressure: 0.5}, {X: 0.2, Y: 0.2, Pressure: 0.5} })`.


#### Pre-Configured Virtual Device Factories

This package provides pre-configured factory functions to create virtual devices that simulate specific devices. These functions simplify the creation of virtual devices by providing ready-made configurations for popular hardware such as Sony PS5, Nintendo Switch Pro controllers, and more.

##### Keyboard

- **`NewGenericKeyboard`**  
  Creates a virtual generic keyboard.

- **`NewLogitechG510`**  
  Creates a virtual keyboard with the layout and features of a Logitech G510 gaming keyboard

##### Mouse

- **`NewGenericMouse`**  
  Creates a virtual generic mouse with basic movement, scrolling, and button support.

- **`NewLogitechG402`**  
  Creates a virtual mouse with the layout and features of a Logitech G402 gaming mouse.

##### Touchpad

- **`NewGenericTouchpad`**  
  Creates a virtual generic touchpad with basic multitouch, button support, and absolute axis handling.

##### Gamepads

- **`NewSonyPS4`**  
  Creates a virtual controller with the layout and behavior of a Sony PS4 DualShock controller.

- **`NewSonyPS5`**  
  Creates a virtual controller with the layout and behavior of a Sony PS5 DualSense controller.

- **`NewSwitchPro`**  
  Creates a virtual controller with the layout and behavior of a Nintendo Switch Pro controller.

- **`NewXBox360`**  
  Creates a virtual controller with the layout and behavior of an Xbox 360 controller.

- **`NewXBoxOneS`**  
  Creates a virtual controller with the layout and behavior of an Xbox One S controller.

- **`NewStadia`**  
  Creates a virtual controller with the layout and behavior of a Google Stadia controller.

- **`NewSN30Pro`**  
  Creates a virtual controller with the layout and behavior of an 8BitDo SN30 Pro controller.


##### **Advantages of Using Pre-Configured Factories**
**Ease of Use** 
No need to manually configure each aspect of the virtual device. These functions provide pre-set mappings and layouts.

**Accuracy**  
Each factory function is designed to closely match the behavior and layout of the real-world hardware.

**Flexibility** 
Use the created devices as starting points and customize them further if needed.

**Rapid Prototyping**  
Quickly test and simulate hardware without requiring physical devices.


#### **Linux Input Constants**

Linux input constants (e.g., keys, buttons, axes, LEDs) used in this package are defined in the `linux` package. These constants are directly mapped from the Linux kernel's [`input.h`](https://github.com/torvalds/linux/blob/master/include/uapi/linux/input.h), [`input-event-codes.h`](https://github.com/torvalds/linux/blob/master/include/uapi/linux/input-event-codes.h) and [`uinput.h`](https://github.com/torvalds/linux/blob/master/include/uapi/linux/uinput.h) files and provide a comprehensive set of identifiers for input events.

**Examples of Constants**

```go
linux.KEY_A  // Represents the "A" key
linux.KEY_ENTER  // Represents the "Enter" key

linux.BTN_LEFT  // Represents the left mouse button
linux.BTN_RIGHT  // Represents the right mouse button

linux.ABS_X  // Represents the X-axis for absolute positioning
linux.REL_Y  // Represents the Y-axis for relative positioning

linux.LED_CAPSL  // Represents the Caps Lock LED
linux.LED_NUML  // Represents the Num Lock LED
```

## Credits
Package freely inspired by [kenshaw/evdev](https://github.com/kenshaw/evdev), [bendahl/uinput](https://github.com/bendahl/uinput) and some others.
