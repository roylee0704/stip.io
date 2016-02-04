// Bubble sorts the first n elements of an array of integers ([]int) into
// accending order
package main

import "fmt"

func main() {
	val := []int{6, 5, 4, 3, 2, 9, 1, -1, 22, 30}

	bubble(val, 400)
	fmt.Printf("%v\n", val)
}

//
func bubble(v []int, limit int) {

	n := len(v)
	if limit > n {
		limit = n
	}

	for pass := limit - 1; pass >= 0; pass-- {
		for swap := 0; swap < pass; swap++ {
			if v[swap] > v[swap+1] {
				v[swap], v[swap+1] = v[swap+1], v[swap]
			}
		}
	}
}
