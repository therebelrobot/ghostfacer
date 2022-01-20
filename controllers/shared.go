package ctrl

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"

	sharedutils "github.com/therebelrobot/ghostfacer/utils/shared"
	tea5767 "github.com/therebelrobot/ghostfacer/utils/tea5767"

	xiaobutton "github.com/therebelrobot/ghostfacer/utils/xiaoexpansion/button"
	display "github.com/therebelrobot/ghostfacer/utils/xiaoexpansion/display"
)

type Data struct {
	shouldRedraw    bool
	Functions       []string
	CurrentFunction string
	isButtonPressed bool
}

type Devices struct {
	radio  tea5767.TEA5767
	screen ssd1306.Device
	button machine.Pin
	dial   machine.ADC
}

type Ctrl struct {
	Data    Data
	Devices Devices

	Fm     Fm
	FmScan FmScan
}

func (c *Ctrl) ChangeFuncCheck(function string) {
	if !c.Devices.button.Get() && !c.Data.isButtonPressed {
		currentFuncIndex := sharedutils.GetIndex(len(c.Data.Functions), func(i int) bool {
			return c.Data.Functions[i] == function
		})
		if currentFuncIndex >= 0 {
			c.Data.isButtonPressed = true
			if currentFuncIndex == len(c.Data.Functions)-1 {
				c.Data.CurrentFunction = c.Data.Functions[0]
			} else {
				c.Data.CurrentFunction = c.Data.Functions[currentFuncIndex+1]
			}
			c.Data.shouldRedraw = true
		}
	} else if c.Devices.button.Get() && c.Data.isButtonPressed {
		c.Data.isButtonPressed = false
	}
}

func (c *Ctrl) SharedInit() {

	c.Data.Functions = []string{"FmScan", "Fm"}
	c.Data.CurrentFunction = "FmScan"
	c.Data.isButtonPressed = false

	sharedutils.LogSetup()
	sharedutils.Log("Initializing")
	c.Devices.screen = display.Init()
	c.Devices.button = xiaobutton.Init()

	machine.InitADC()
	c.Devices.dial = machine.ADC{Pin: machine.A0}
	c.Devices.dial.Configure(machine.ADCConfig{})

	c.FmInit()

	c.Data.shouldRedraw = false
}
