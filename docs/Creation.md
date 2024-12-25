## How to Create a New Virtual Device Profile

The goal is to simulate the behavior of a real device as accurately as possible. To achieve this, follow these steps using a real device connected to an Ubuntu machine.

### 1. Identify and Select the Target Device

- Plug the device to your linux system
- Run the `evtest` tool and select the connected device.

Example: Select device `22` to target:

```shell
/dev/input/event22:	    Microsoft X-Box One S pad
```

### 2. Configure the virtual device

The first screen in `evtest` displays the configuration of the selected device. Use this information to set up the virtual device programmatically.

Example output:

```shell
Select the device event number [0-22]: 22
Input driver version is 1.0.1
Input device ID: bus 0x3 vendor 0x45e product 0x2ea version 0x408
Input device name: "Microsoft X-Box One S pad"
Supported events:
  Event type 0 (EV_SYN)
  Event type 1 (EV_KEY)
    Event code 304 (BTN_SOUTH)
    Event code 305 (BTN_EAST)
    Event code 307 (BTN_NORTH)
    Event code 308 (BTN_WEST)
    Event code 310 (BTN_TL)
    Event code 311 (BTN_TR)
    Event code 314 (BTN_SELECT)
    Event code 315 (BTN_START)
    Event code 316 (BTN_MODE)
    Event code 317 (BTN_THUMBL)
    Event code 318 (BTN_THUMBR)
  Event type 3 (EV_ABS)
    Event code 0 (ABS_X)
      Value      0
      Min   -32768
      Max    32767
      Fuzz      16
      Flat     128
    Event code 1 (ABS_Y)
      Value      0
      Min   -32768
      Max    32767
      Fuzz      16
      Flat     128
    Event code 2 (ABS_Z)
      Value      0
      Min        0
      Max     1023
    Event code 3 (ABS_RX)
      Value      0
      Min   -32768
      Max    32767
      Fuzz      16
      Flat     128
    Event code 4 (ABS_RY)
      Value      0
      Min   -32768
      Max    32767
      Fuzz      16
      Flat     128
    Event code 5 (ABS_RZ)
      Value      0
      Min        0
      Max     1023
    Event code 16 (ABS_HAT0X)
      Value      0
      Min       -1
      Max        1
    Event code 17 (ABS_HAT0Y)
      Value      0
      Min       -1
      Max        1
  Event type 21 (EV_FF)
    Event code 80 (FF_RUMBLE)
    Event code 81 (FF_PERIODIC)
    Event code 88 (FF_SQUARE)
    Event code 89 (FF_TRIANGLE)
    Event code 90 (FF_SINE)
    Event code 96 (FF_GAIN)
Properties:
Testing ... (interrupt to exit)
```
Use this output to configure your virtual device to mimic the original device as closely as possible.  
Map supported events, axes, and properties accordingly.


### 3. Record Event Behavior

While still in `evtest`, interact with the real device (press buttons, move axes, etc.) to record the events it generates.  
These events will define the behavior of your virtual device.

