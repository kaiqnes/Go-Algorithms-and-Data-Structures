package utils

import (
	"math/rand"
	"time"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func GetRandomIntSlice(size int) (arr []int) {
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(100))
	}
	return
}

func GetSortedIntSlice(size int) (arr []int) {
	for i := 0; i < size; {
		i++
		arr = append(arr, i)
	}
	return
}
