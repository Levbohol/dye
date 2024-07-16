package dye

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
