package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	f := fib(5)
	fmt.Printf("fib 5 is %d \n", f)
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
