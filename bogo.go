package bogo

import "fmt"
import "sort"
import "math/rand"

func Sort(arr []int) {
	for {
		if sort.IntsAreSorted(arr) {
			fmt.Println("\nArray sorted!")
			fmt.Println("=============\n")

			for _, element := range arr {
				fmt.Printf("%d ", element)
			}
			break
		}

		arr = shuffle(arr[:])
		fmt.Println(arr)
	}
}

func shuffle(arr []int) []int {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
