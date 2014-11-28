package bogo

import "sort"
import "math/rand"

func Sort(arr []int) []int {
	for {
		if sort.IntsAreSorted(arr) {
			return arr
		}

		arr = shuffle(arr[:])
	}
}

func shuffle(arr []int) []int {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
