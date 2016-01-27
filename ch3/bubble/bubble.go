// Bubble sorts the first n elements of an array of integers ([]int) into
// accending order
package main

import "fmt"

func main() {
	val := []int{6, 5, 4, 3, 2, 9}

	bubble(val, 400)
	fmt.Printf("%v\n", val)

}

//
func bubble(v []int, n int) {
	for pass := len(v) - 1; pass >= 0; pass-- {
		for compare := 0; compare < pass; compare++ {
			if v[compare] > v[compare+1] {
				//swap
				v[compare], v[compare+1] = v[compare+1], v[compare]
			}
		}
	}

}
