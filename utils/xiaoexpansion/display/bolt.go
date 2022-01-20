package xiaodisplay

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

func Bolt(display ssd1306.Device, x int16, y int16, c color.RGBA) {

	/*
		_____X_
		____X__
		___XX__
		__X_X__
		_X__X__
		XX_-_XX
		__X__X_
		__X_X__
		__XX___
		__X____
		_X_____

	*/
	display.SetPixel(x+2, y-5, c)

	display.SetPixel(x+1, y-4, c)

	display.SetPixel(x, y-3, c)
	display.SetPixel(x+1, y-3, c)

	display.SetPixel(x-1, y-2, c)
	display.SetPixel(x+1, y-2, c)

	display.SetPixel(x-2, y-1, c)
	display.SetPixel(x+1, y-1, c)

	display.SetPixel(x-3, y, c)
	display.SetPixel(x-2, y, c)
	display.SetPixel(x+2, y, c)
	display.SetPixel(x+3, y, c)

	display.SetPixel(x-1, y+1, c)
	display.SetPixel(x+2, y+1, c)

	display.SetPixel(x-1, y+2, c)
	display.SetPixel(x+1, y+2, c)

	display.SetPixel(x-1, y+3, c)
	display.SetPixel(x, y+3, c)

	display.SetPixel(x-1, y+4, c)

	display.SetPixel(x-2, y+5, c)
}

func FilledBolt(display ssd1306.Device, x int16, y int16, c color.RGBA) {

	/*
		_____X_
		____X__
		___XX__
		__XXX__
		_XXXX__
		XXX0XXX
		__XXXX_
		__XXX__
		__XX___
		__X____
		_X_____

	*/
	display.SetPixel(x+2, y-5, c)

	display.SetPixel(x+1, y-4, c)

	display.SetPixel(x, y-3, c)
	display.SetPixel(x+1, y-3, c)

	display.SetPixel(x-1, y-2, c)
	display.SetPixel(x, y-2, c)
	display.SetPixel(x+1, y-2, c)

	display.SetPixel(x-2, y-1, c)
	display.SetPixel(x-1, y-1, c)
	display.SetPixel(x, y-1, c)
	display.SetPixel(x+1, y-1, c)

	display.SetPixel(x-3, y, c)
	display.SetPixel(x-2, y, c)
	display.SetPixel(x-1, y, c)
	display.SetPixel(x, y, c)
	display.SetPixel(x+1, y, c)
	display.SetPixel(x+2, y, c)
	display.SetPixel(x+3, y, c)

	display.SetPixel(x-1, y+1, c)
	display.SetPixel(x, y+1, c)
	display.SetPixel(x+1, y+1, c)
	display.SetPixel(x+2, y+1, c)

	display.SetPixel(x-1, y+2, c)
	display.SetPixel(x, y+2, c)
	display.SetPixel(x+1, y+2, c)

	display.SetPixel(x-1, y+3, c)
	display.SetPixel(x, y+3, c)

	display.SetPixel(x-1, y+4, c)

	display.SetPixel(x-2, y+5, c)
}
