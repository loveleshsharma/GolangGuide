package main

import (
	"fmt"
)

func main() {
	defer check()
	ch := make(chan string)
	go pattern(5, ch)

	// time.Sleep(3 * time.Second)
	for p := range ch {
		fmt.Println(p)
	}
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-ch)
	// }
}

func pattern(n int, c chan string) {
	var str string
	for i := 1; i <= n; i++ {
		str = ""
		for j := 0; j < i; j++ {
			str = str + "*"
		}
		c <- str
	}
	// close(c)
}

func check() {
	fmt.Println("check called!")
}
