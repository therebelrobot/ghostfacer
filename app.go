package main

// https://create.arduino.cc/projecthub/alankrantas/tinygo-on-arduino-uno-an-introduction-6130f6

import (
	ctrl "github.com/therebelrobot/ghostfacer/controllers"
	utils "github.com/therebelrobot/ghostfacer/utils/shared"
)

func main() {
	c := ctrl.Ctrl{}
	c.SharedInit()
	utils.Log("starting loop")
	for {
		for {
			utils.Log("Current Function: " + c.Data.CurrentFunction)

			if c.Data.CurrentFunction != c.Data.Functions[0] {
				utils.Log("Breaking FmScan")
				break
			}
			c.FmScanLoop()
			// shouldChangeFunction := ctrl.FmScanLoop(screen, radio, button, currentFreq, currentFrame, goingUp, numberOfCycles, highLvl, totalHighLevel, maxLevel)

			utils.Log("hello")
		}
		for {
			utils.Log("Current Function: " + c.Data.CurrentFunction)

			if c.Data.CurrentFunction != c.Data.Functions[1] {
				utils.Log("Breaking FmScan")
				break
			}
			utils.Log("goodbye")
			c.FmLoop()
		}
	}
}
