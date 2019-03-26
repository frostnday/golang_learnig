package main

// 使用するモジュールをinport する
import (
	"fmt"
)

// mapとは連想配列みたいなやつ
func init() {
	fmt.Println("map")
}

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	// => map[banana:200 apple:100]

	fmt.Println(m["apple"])
	// => 100

	m["banana"] = 300
	fmt.Println(m["banana"])
	// => 300

	m["new"] = 500
	fmt.Println(m)
	// => map[apple:100 banana:300 new:500]

	fmt.Println(m["nothing"])
	// => 0

	v, ok := m["apple"]
	fmt.Println(v, ok)
	// => 100 true
	// 戻り値として値と値の存在可否が受け取れる

	// メモリ上にからのmapを作ってからそれに対して代入する
	m2 := make(map[string]int)
	m2["pc"] = 5000

	fmt.Println(m2)
	// => map[pc:5000]

	// こちらは宣言はしてあるが、メモリ上に確保していないので
	// エラーになる
	// var m3 map[string]int
	// m3["pc"] = 5000
}
