package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	if !strings.HasPrefix(os.Args[1], "--align=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	align := strings.TrimPrefix(os.Args[1], "--align=")
	if !IsValidAlign(align) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	banner := os.Args[3]
	if !IsValidateBanner(banner) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}
	ascii := LoadBanner(os.Args[2], banner+".txt")
	fmt.Print(ApplyAlignment(ascii, align))
}
