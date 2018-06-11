package main

import "fmt"

func main() {
	f := fib()
	for i := 1; i < 100; i++ {
		fmt.Println(f())
	}
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
