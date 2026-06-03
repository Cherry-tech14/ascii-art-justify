package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	out, err := exec.Command("tput", "cols").Output()
	if err != nil {
		return 80
	}
	width, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 80
	}
	return width
}
