package main

import "fmt"

func fibonacci() func() int {
	pre := 0
	next := 1
	expected := 0

	return func() int {
		defer func() {
			expected = next
			next = expected + pre
			pre = expected
		}()
		return expected
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
