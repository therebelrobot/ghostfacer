package xiaobutton

import (
	"machine"
)

func Init() machine.Pin {
	button := machine.D1
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return button
}
