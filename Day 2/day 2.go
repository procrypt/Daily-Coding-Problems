/*
This problem was asked by Uber.

Given an array of integers, return a new array such that each element at
index i of the new array is the product of all the numbers in the original array except the one at i.
For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Edge cases.
If the array contains one 0, then all the other elements will be 0, except the 0 element, which will be the product of all other elements.
If there are two or more 0, then the entire array will be zero.

To get the product at index i: Calculate the total product of the array and at index i divide the product by the value at that index.

*/


package main

import "fmt"


// Time Complexity O(n) | Space O(n)
func solution(a []int) []int {
	// If the length of the array is 1, there is no element we can divide it to so we just return the array.
	if len(a) < 2 {
		return a
	}

	// Range over the array and count the number of zeros
	zero := 0
	product := 1
	for _, v := range a {
		if v == 0 {
			zero++
			// Don't execute the next instruction. Just start the loop again.
			continue
		}
		product *= v
	}
	if zero > 1 {
		return make([]int, len(a))
	}

	for i := range a {
		if a[i] == 0 {
			a[i] = product
			continue
		}
		if zero == 1 {
			a[i] = 0
			continue
		}
		a[i] = product/a[i]
	}
	return a
}


func main() {

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(solution(a))

}
