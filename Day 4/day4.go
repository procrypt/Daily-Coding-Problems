/*
This problem was asked by Stripe.
Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.
For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(a []int) int {
	sort.Ints(a)
	fmt.Println(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] < 0 {
			continue
		}
		if max(a[i], a[i+1]) == 1 {
			continue
		}
		if max(a[i], a[i+1]) > 1 {
			return a[i]+1
		}
	}
	return a[len(a)-1]+1
}

func max(a, b int) int {
	out := math.Abs(float64(a)-float64(b))
	return int(out)
}

func main() {
	a := []int{1,2,3,4,5}
	fmt.Println(solution(a))
}