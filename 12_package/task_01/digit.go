package task_01

import (
	"hw/module03/wordz"
)

func Digit() string {
	wordz.Prefix = ""
	wordz.Words = []string{"one", "two", "three ", "four", "five", "six", "seven"}

	return wordz.Random()
}
