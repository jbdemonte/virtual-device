package keyboard

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func getKeymap() KeyMap {
	layout, _ := getKeyboardLayout()

	switch layout {
	case "fr", "azerty":
		return azertyKeyMap
	case "us", "qwerty":
		return qwertyKeyMap
	default:
		return qwertyKeyMap // Default to US QWERTY
	}
}

// getKeyboardLayout detects the keyboard layout using setxkbmap and localectl as a fallback.
func getKeyboardLayout() (string, error) {
	// Try to use setxkbmap first
	layout, err := getKeyboardLayoutWithSetxkbmap()
	if err == nil {
		return layout, nil
	}

	// If setxkbmap fails, try localectl
	layout, err = getKeyboardLayoutWithLocalectl()
	if err == nil {
		return layout, nil
	}

	// If both fail, return an error
	return "", fmt.Errorf("unable to detect keyboard layout with setxkbmap or localectl")
}

// getKeyboardLayoutWithSetxkbmap detects the layout using setxkbmap.
func getKeyboardLayoutWithSetxkbmap() (string, error) {
	cmd := exec.Command("setxkbmap", "-query")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("setxkbmap failed: %v", out.String())
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "layout:") {
			fields := strings.Fields(line)
			if len(fields) == 2 {
				return fields[1], nil
			}
		}
	}

	return "", fmt.Errorf("setxkbmap did not return a valid layout")
}

// getKeyboardLayoutWithLocalectl detects the layout using localectl.
func getKeyboardLayoutWithLocalectl() (string, error) {
	cmd := exec.Command("localectl", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("localectl failed: %v", out.String())
	}

	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "X11 Layout:") {
			fields := strings.Fields(line)
			if len(fields) == 3 {
				return fields[2], nil
			}
		}
	}

	return "", fmt.Errorf("localectl did not return a valid layout")
}
