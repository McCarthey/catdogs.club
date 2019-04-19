package libs

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digital     = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandInt(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digital[rand.Int63()%int64(len(digital))]
	}
	return string(b)
}
