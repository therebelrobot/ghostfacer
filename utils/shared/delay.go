package utils

import (
	"time"
)

func Delay(ms int64) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
