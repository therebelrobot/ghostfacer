package xiaodisplay

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

func Rect(display ssd1306.Device, x1 int16, x2 int16, y1 int16, y2 int16, c color.RGBA) {
	for x := x1; x < x2; x++ {
		display.SetPixel(x, y1, c)
		display.SetPixel(x, y2, c)
	}
	for y := y1; y < y2; y++ {
		display.SetPixel(x1, y, c)
		display.SetPixel(x2, y, c)
	}
}

func FilledRect(display ssd1306.Device, x1 int16, x2 int16, y1 int16, y2 int16, c color.RGBA) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			display.SetPixel(x, y, c)
		}
	}
}

func DottedRect(display ssd1306.Device, x1 int16, x2 int16, y1 int16, y2 int16, c color.RGBA) {
	for x := x1; x <= x2; x += 2 {
		for y := y1; y <= y2; y += 2 {
			display.SetPixel(x, y, c)
		}
	}
}
