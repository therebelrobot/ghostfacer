package tea5767

import (
	"machine"
)

type Constants struct {
	FreqRangeUs []float64
	FreqRangeJp []float64
	Adc         []int
	AdcBit      []int
}

type TEA5767 struct {
	I2c                       *machine.I2C
	Address                   uint16
	Frequency                 float64
	BandLimits                string
	StandbyMode               bool
	MuteMode                  bool
	SoftMuteMode              bool
	SearchMode                bool
	SearchDirection           int
	SearchAdcLevel            int
	StereoMode                bool
	StereoNoiseCancellingMode bool
	HighCutMode               bool
	IsReady                   bool
	IsStereo                  bool
	SignalAdcLevel            int
	Constants                 Constants
}
