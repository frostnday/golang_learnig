package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	// => [72 73]

	fmt.Println(string(b))
	// => HI
}
