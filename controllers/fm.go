package ctrl

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6
// https://github.com/tinygo-org/drivers/blob/v0.18.0/examples/ws2812/main.go

import (
	"strconv"
	"time"

	sharedutils "github.com/therebelrobot/ghostfacer/utils/shared"
	tea5767 "github.com/therebelrobot/ghostfacer/utils/tea5767"

	display "github.com/therebelrobot/ghostfacer/utils/xiaoexpansion/display"
)

type Fm struct {
	radio       tea5767.TEA5767
	currentFreq float64
}

func (c *Ctrl) FmInit() {
	sharedutils.Log("EMF")
	c.Devices.radio = tea5767.Init()

	c.Fm.currentFreq = c.Devices.radio.Constants.FreqRangeUs[0]
	c.FmScan.currentFrame = 0
	c.FmScan.goingUp = true
	c.FmScan.numberOfCycles = 1

	tea5767.SetFrequency(c.Devices.radio, c.Fm.currentFreq)

	c.FmScan.highLvl = make(map[int64][]float64)
	c.FmScan.totalHighLevel = 0
	c.FmScan.maxLevel = 0
}

func (c *Ctrl) FmLoop() {

	dialValue := sharedutils.Constrain(float64(c.Devices.dial.Get()), 0, 1023)
	outVal := sharedutils.Scale(dialValue, 0, 1023, 0, 100)

	sharedutils.Log("dialValue " + strconv.FormatInt(int64(dialValue), 10) + " " + strconv.FormatFloat(outVal, 'f', 1, 64))

	if c.Data.shouldRedraw {
		c.Data.shouldRedraw = false
		display.Clear(c.Devices.screen)
		display.Text(
			c.Devices.screen,
			2,
			15,
			"FM Band ∿")
		display.Text(
			c.Devices.screen,
			2,
			55,
			"Freq: "+strconv.FormatFloat(c.Fm.currentFreq, 'f', 1, 64))
		display.Text(
			c.Devices.screen,
			2,
			55,
			"Freq: "+strconv.FormatFloat(c.Fm.currentFreq, 'f', 1, 64))
		c.Devices.screen.Display()

	}

	c.ChangeFuncCheck("Fm")

	time.Sleep(50 * time.Millisecond)

}
