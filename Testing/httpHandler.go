package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = ":8080"
)

func MyServer() {

	serverMux := http.NewServeMux()
	httpServer := &http.Server{Addr: PORT, Handler: serverMux}

	serverMux.HandleFunc("/meet-and-greet/", meetAndGreetHandler)
	serverMux.HandleFunc("/", slashHandler)

	httpServer.ListenAndServe()
}

func meetAndGreetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text")
	msg := "Greetings!!! You reached the new Golang Server..."
	fmt.Fprintf(w,msg)
}

func slashHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text")
	msg := "nothing here"
	fmt.Fprintf(w,msg)
}

func main() {

	MyServer()

}
