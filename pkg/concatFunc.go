package pkg

import (
	"math/rand"
	"strings"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length  = 25
)

func Concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func СoncatOptimized(str []string) string {
	var buffer strings.Builder
	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func GenerateStrings(n int) []string {
	rand.Seed(time.Now().UnixNano())
	var strs []string
	for i := 0; i < n; i++ {
		var b strings.Builder
		for j := 0; j < length; j++ {
			b.WriteByte(charset[rand.Intn(len(charset))])
		}
		strs = append(strs, b.String())
	}
	return strs
}
