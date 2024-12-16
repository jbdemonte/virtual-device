package gamepad

import (
	"fmt"
	virtual_device "github.com/jbdemonte/virtual-device"
	"github.com/jbdemonte/virtual-device/linux"
)

func convertHatToAbsAxis(hatEvents []HatEvent) []virtual_device.AbsAxis {
	result := make([]virtual_device.AbsAxis, 0)
	hatToAbsEvents := map[linux.AbsoluteAxis][]int32{}

	for _, hatEvent := range hatEvents {
		_, exists := hatToAbsEvents[hatEvent.Axis]
		if !exists {
			hatToAbsEvents[hatEvent.Axis] = make([]int32, 0)
		}
		hatToAbsEvents[hatEvent.Axis] = append(hatToAbsEvents[hatEvent.Axis], hatEvent.Value)
	}
	for axis, values := range hatToAbsEvents {
		minValue, maxValue, err := minMax(values)

		if err != nil {
			fmt.Printf("Error converting hat to absolute axis: %s\n", err)
			continue
		}

		e := virtual_device.AbsAxis{
			Axis:  axis,
			Value: minValue + (maxValue-minValue)/2,
			Min:   minValue,
			Max:   maxValue,
		}

		result = append(result, e)
	}

	return result
}

func minMax(values []int32) (int32, int32, error) {
	if len(values) == 0 {
		return 0, 0, fmt.Errorf("MinMax failed, source array is empty")
	}

	minValue, maxValue := values[0], values[0]

	for _, value := range values {
		if value < minValue {
			minValue = value
		}
		if value > maxValue {
			maxValue = value
		}
	}
	return minValue, maxValue, nil
}
