package utils

import (
	"bytes"
	"strconv"
)

func FloatArrayToString(A []float64, delim string) string {

	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.FormatFloat(A[i], 'f', 1, 64))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}
