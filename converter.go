package dye

import (
	"strconv"
	"strings"
)

type RGB struct {
	Red   int
	Green int
	Blue  int
}

func (color Color) ConvertToRGB() RGB {
	return RGB{
		Red:   int((color >> 16) & 0xFF),
		Green: int((color >> 8) & 0xFF),
		Blue:  int(color & 0xFF),
	}
}

func ColorFromHTML(color string) Color {
	color = strings.TrimPrefix(color, "#")

	hex, err := strconv.ParseUint(color, 16, 32)
	if err != nil {
		return Color(0xFFFFFF)
	}

	return Color(hex)
}
