package xiaobattery

import (
	"machine"
)

// ChargeStatus returns the state of the battery.
type ChargeStatus uint8

// Charge status of the battery: discharging, charging, and fully charged.
const (
	Discharging ChargeStatus = iota + 1
	Charging
	FullyCharged
)

type voltagePercentPosition struct {
	millivolts int
	percent    int
}

// The two pins connected to the power regulator chip, that indicate charging
// status and power presence.
const (
	pinBatteryCharging machine.Pin = 12
	pinPowerConnected  machine.Pin = 19
)

// Voltage to percent mappings. Values in between can be linearly approximated.
// This is a rough fitting, better fits are likely possible.
var voltagePercentPositions = []voltagePercentPosition{
	{3880, 100},
	{3780, 80},
	{3690, 60},
	{3640, 40},
	{3610, 20},
	{3520, 0},
}

// voltageToPercent calculates the percentage the battery is full based on a
// linear approximation with multiple points on the graph. The points must be in
// order, from full to empty. The first entry in the slice must have
// percent==100, the last entry must have percent==0.
func voltageToPercent(millivolts int, pointsOnGraph []voltagePercentPosition) int {
	if millivolts >= pointsOnGraph[0].millivolts {
		return 100
	}
	if millivolts <= pointsOnGraph[len(pointsOnGraph)-1].millivolts {
		return 0
	}
	for i := 0; i < len(pointsOnGraph)-1; i++ {
		if millivolts < pointsOnGraph[i+1].millivolts {
			continue
		}
		// Voltage is between pointsOnGraph[i] and pointsOnGraph[i+1].
		high := pointsOnGraph[i]
		low := pointsOnGraph[i+1]
		return high.percent + (high.percent-low.percent)*(millivolts-high.millivolts)/(high.millivolts-low.millivolts)
	}
	// unreachable
	return 0
}

// ---------------------

// BatteryStatusRaw reads and returns the current battery voltage in millivolts
// (mV) and returns the current charging status of the battery (discharging,
// charging, full).
func BatteryStatusRaw() (millivolt int, status ChargeStatus) {
	if !pinPowerConnected.Get() {
		// Power is connected.
		if !pinBatteryCharging.Get() {
			// Battery is charging.
			status = Charging
		} else {
			status = FullyCharged
		}
	} else {
		status = Discharging
	}
	value := machine.ADC{31}.Get()
	return int(value) * 2000 / (65535 / 3), status
}

// BatteryStatus reads and returns the current battery status (percent and
// charge status).
func BatteryStatus() (millivolts, percent int, status ChargeStatus) {
	millivolts, status = BatteryStatusRaw()
	percent = voltageToPercent(millivolts, voltagePercentPositions)
	return
}
