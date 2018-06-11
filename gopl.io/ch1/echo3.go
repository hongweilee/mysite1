package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "), time.Since(start))
	fmt.Println(os.Args[1:], time.Since(start))
}
