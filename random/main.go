package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(700)
	fmt.Println(n)
}
