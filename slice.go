package main

// 使用するモジュールをinport する
import (
	"fmt"
)

func init() {
	fmt.Println("slice")
}

func main() {
	/**
	--------配列とスライス----------
	**/
	var list [2]int
	list[0] = 100
	list[1] = 200
	fmt.Println(list)
	// => [100 200]

	// デフォルトで入れておく場合
	// 配列は要素数と型は変更できないので注意
	var list2 [2]int = [2]int{100, 200}

	// 要素を追加するメソッドを使ってもエラーとなる
	// list2 = append(b, 300)
	fmt.Println(list2)

	// スライス 要素数を宣言しなければ追加できる
	var list3 []int = []int{100, 200}
	list3 = append(list3, 300)
	fmt.Println(list3)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	// => [1 2 3 4 5 6]

	fmt.Println(slice[2])
	// => 3

	fmt.Println(slice[2:4])
	// => [3 4]

	fmt.Println(slice[:2])
	// => [1 2]

	fmt.Println(slice[2:])
	// => [3 4 5 6]

	//入れ子のスライス
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}

	fmt.Println(board)
	// => [[0 1 2] [3 4 5] [6 7 8]]

	// 長さが3、キャパシティ5のスライス
	// 引数を1つにすると長さもキャパシティも同時に宣言できる
	slice2 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(slice2), cap(slice2), slice2)
	// => len=3 cap=5 value=[0 0 0]

	slice2 = append(slice2, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(slice2), cap(slice2), slice2)
	// => len=5 cap=5 value=[0 0 0 0 0]

	slice2 = append(slice2, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(slice2), cap(slice2), slice2)

	slice3 := make([]int, 0) // メモリにsliceを確保する
	var slice4 []int         // メモリにsliceを確保しない
	fmt.Printf("len=%d cap=%d value=%v\n", len(slice3), cap(slice3), slice3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(slice4), cap(slice4), slice4)

	fmt.Println("slice5\n")
	slice5 := make([]int, 5)
	for q := 0; q < 5; q++ {
		slice5 = append(slice5, q)
		fmt.Println(slice5)
	}
	fmt.Println(slice5)
	// [0 0 0 0 0 0]
	// [0 0 0 0 0 0 1]
	// [0 0 0 0 0 0 1 2]
	// [0 0 0 0 0 0 1 2 3]
	// [0 0 0 0 0 0 1 2 3 4]
	// [0 0 0 0 0 0 1 2 3 4]

	fmt.Println("slice6\n")
	slice6 := make([]int, 0, 5)
	for z := 0; z < 5; z++ {
		slice6 = append(slice6, z)
		fmt.Println(slice6)
	}
	fmt.Println(slice6)

	// [0]
	// [0 1]
	// [0 1 2]
	// [0 1 2 3]
	// [0 1 2 3 4]
	// [0 1 2 3 4]

	// 長さを宣言すると先のその分の領域が確保されるので
	// appendすると長さ + 追加要素となる？
}
