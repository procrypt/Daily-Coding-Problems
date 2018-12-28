/*
This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

Implement car and cdr.
*/

package main

import "fmt"

func con(a, b interface{}) []interface{} {
	var out []interface{}
	out = append(out, a, b)
	return out
}

func car(array []interface{}) interface{} {
	p := func(array []interface{}) interface{} {
		return array[0]
	}(array)
	return p
}

func cdr(array []interface{}) interface{} {
	q := func(array []interface{}) interface{} {
		return array[1]
	}(array)
	return q
}

func main() {
	data := con("Hey", "Hi")
	fmt.Println(car(data))
	fmt.Println(cdr(data))
}
