package utils

func Constrain(num float64, min float64, max float64) float64 {
	if num > max {
		return max
	}
	if num < min {
		return min
	}
	return num
}
