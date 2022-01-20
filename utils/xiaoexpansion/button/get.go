package xiaobutton

import (
	"machine"
)

func Get(button machine.Pin) bool {
	return !button.Get()
}
