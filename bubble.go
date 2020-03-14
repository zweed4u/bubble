package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	length_of_array := 10
	array := make([]int, length_of_array)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range array {
		array[i] = rand.Intn(10+1) // 0-10 - rand.Intn [0, n)
	}
	fmt.Println(array)
	start := time.Now()
	for passes := 0; passes < length_of_array; passes++ {
		array_changed := false
		for i := 0; i < length_of_array-passes; i++ {
			//fmt.Printf("passes: %d, i: %d\n", passes, i)
			if i == length_of_array - 1 {
				break
			}
			v1 := array[i]
			v2 := array[i+1]
			if v1 > v2 {
				array_changed = true
				array[i+1] = v1
				array[i]  = v2
			}
		}
		if !array_changed {
			break
		}
	}
	total_time := time.Since(start)
	fmt.Println(array)
	fmt.Printf("sort took %s to sort array of %d elements\n", total_time, length_of_array)
}
