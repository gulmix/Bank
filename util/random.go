package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var st strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		st.WriteByte(c)
	}
	return st.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"EURO", "USD", "CAD", "TON"}
	return currencies[rand.Intn(len(currencies))]
}
