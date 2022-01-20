package xiaodisplay

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"
)

func Init() ssd1306.Device {
	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()

	return display
}
