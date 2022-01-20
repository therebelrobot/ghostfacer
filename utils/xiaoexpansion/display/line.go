package xiaodisplay

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

func Line(display ssd1306.Device, start_x int32, end_x int32, start_y int32, end_y int32, color color.RGBA) {

	// Bresenham's
	var cx int32 = start_x
	var cy int32 = start_y

	var dx int32 = end_x - cx
	var dy int32 = end_y - cy
	if dx < 0 {
		dx = 0 - dx
	}
	if dy < 0 {
		dy = 0 - dy
	}

	var sx int32
	var sy int32
	if cx < end_x {
		sx = 1
	} else {
		sx = -1
	}
	if cy < end_y {
		sy = 1
	} else {
		sy = -1
	}
	var err int32 = dx - dy

	var n int32
	for n = 0; n < 1000; n++ {
		display.SetPixel(int16(cx), int16(cy), color)
		if (cx == end_x) && (cy == end_y) {
			return
		}
		var e2 int32 = 2 * err
		if e2 > (0 - dy) {
			err = err - dy
			cx = cx + sx
		}
		if e2 < dx {
			err = err + dx
			cy = cy + sy
		}
	}
}
