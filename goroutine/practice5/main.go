package main

import (
	"fmt"
	"sync"
)

func main() {

	count := 100
	amount := 0

	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func(num int, wg *sync.WaitGroup) {
			defer wg.Done()

			mx.Lock()
			defer mx.Unlock()
			amount += num

		}(i, wg)
	}

	wg.Wait()

	fmt.Println(amount) // 499500
}
