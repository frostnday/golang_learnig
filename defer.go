package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

// 遅延実行のこと
// 記述された関数内の最後に実行される
func main() {

	foo()

	defer fmt.Println("world")

	fmt.Println("hello") // こっちが先に実行される

	// 一番最初に入れたものが最後に実行される
	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	// => run success 3 2 1の順で表示される

	// deferってどんなときにつかうのか

	// ファイルを開きます
	file, _ := os.Open("./for.go")

	// 一番最後の処理でファイルをcloseします
	defer file.Close()

	// バイト配列を宣言します
	data := make([]byte, 100)

	// データを読み込みます
	file.Read(data)

	// データを出力します
	fmt.Println(string(data))
}
