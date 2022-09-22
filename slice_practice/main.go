package main

import "fmt"

func main() {

	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res := removeElement(testSlice, 5)
	fmt.Println(res) // [1 2 3 4 6 7 8 9 10]
}

func removeElement(slice []int, s int) []int {
	idx := 0
	for _, v := range slice {
		if v != s {
			slice[idx] = v
			idx++
		}

	}
	return slice[:idx]
}
