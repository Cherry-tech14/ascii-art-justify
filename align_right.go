package main

import (
	"strings"
)

func AlignRight(ascii string) string {
	terminalWidth := getTerminalWidth()
	lines := strings.Split(ascii, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		spaces := terminalWidth - len(line)
		if spaces < 0 {
			spaces = 0
		}
		result.WriteString(strings.Repeat(" ", spaces) + line + "\n")
	}
	return result.String()
}

