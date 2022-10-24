package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 123
	text := "456"

	itoaResult := strconv.Itoa(num)
	fmt.Println(itoaResult == "123") // true

	atoiResult, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(atoiResult == 456) // true

	invalidNumStr := "abc123"
	wrongAtoiResult, err := strconv.Atoi(invalidNumStr)
	if err != nil {
		fmt.Println(err)
		// strconv.Atoi: parsing "abc123": invalid syntax
	}
	fmt.Println(wrongAtoiResult)
}
