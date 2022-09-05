package util

import (
	"math/rand"
	"time"
)

func IsItemInSlice[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func GetIndexOfItem[T comparable](arr []T, item T) int {
	for i, v := range arr {
		if item == v {
			return i
		}
	}
	return -1
}

func RandomSelect[T any](arr []T) T {
	rand.Seed(time.Now().Unix())
	return arr[rand.Intn(len(arr))]
}
