package main

import "strings"

func AlignJustify(ascii string, originalText string, banner string) string {
	terminalWidth := getTerminalWidth()

	words := strings.Fields(originalText)

	if len(words) <= 1 {
		return AlignCenter(ascii)
	}

	var wordBlocks [][]string
	for _, word := range words {
		wordArt := LoadBanner(word, banner)
		rows := strings.Split(strings.TrimRight(wordArt, "\n"), "\n")
		wordBlocks = append(wordBlocks, rows)
	}

	totalChars := 0
	for _, block := range wordBlocks {
		totalChars += len(block[0])
	}

	gaps := len(words) - 1
	totalSpaces := terminalWidth - totalChars
	spacePerGap := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var result strings.Builder
	for row := 0; row < 8; row++ {
		for i, block := range wordBlocks {
			if row < len(block) {
				result.WriteString(block[row])
			}
			if i < gaps {
				space := spacePerGap
				if i < extraSpaces {
					space++
				}
				result.WriteString(strings.Repeat(" ", space))
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}
