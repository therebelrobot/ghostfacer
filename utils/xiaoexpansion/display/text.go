package xiaodisplay

import (
	"image/color"

	utils "github.com/therebelrobot/ghostfacer/utils/shared"

	"tinygo.org/x/tinyfont/freesans"

	"tinygo.org/x/drivers/ssd1306"
)

func Text(display ssd1306.Device, x int16, y int16, text string) {
	c := color.RGBA{100, 100, 100, 255}
	utils.Log("Text: " + text)

	WriteLine(display, &freesans.Regular9pt7b, x, y, text, c)

}

func Heading(display ssd1306.Device, x int16, y int16, text string) {
	c := color.RGBA{100, 100, 100, 255}
	utils.Log("Heading: " + text)

	WriteLine(display, &freesans.Bold12pt7b, x, y, text, c)

}

func TextRotate(display ssd1306.Device, x int16, y int16, text string, rotation Rotation) {
	c := color.RGBA{100, 100, 100, 255}
	utils.Log("Text: " + text)

	WriteLineRotated(display, &freesans.Regular9pt7b, x, y, text, c, rotation)

}

func HeadingRotate(display ssd1306.Device, x int16, y int16, text string, rotation Rotation) {
	c := color.RGBA{100, 100, 100, 255}
	utils.Log("Heading: " + text)

	WriteLineRotated(display, &freesans.Bold12pt7b, x, y, text, c, rotation)

}
