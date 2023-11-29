package commons

import "github.com/fatih/color"

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
