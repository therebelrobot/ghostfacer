package utils

import (
	"machine"
)

var BAUD_RATE uint32 = 9600

func LogSetup() {
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: BAUD_RATE,
		TX:       6,
		RX:       7,
	})
}

func Log(message string) {
	machine.Serial.Write([]byte(message + "\n"))
}
