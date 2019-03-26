package main

import "fmt"

// 値渡し
func one(x int) {
	x = 1
}

// 参照渡し
func one2(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	fmt.Println(n)

	// &をつけることで100を入れたメモリのアドレスが表示される
	fmt.Println(&n)

	// *intでメモリ領域を入れる用の型を宣言でき、そこに100が入っているメモリアドレスを入れる
	var p *int = &n

	// => メモリアドレスが表示される
	fmt.Println(p)

	// *をつけることでそのメモリアドレスに実際に履いている値が表示される
	// => 100
	fmt.Println(*p)

	var m int = 100
	one(m)
	fmt.Println(m)
	// => 100

	one2(&m)
	fmt.Println(m)
}
