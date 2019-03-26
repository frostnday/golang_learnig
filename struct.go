package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

// こっちは値渡し
func changeVertex(v Vertex) {
	v.X = 1000
}

// こっちは参照渡し
func changeVertex2(v *Vertex) {
	// 本来であればこちらを使うべきだが、structの場合は*を使えば
	// 自動的にポインタの実態を指してくれるので問題ない
	// (*v).X = 1000

	v.X = 1000
}

// staticなObjectみたいな感じ
func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	// => {1 2 }
	fmt.Println(v.X, v.Y)
	// => 1 2

	v.X = 100
	fmt.Println(v.X, v.Y)
	// => 100 2

	v2 := Vertex{X: 1}
	fmt.Println(v2)
	// => {1 0}

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)
	// => {1 2 test}

	v4 := Vertex{}
	fmt.Println(v4)
	// => {0 0}

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)
	// => main.Vertex {0 0 }

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)
	// => *main.Vertex &{0 0 }

	// structの場合はこっちを使うことがおおい
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)
	// => *main.Vertex &{0 0 }

	vf := Vertex{1, 2, "test"}
	changeVertex(vf)
	fmt.Println(vf)

	vf2 := &Vertex{1, 2, "test"} //&をつけているのでアドレスになる
	changeVertex2(vf2)
	fmt.Println(*vf2)

}
