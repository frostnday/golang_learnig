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

type Vertex3D struct {
	// ここに他のstructを宣言することで継承したことになる
	Vertex

	// 追加も可能
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i

}

// コンストラクターのような存在のもの
// package.Newでstructを返却することがGoのお作法
// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	// fmt.Println(Area(v))
	v.Scale3D(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())

}
