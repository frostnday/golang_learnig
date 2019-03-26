package main

import "fmt"

func main() {
	// メモリの領域は確保している
	var p *int = new(int)
	fmt.Println(p)
	// => メモリのアドレス

	var p2 *int
	fmt.Println(p2)
	// => nil

	// データ構造を初期化し、使用可能となるよう値を準備
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	// メモリ領域を確保し*Tの初期値をいれておく
	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
