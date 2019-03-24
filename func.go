package main

import "fmt"

func add(x, y int) (int, int) {
	fmt.Println("add function")
	fmt.Println(x + y)
	return x + y, x - y
}

// 戻り値が複数ある場合はresultのように名前をつける
func cal(price, item int) (result int) {
	result = price * item
	// 戻り値指定のところで宣言しているので
	// ここでは変数を宣言できない
	// result := price * item

	// resultと指定しなくてもいい
	return
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)
}
