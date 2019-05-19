package main

import (
	"bytes"
	"net/http"
)

func main() {

	barray := []byte("hello world")

	bytes.NewBuffer(barray)

	resp, err := http.Post("httpL//google.com","application/octec-stream",io.reader)

}
