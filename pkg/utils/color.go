package utils

import (
	"fmt"
)

func Colorize(colorize bool, color, text string) string {
	if colorize {
		return fmt.Sprintf("%%{F%s}%s%%{F-}", color, text)
	}
	return text
}

func Fontify(colorize bool, font, text string) string {
	if colorize {
		return fmt.Sprintf("%%{T%s}%s%%{T-}", font, text)
	}
	return text
}
