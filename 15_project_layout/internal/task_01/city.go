package task_01

import (
	"gopackages-layout/pkg/wordz"
)

func City() string {
	wordz.Prefix = ""
	wordz.Words = []string{"Moscow", "Berlin", "Paris ", "Dublin", "Astana"}

	return wordz.Random()
}
