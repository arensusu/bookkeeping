package random

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphanumeric = "abcdefghijklmnopqruvstwxyz0123456789"
)

var RandGen *rand.Rand

func init() {
	RandGen = rand.New(rand.NewSource(time.Now().UnixMilli()))
}

func randomInt64(min, max int64) int64 {
	return RandGen.Int63n(max-min+1) + min
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphanumeric)
	for i := 0; i < n; i += 1 {
		sb.WriteByte(alphanumeric[RandGen.Intn(k)])
	}
	return sb.String()
}

func randomBool() bool {
	return RandGen.Intn(2)%2 == 0
}

func Username() string {
	return randomString(6)
}

func Password() string {
	return randomString(8)
}

func IsAdmin() bool {
	return randomBool()
}

func Category() string {
	return randomString(10)
}

func Cost() int64 {
	return randomInt64(50, 1000)
}

func Date() time.Time {
	return time.UnixMilli(time.Now().UnixMilli() + randomInt64(-36400, 36400))
}
