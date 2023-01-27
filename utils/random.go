package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefhijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer btw min & max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random name
func RandomOwner() string {
	return RandomString(6)
}

// RandomAmount generates a random amount
func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{
		"USD",
		"GBP",
		"EUR",
	}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}
