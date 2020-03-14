package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lengthOfArray := 10
	array := make([]int, lengthOfArray)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range array {
		array[i] = rand.Intn(10 + 1) // 0-10 - rand.Intn [0, n)
	}
	fmt.Println(array)
	start := time.Now()
	for passes := 0; passes < lengthOfArray; passes++ {
		arrayChanged := false
		for i := 0; i < lengthOfArray-passes; i++ {
			//fmt.Printf("passes: %d, i: %d\n", passes, i)
			if i == lengthOfArray-1 {
				break
			}
			v1 := array[i]
			v2 := array[i+1]
			if v1 > v2 {
				arrayChanged = true
				array[i+1] = v1
				array[i] = v2
			}
		}
		if !arrayChanged {
			break
		}
	}
	totalTime := time.Since(start)
	fmt.Println(array)
	fmt.Printf("sort took %s to sort array of %d elements\n", totalTime, lengthOfArray)
}