Example event log:
```shell
Event: time 1735001653.107368, -------------- SYN_REPORT ------------
Event: time 1735001653.235435, type 1 (EV_KEY), code 308 (BTN_WEST), value 0
Event: time 1735001653.235435, -------------- SYN_REPORT ------------
Event: time 1735001656.656167, type 1 (EV_KEY), code 305 (BTN_EAST), value 1
Event: time 1735001656.656167, -------------- SYN_REPORT ------------
Event: time 1735001656.768192, type 1 (EV_KEY), code 305 (BTN_EAST), value 0
Event: time 1735001656.768192, -------------- SYN_REPORT ------------
Event: time 1735001658.824641, type 1 (EV_KEY), code 304 (BTN_SOUTH), value 1
Event: time 1735001658.824641, -------------- SYN_REPORT ------------
Event: time 1735001658.872416, type 1 (EV_KEY), code 304 (BTN_SOUTH), value 0
Event: time 1735001658.872416, -------------- SYN_REPORT ------------
Event: time 1735001660.184935, type 1 (EV_KEY), code 304 (BTN_SOUTH), value 1
Event: time 1735001660.184935, -------------- SYN_REPORT ------------
Event: time 1735001660.324974, type 1 (EV_KEY), code 304 (BTN_SOUTH), value 0
Event: time 1735001660.324974, -------------- SYN_REPORT ------------
Event: time 1735001662.549393, type 3 (EV_ABS), code 17 (ABS_HAT0Y), value -1
Event: time 1735001662.549393, -------------- SYN_REPORT ------------
Event: time 1735001662.693485, type 3 (EV_ABS), code 17 (ABS_HAT0Y), value 0
Event: time 1735001662.693485, -------------- SYN_REPORT ------------
Event: time 1735001663.217388, type 3 (EV_ABS), code 17 (ABS_HAT0Y), value 1
Event: time 1735001663.217388, -------------- SYN_REPORT ------------
Event: time 1735001663.397637, type 3 (EV_ABS), code 17 (ABS_HAT0Y), value 0
Event: time 1735001663.397637, -------------- SYN_REPORT ------------
Event: time 1735001663.913752, type 3 (EV_ABS), code 16 (ABS_HAT0X), value 1
Event: time 1735001663.913752, -------------- SYN_REPORT ------------
Event: time 1735001664.033770, type 3 (EV_ABS), code 16 (ABS_HAT0X), value 0
Event: time 1735001664.033770, -------------- SYN_REPORT ------------
Event: time 1735001664.517828, type 3 (EV_ABS), code 16 (ABS_HAT0X), value -1
Event: time 1735001664.517828, -------------- SYN_REPORT ------------
Event: time 1735001664.637858, type 3 (EV_ABS), code 16 (ABS_HAT0X), value 0
Event: time 1735001664.637858, -------------- SYN_REPORT ------------
Event: time 1735001664.921977, type 3 (EV_ABS), code 3 (ABS_RX), value 200
Event: time 1735001664.921977, -------------- SYN_REPORT ------------
Event: time 1735001664.929952, type 3 (EV_ABS), code 3 (ABS_RX), value 349
Event: time 1735001664.929952, -------------- SYN_REPORT ------------
Event: time 1735001664.937974, type 3 (EV_ABS), code 3 (ABS_RX), value 497
Event: time 1735001700.810793, -------------- SYN_REPORT ------------
Event: time 1735001700.818790, type 3 (EV_ABS), code 0 (ABS_X), value 30734
Event: time 1735001700.818790, type 3 (EV_ABS), code 1 (ABS_Y), value 15468
Event: time 1735001700.818790, -------------- SYN_REPORT ------------
Event: time 1735001700.826785, type 3 (EV_ABS), code 0 (ABS_X), value 30925
Event: time 1735001700.826785, type 3 (EV_ABS), code 1 (ABS_Y), value 14386
Event: time 1735001700.826785, -------------- SYN_REPORT ------------
Event: time 1735001700.834780, type 3 (EV_ABS), code 1 (ABS_Y), value 13289
Event: time 1735001700.834780, -------------- SYN_REPORT ------------
Event: time 1735001700.842777, type 3 (EV_ABS), code 0 (ABS_X), value 31044
Event: time 1735001700.842777, type 3 (EV_ABS), code 1 (ABS_Y), value 12281
Event: time 1735001700.842777, -------------- SYN_REPORT ------------
Event: time 1735001700.850771, type 3 (EV_ABS), code 0 (ABS_X), value 30826
Event: time 1735001700.850771, type 3 (EV_ABS), code 1 (ABS_Y), value 11049
Event: time 1735001700.850771, -------------- SYN_REPORT ------------
```

In some cases, `evtest` logs might not display certain events, such as `ABS_MT_SLOT`. 
This is because the Linux kernel optimizes event reporting to avoid unnecessary messages.   

For example, if a multitouch slot remains unchanged, the kernel may skip logging it to reduce overhead. 
Keep this in mind when analyzing event logs.



### 3. Implement the Virtual Device

Using the data collected in steps 2 and 3, configure your virtual device with the appropriate axes, buttons, and properties.
This ensures that the virtual device closely mimics the original device's functionality.

```go
device := virtual_device.NewVirtualDevice().
    WithBusType(linux.BUS_USB).
    WithVendor(0x45e).
    WithProduct(0x2ea).
    WithVersion(0x408).
    WithName("Microsoft X-Box One S pad").
    WithKeys([]linux.Key{
        linux.BTN_SOUTH, linux.BTN_EAST, linux.BTN_NORTH, linux.BTN_WEST,
        linux.BTN_TL, linux.BTN_TR, linux.BTN_SELECT, linux.BTN_START,
    }).
    WithAbsAxes([]virtual_device.AbsAxis{
        {Axis: linux.ABS_X, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
        {Axis: linux.ABS_Y, Min: -32768, Max: 32767, Fuzz: 16, Flat: 128},
    }).
    WithProperties([]linux.InputProp{linux.INPUT_PROP_BUTTONPAD})

err := device.Register()
if err != nil {
    log.Fatalf("Failed to register virtual device: %v", err)
}
defer device.Unregister()
```

## Notes
These steps ensure that your virtual device behaves as closely as possible to the original hardware.  
While this guide uses `evtest` for input collection, other tools like `libevdev` can also provide detailed information about device behavior.

## Debugging
You're free to manage the code as you see fit. However, when working with the binary, you can monitor the `ioctl` commands and their results using the following command:
```shell
strace -e ioctl  ./main
```