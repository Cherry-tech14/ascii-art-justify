package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	if len(args) == 1 {
		if strings.HasPrefix(args[0], "--") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		ascii := LoadBanner(args[0], "standard.txt")
		fmt.Print(ascii)
		return
	}

	if len(args) == 2 {
		if strings.HasPrefix(args[0], "--") {
			fmt.Println("Usage: go run . OPTION STRING BANNER")
			return
		}
		banner := args[1]
		if !IsValidateBanner(banner) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		ascii := LoadBanner(args[0], banner+".txt")
		fmt.Print(ascii)
		return
	}

	if len(args) == 3 {
		if !strings.HasPrefix(args[0], "--align=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		align := strings.TrimPrefix(args[0], "--align=")
		if !IsValidAlign(align) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		banner := args[2]
		if !IsValidateBanner(banner) {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			return
		}
		ascii := LoadBanner(args[1], banner+".txt")
		fmt.Print(ApplyAlignment(ascii, align, args[1], banner+".txt"))
		return
	}

	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
}
