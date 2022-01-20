package tea5767

import (
	"machine"
)

func Init() TEA5767 {
	config := machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
		SCL:       machine.SCL_PIN,
		SDA:       machine.SDA_PIN,
	}
	machine.I2C0.Configure(config)
	return TEA5767{
		I2c:                       machine.I2C0,
		Address:                   0x60,
		Frequency:                 0.0,
		BandLimits:                "US",
		StereoMode:                true,
		SoftMuteMode:              true,
		StereoNoiseCancellingMode: true,
		HighCutMode:               true,
		StandbyMode:               false,
		MuteMode:                  false,
		SearchMode:                false,
		SearchDirection:           1,
		SearchAdcLevel:            7,
		IsReady:                   false,
		IsStereo:                  false,
		SignalAdcLevel:            0,
		Constants: Constants{
			FreqRangeUs: FREQ_RANGE_US,
			FreqRangeJp: FREQ_RANGE_JP,
			Adc:         ADC,
			AdcBit:      ADC_BIT,
		}}
}
