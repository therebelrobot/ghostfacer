package xiaobuzzer

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

func Init() (buzzer.Device, machine.Pin) {
	buzrPin := machine.D3
	buzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker := buzzer.New(buzrPin)

	return speaker, buzrPin
}
