package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	// to make it random, you should do the following
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
