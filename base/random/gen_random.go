package random

import (
	"math/rand"
	"time"
)

const (
	choicesStr   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	choicesAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	choicesNum   = "0123456789"
)

func genRandom(choices string, length int) string {
	bytes := []byte(choices)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GenRandomStr(length int) string {
	return genRandom(choicesStr, length)
}

func GenRandomAlpha(length int) string {
	return genRandom(choicesAlpha, length)
}

func GenRandomNum(length int) string {
	return genRandom(choicesNum, length)
}
