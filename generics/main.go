package main

import "fmt"

type MyNumType interface {
	int64 | float64
}

func sumNums[numType MyNumType](nums []numType) numType {
	sum := numType(0)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {

	nums := []float64{1.1, 2, 3, 4, 5}
	sum := sumNums[float64](nums)
	fmt.Printf("Sum: %v, Type: %T \n", sum, sum)

	nums2 := []int64{1, 2, 3, 4, 5}
	sum2 := sumNums[int64](nums2)
	fmt.Printf("Sum: %v, Type: %T \n", sum2, sum2)
}
