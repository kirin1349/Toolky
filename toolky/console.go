package toolky

import (
	"fmt"
	"runtime"
)

func PrintInfo(content string) {
	if runtime.GOOS == "windows" {
		fmt.Println(content)
	} else {
		fmt.Printf("%c[1;40;36m %s %c[0m\n", 0x1B, content, 0x1B)
	}
}

func PrintLog(content string) {
	if runtime.GOOS == "windows" {
		fmt.Println(content)
	} else {
		fmt.Printf("%c[1;40;32m %s %c[0m\n", 0x1B, content, 0x1B)
	}
}

func PrintWarning(content string) {
	if runtime.GOOS == "windows" {
		fmt.Println(content)
	} else {
		fmt.Printf("%c[1;40;33m %s %c[0m\n", 0x1B, content, 0x1B)
	}
}

func PrintError(content string) {
	if runtime.GOOS == "windows" {
		fmt.Println(content)
	} else {
		fmt.Printf("%c[1;40;31m %s %c[0m\n", 0x1B, content, 0x1B)
	}
}