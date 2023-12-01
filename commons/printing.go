package commons

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintErr(msg string) {
	red := color.New(color.FgRed, color.Bold)
	red.Printf(msg)
	red.Println()
}

func PrintInfo(msg string) {
	green := color.New(color.FgGreen, color.Bold)

	green.Printf(msg)
	green.Println()
}

func PrintInfoFormat(msgFormat string, a ...any ) {
	PrintInfo(fmt.Sprintf(msgFormat, a...))
}

func PrintDebug(msg string) {
	fmt.Println(msg)
} 

func PrintDebugFormat(msgFormat string, a ...any ) {
	PrintDebug(fmt.Sprintf(msgFormat, a...))
}