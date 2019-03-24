package main

// 使用するモジュールをinport する
import (
	"fmt"
	"os/user"
	"strconv"
	"time"
)

// 定数(変更されたら困るデータ)
// 型は実行時に解釈される
// 定数が untyped であることを許容しており、変数に代入するときや式の中で型が決定する
const pi = 3.14

// 1 文字目が 小文字 の場合は、そのパッケージだけで見える変数
// 1 文字目が 大文字 の場合は、他のパッケージからも見える変数
const (
	Username = "test_user"
	Password = "test_password"
)

func init() {
	fmt.Println("basic")
}

func main() {
	/**
	--------変数----------
	**/

	// ()でくくることで、型の違う複数の変数を宣言できる
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

	// 関数内でしか使用できない
	// :=とすることでvarを省略し、型宣言と代入同時にできる

	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
	fmt.Printf("%T\n", xf64)

	fmt.Println("hello world\n", time.Now())
	fmt.Println(user.Current())

	fmt.Println(pi, Username, Password)

	/**
	--------型変換----------
	**/
	var x int = 1
	xx := float64(x)

	// 変数の型、値、型変換での値
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var st string = "14s"
	// 文字列を数値には変換できない
	// z = int(s)

	// ここの_は本来はエラーが返ってくるが、エラーをロジック内で使用する必要のない場合は
	// _ とする
	// iii, _ := strconv.Atoi(st)

	eee, err := strconv.Atoi(st)
	if err != nil {
		// もしエラーが空じゃなかったら
		fmt.Println("数値に変換できない文字列です")
	}
	fmt.Printf("%T %v", eee, eee)

	hh := "Hello World"
	fmt.Println(hh[0])
	// => 72
	// 先頭の文字hのアスキーコードのが表示される

	fmt.Println(string(hh[0]))
	// => h
}
