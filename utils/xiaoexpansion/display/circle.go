package xiaodisplay

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

func Circle(display ssd1306.Device, x0 int16, y0 int16, r int16, c color.RGBA) {
	x, y, dx, dy := r-int16(1), int16(0), int16(1), int16(1)
	err := dx - (r * 2)

	for x >= y {
		display.SetPixel(x0+x, y0+y, c)
		display.SetPixel(x0+y, y0+x, c)
		display.SetPixel(x0-y, y0+x, c)
		display.SetPixel(x0-x, y0+y, c)
		display.SetPixel(x0-x, y0-y, c)
		display.SetPixel(x0-y, y0-x, c)
		display.SetPixel(x0+y, y0-x, c)
		display.SetPixel(x0+x, y0-y, c)

		if err <= int16(0) {
			y++
			err += dy
			dy += int16(2)
		}
		if err > int16(0) {
			x--
			dx += int16(2)
			err += dx - (r * int16(2))
		}
	}
}
func FilledCircle(display ssd1306.Device, x0 int16, y0 int16, r int16, c color.RGBA) {
	for rd := r; rd >= 0; rd-- {
		Circle(display, x0, y0, rd, c)
	}
}
