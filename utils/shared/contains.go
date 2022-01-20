package utils

func ContainsInt(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}
