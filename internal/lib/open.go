package lib

import (
	"fmt"
	"os/exec"
	"runtime"
)

func checkCommandExsists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}

func GetOpenCommand() (string, error) {
	switch runtime.GOOS {
	case "darwin": // macOS
		if checkCommandExsists("open") {
			return "open", nil
		}
	case "windows":
		if checkCommandExsists("explorer.exe") {
			return "explorer.exe", nil
		}
	case "linux":
		if checkCommandExsists("xdg-open") {
			return "xdg-open", nil
		}
	}
	return "", fmt.Errorf("no suitable open command found")
}
