package wordz

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Hello = "This is wordz package"
var Prefix = "Random word: "
var Words = []string{"One", "Two", "Three", "Four", "Five"}

func init() {
	fmt.Println("Function init in package wordz")
}

func Random() string {
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + Words[index]
}
