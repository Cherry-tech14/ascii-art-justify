package main

func ApplyAlignment(ascii, align string) string {
	switch align {
	case "left":
		return ascii
	case "right":
		return AlignRight(ascii)
	case "center":
		return AlignCenter(ascii)
	case "justify":
		return AlignJustify(ascii)
	}
	return ascii
}
