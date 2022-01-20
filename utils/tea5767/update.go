package tea5767

import (
	"math"

	sharedutils "github.com/therebelrobot/ghostfacer/utils/shared"
)

var bool2int = map[bool]int{false: 0, true: 1}

func Update(dev TEA5767) []byte {
	if dev.BandLimits == "JP" {
		dev.Frequency = math.Min(math.Max(dev.Frequency, dev.Constants.FreqRangeJp[0]), dev.Constants.FreqRangeJp[1])
	} else {
		dev.BandLimits = "US"
		dev.Frequency = math.Min(math.Max(dev.Frequency, dev.Constants.FreqRangeUs[0]), dev.Constants.FreqRangeUs[1])
	}

	freqB := 4 * (dev.Frequency*1000000 + 225000) / 32768
	buf := [5]byte{}
	buf[0] = byte(int(freqB)>>8 | bool2int[dev.MuteMode]<<7 | bool2int[dev.SearchMode]<<6)
	buf[1] = byte(int(freqB) & 0xff)
	buf[2] = byte(dev.SearchDirection<<7 | 1<<4 | bool2int[dev.StereoMode]<<3)
	// buf[2] += byte(dev.Constants.adcBit[dev.Constants.adc.index(dev.SearchAdcLevel)] << 5)

	buf[3] = byte(bool2int[dev.StandbyMode]<<6 | bool2int[(dev.BandLimits == "JP")]<<5 | 1<<4)
	buf[3] += byte(bool2int[dev.SoftMuteMode]<<3 | bool2int[dev.HighCutMode]<<2 | bool2int[dev.StereoNoiseCancellingMode]<<1)
	buf[4] = 0
	response := make([]byte, 4)
	dev.I2c.Tx(dev.Address, buf[:], response)
	return response
}

func SetFrequency(dev TEA5767, freq float64) []byte {
	dev.Frequency = freq
	return Update(dev)
}
func GetFrequency(dev TEA5767) float64 {
	return dev.Frequency
}
func ChangeFrequency(dev TEA5767, change float64) {
	dev.Frequency += change
	if change >= 0 {
		dev.SearchDirection = 1
	} else {
		dev.SearchDirection = 0
	}
	Update(dev)
}

func Search(dev TEA5767, mode bool, dir int, adc int) {
	dev.SearchMode = mode
	dev.SearchDirection = dir
	if sharedutils.ContainsInt(dev.Constants.Adc, adc) {
		dev.SearchAdcLevel = adc
	} else {
		dev.SearchAdcLevel = 7
	}
	Update(dev)
}

func Mute(dev TEA5767, mode bool) {
	dev.MuteMode = mode
	Update(dev)
}

func Standby(dev TEA5767, mode bool) {
	dev.StandbyMode = mode
	Update(dev)
}
