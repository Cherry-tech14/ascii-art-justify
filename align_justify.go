package main

import "strings"

func AlignJustify(ascii string) string {
	terminalWidth := GetTerminalWidth()
	lines := strings.Split(ascii, "\n")
	var result strings.Builder

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		words := strings.Fields(line)
		if len(words) <= 1 {
			result.WriteString(line + "\n")
			continue
		}
		totalChars := 0
		for _, word := range words {
			totalChars += len(word)
		}
		totalSpaces := terminalWidth - totalChars
		gaps := len(words) - 1
		spacePerGap := totalSpaces / gaps
		extraSpaces := totalSpaces % gaps

		var justified strings.Builder
		for i, word := range words {
			justified.WriteString(word)
			if i < gaps {
				space := spacePerGap
				if i < extraSpaces {
					space++
				}
				justified.WriteString(strings.Repeat(" ", space))
			}
		}
		result.WriteString(justified.String() + "\n")
	}
	return result.String()
}
