package main

import "fmt"

type Vertex struct {
	// 小文字で宣言すると他のパッケージから見えないようになる
	// privateみたいな
	x, y int
}

// 最初にstructをつけてメソッドを宣言すると
// struct紐つくメソッドとして扱うことができる
// 値レシーバー
func (v Vertex) Area() int {
	return v.x * v.y
}

// vertexのstructを引数に取ることでstructの値を直接書き換えるメソッドを定義できる
// setterのようなもの
// ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// func Area(v Vertex) int {
// 	return v.x * v.y
// }

// コンストラクターのような存在のもの
// package.Newでstructを返却することがGoのお作法
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	// fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())
}
