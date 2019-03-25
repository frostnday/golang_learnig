package main

import "fmt"

// 可変長引数を受け取れる
func foo(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func main() {
	foo()
	foo(10, 20)
	foo(110, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)

	// sliceを展開している
	foo(s...)
}
