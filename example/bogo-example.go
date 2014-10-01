package main

import "github.com/jsec/bogo"

func main() {
	var array = []int{42, 51, 62, 21, 91, 12, 33, 102, 49, 88}
	bogo.Sort(array)
}
