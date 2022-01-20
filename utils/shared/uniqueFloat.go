package utils

func UniqueFloat(floatSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range floatSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
