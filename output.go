package dye

import "fmt"

func Print(message string, color Color) {
	colorRGB := color.ConvertToRGB()
	formattedMessage := fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", colorRGB.Red, colorRGB.Green, color.ConvertToRGB().Blue, message)

	fmt.Print(formattedMessage)
}
