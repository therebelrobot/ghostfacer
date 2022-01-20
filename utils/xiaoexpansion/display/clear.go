package xiaodisplay

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

func Clear(display ssd1306.Device) {
	FilledRect(display, 0, 128, 0, 64, color.RGBA{0, 0, 0, 255})
}
