package main

import "strings"

func AlignCenter(ascii string) string {
	terminalWidth := GetTerminalWidth()
	lines := strings.Split(ascii, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		spaces := (terminalWidth - len(line)) / 2
		if spaces < 0 {
			spaces = 0
		}
		result.WriteString(strings.Repeat(" ", spaces) + line + "\n")
	}
	return result.String()
}
