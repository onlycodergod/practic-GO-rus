package task_01

import "hw/module03/wordz"

func City() string {
	wordz.Prefix = ""
	wordz.Words = []string{"Moscow", "Berlin", "Paris ", "Dublin", "Astana"}

	return wordz.Random()
}
