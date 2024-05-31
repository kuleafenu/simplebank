package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of n
func RandomString(n int) string {
	var stringBuilder strings.Builder
	alphaLen := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(alphaLen)]
		stringBuilder.WriteByte(c)

	}

	return stringBuilder.String()

}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
