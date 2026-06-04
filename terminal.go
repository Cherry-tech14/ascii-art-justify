package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80
	}
	parts := strings.Fields(string(out))
	if len(parts) >= 2 {
		width, err := strconv.Atoi(parts[1])
		if err == nil {
			return width
		}
	}
	return 80
}
