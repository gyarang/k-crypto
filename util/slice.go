package util

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
