package main

import (
	"fmt"
	"net/http"
)

func main() {
	defer check()
	links := []string{
		"https://facebook.com",
		"https://google.com",
		"https://linkedin.com",
		"https://youtube.com",
	}

	// Channel
	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	// for l := range ch {
	// 	// func() {
	// 	// time.Sleep(time.Second)
	// 	fmt.Println(l)
	// 	// }()
	// }

	if <-ch == "is up" {
		fmt.Println("it is up")
	} else {
		fmt.Println("it is down")
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, " is down!")
		ch <- "is down"
		return
	}
	// fmt.Println(link, "is up!")
	ch <- "is up"
}

func check() {
	fmt.Println("check called!")
}
