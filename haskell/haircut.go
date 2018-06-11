package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	shop := make(chan int, 3)
	go barber(shop)
	start := time.Now()
	for i := 0; ; i++ {
		r := 100 + rand.Intn(200)
		//fmt.Println(r)
		time.Sleep(time.Duration(r) * time.Millisecond)
		if time.Since(start) > time.Second*100 {
			break
		}
		customer(i, shop)
	}
}
func customer(id int, shop chan<- int) {
	if len(shop) < cap(shop) {
		shop <- id
	}
}

func barber(shop <-chan int) {
	count := 1
	for {
		fmt.Printf("Barber cuts hair of customer %d, and he is the %dth\n", <-shop, count)
		time.Sleep(200 * time.Millisecond)
		count = count + 1
	}
}
