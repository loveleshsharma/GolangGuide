package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"
)

func main() {
	// input := "foo\nbar\nbaz"
	input := make([]byte, 1000)
	input[0] = byte("foo\n")
	input[1] = "bar\n"
	input[2] = "abc"

	scanner := bufio.NewScanner(bytes.NewReader(input))
	// Not actually needed since it’s a default split function.
	scanner.Split(mySplit)
	buf := make([]byte, 100)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)
	for scanner.Scan() {
		str := scanner.Text()

		if str != "" {
			fmt.Println("->", str)
		}
		time.Sleep(200 * time.Millisecond)

	}

	prefix := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	s := string(prefix)
	fmt.Println("#", strings.Trim(s, " "), "#")
	fmt.Println("#", len(string(prefix)), "#")
	// if !bytes.Contains(prefix, []byte{'\n'}) {
	if string(prefix) == " " {
		fmt.Println("found")
	}

}

func mySplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		fmt.Print("1")
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		fmt.Print("2")
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		fmt.Print("3")
		fmt.Println("advance: ", len(data))
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
