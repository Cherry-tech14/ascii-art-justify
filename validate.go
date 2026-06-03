package main

func IsValidateBanner(banner string) bool {
	switch banner {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}

func IsValidAlign(align string) bool {
	switch align {
	case "left", "right", "center", "justify":
		return true
	}
	return false
}
