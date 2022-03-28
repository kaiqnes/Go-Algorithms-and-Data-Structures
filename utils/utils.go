package utils

import (
	"math/rand"
	"time"
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
