package main

import "fmt"

func main() {

	// fmt.Println(1)

	// defer fmt.Println(2)

	// defer fmt.Println(3)

	// defer fmt.Println(4)

	// fmt.Println(5)

	num := 1

	// defer deferPrint(num)

	defer func() {
		deferPrint(num)
	}()

	num++
}

func deferPrint(num int) {
	fmt.Println(num)
}
