package dye

import "fmt"

func printWithColor(message string, color Color, newLine bool) {
	colorRGB := color.ConvertToRGB()
	formattedMessage := fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", colorRGB.Red, colorRGB.Green, colorRGB.Blue, message)

	if newLine {
		formattedMessage += "\n"
	}

	fmt.Print(formattedMessage)
}

func Print(message string, color Color) {
	printWithColor(message, color, false)
}

func Println(message string, color Color) {
	printWithColor(message, color, true)
}
