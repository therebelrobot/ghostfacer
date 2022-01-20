package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"image/color"
	"strconv"
	"time"

	sharedutils "github.com/therebelrobot/ghostfacer/utils/shared"
	tea5767 "github.com/therebelrobot/ghostfacer/utils/tea5767"

	display "github.com/therebelrobot/ghostfacer/utils/xiaoexpansion/display"
)

type FmScan struct {
	currentFrame   int
	goingUp        bool
	numberOfCycles int
	highLvl        map[int64][]float64
	totalHighLevel int
	maxLevel       int64
}

func (c *Ctrl) FmScanLoop() {
	sharedutils.Log("Starting FmScanLoop")
	isInitializing := c.FmScan.numberOfCycles <= 4
	var finalFreqs []float64
	if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel-2]; ok {
		c.FmScan.totalHighLevel = len(c.FmScan.highLvl[c.FmScan.maxLevel]) + len(c.FmScan.highLvl[c.FmScan.maxLevel-1]) + len(c.FmScan.highLvl[c.FmScan.maxLevel-2])
	} else if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel-1]; ok {
		c.FmScan.totalHighLevel = len(c.FmScan.highLvl[c.FmScan.maxLevel]) + len(c.FmScan.highLvl[c.FmScan.maxLevel-1])
	} else if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel]; ok {
		c.FmScan.totalHighLevel = len(c.FmScan.highLvl[c.FmScan.maxLevel])
	}
	finalFreqs = make([]float64, c.FmScan.totalHighLevel)
	if isInitializing {
		if c.FmScan.goingUp && c.Fm.currentFreq <= c.Devices.radio.Constants.FreqRangeUs[1] {
			c.Fm.currentFreq += 0.1
		} else if c.FmScan.goingUp && c.Fm.currentFreq > c.Devices.radio.Constants.FreqRangeUs[1] {
			c.Fm.currentFreq = c.Devices.radio.Constants.FreqRangeUs[1]
			c.FmScan.goingUp = false
			c.FmScan.numberOfCycles += 1
		} else if !c.FmScan.goingUp && c.Fm.currentFreq >= c.Devices.radio.Constants.FreqRangeUs[0] {
			c.Fm.currentFreq -= 0.3
		} else {
			c.FmScan.numberOfCycles += 1
			c.Fm.currentFreq = c.Devices.radio.Constants.FreqRangeUs[0]
			c.FmScan.goingUp = true
		}
	} else {
		if len(finalFreqs) != c.FmScan.totalHighLevel || finalFreqs[0] != c.FmScan.highLvl[c.FmScan.maxLevel][0] {
			if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel-2]; ok {
				finalFreqs = append(c.FmScan.highLvl[c.FmScan.maxLevel], c.FmScan.highLvl[c.FmScan.maxLevel-1]...)
				finalFreqs = append(finalFreqs, c.FmScan.highLvl[c.FmScan.maxLevel-2]...)
			} else if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel-1]; ok {
				finalFreqs = append(c.FmScan.highLvl[c.FmScan.maxLevel], c.FmScan.highLvl[c.FmScan.maxLevel-1]...)
			} else if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel]; ok {
				finalFreqs = c.FmScan.highLvl[c.FmScan.maxLevel]
			}
		}

		if c.FmScan.goingUp && c.FmScan.currentFrame < len(finalFreqs)-1 {
			c.FmScan.currentFrame += 1
		} else if c.FmScan.goingUp && c.FmScan.currentFrame == len(finalFreqs)-1 {
			c.FmScan.currentFrame -= 1
			c.FmScan.goingUp = false
		} else if !c.FmScan.goingUp && c.FmScan.currentFrame > 0 {
			c.FmScan.currentFrame -= 1
		} else {
			c.FmScan.currentFrame += 1
			c.FmScan.goingUp = true
		}
		sharedutils.Log("3 " + strconv.FormatInt(int64(c.FmScan.currentFrame+1), 10) + " " + strconv.FormatFloat(finalFreqs[c.FmScan.currentFrame], 'f', 1, 64))

		// if _, ok := finalFreqs[c.FmScan.currentFrame]; ok {
		c.Fm.currentFreq = finalFreqs[c.FmScan.currentFrame]
		// }
	}

	response := tea5767.SetFrequency(c.Devices.radio, c.Fm.currentFreq)

	var signalStrength int64

	if isInitializing {

		signalStrength = int64(response[3] >> 4)
		if c.FmScan.maxLevel < signalStrength {
			c.FmScan.maxLevel = signalStrength

			// cleanup the lower-signal
			for i := 0; i < (int(c.FmScan.maxLevel) - 2); i++ {
				if _, ok := c.FmScan.highLvl[c.FmScan.maxLevel]; ok {
					delete(c.FmScan.highLvl, c.FmScan.maxLevel)
				}
			}
		}

		// only add if in the top channels
		if signalStrength >= (c.FmScan.maxLevel - 2) {
			c.FmScan.highLvl[signalStrength] = append(c.FmScan.highLvl[signalStrength], c.Fm.currentFreq)
			c.FmScan.highLvl[signalStrength] = sharedutils.UniqueFloat(c.FmScan.highLvl[signalStrength])
		}

	}

	display.Clear(c.Devices.screen)

	if isInitializing {
		display.Text(
			c.Devices.screen,
			2,
			15,
			"Initializing...")
		display.Text(
			c.Devices.screen,
			2,
			35,
			"# signal: "+strconv.FormatInt(int64(c.FmScan.totalHighLevel), 10))
		display.Text(
			c.Devices.screen,
			2,
			55,
			"Freq: "+strconv.FormatFloat(c.Fm.currentFreq, 'f', 1, 64))
		if c.FmScan.numberOfCycles > 1 {
			display.FilledRect(
				c.Devices.screen,
				107,
				117,
				1,
				11,
				color.RGBA{100, 100, 100, 255},
			)
		} else {
			display.Rect(
				c.Devices.screen,
				107,
				117,
				1,
				11,
				color.RGBA{100, 100, 100, 255},
			)
		}
		if c.FmScan.numberOfCycles > 2 {
			display.FilledRect(
				c.Devices.screen,
				117,
				127,
				1,
				11,
				color.RGBA{100, 100, 100, 255},
			)

		} else {
			display.Rect(
				c.Devices.screen,
				117,
				127,
				1,
				11,
				color.RGBA{100, 100, 100, 255},
			)
		}
		if c.FmScan.numberOfCycles > 3 {
			display.FilledRect(
				c.Devices.screen,
				107,
				117,
				11,
				21,
				color.RGBA{100, 100, 100, 255},
			)

		} else {
			display.Rect(
				c.Devices.screen,
				107,
				117,
				11,
				21,
				color.RGBA{100, 100, 100, 255},
			)
		}

		display.Rect(
			c.Devices.screen,
			117,
			127,
			11,
			21,
			color.RGBA{100, 100, 100, 255},
		)
	} else {

		display.FilledRect(
			c.Devices.screen,
			107,
			127,
			1,
			21,
			color.RGBA{100, 100, 100, 255},
		)
		display.FilledBolt(
			c.Devices.screen,
			117,
			11,
			color.RGBA{0, 0, 0, 255},
		)

		var dirStr string
		if c.FmScan.goingUp {
			dirStr = "up"
		} else {
			dirStr = "dn"
		}

		display.Text(
			c.Devices.screen,
			2,
			15,
			"Scan "+dirStr+" "+strconv.FormatInt(int64(c.FmScan.currentFrame), 10))
		display.Text(
			c.Devices.screen,
			2,
			35,
			"# channels: "+strconv.FormatInt(int64(c.FmScan.totalHighLevel), 10))
		display.Text(
			c.Devices.screen,
			2,
			55,
			"Freq: "+strconv.FormatFloat(c.Fm.currentFreq, 'f', 1, 64))
	}

	c.Devices.screen.Display()

	c.ChangeFuncCheck("FmScan")

	if isInitializing {
		time.Sleep(50 * time.Millisecond)

	} else {
		time.Sleep(300 * time.Millisecond)

	}
}
