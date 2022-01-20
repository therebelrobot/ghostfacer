package utils

func Scale(num float64, in_min float64, in_max float64, out_min float64, out_max float64) float64 {
	return (num-in_min)*(out_max-out_min)/(in_max-in_min) + out_min
}
