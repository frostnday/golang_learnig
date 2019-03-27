package main

import "fmt"

// 独自の型を定義
type MyInt int

// 独自定義した型に独自のメソッドをもたせることもできる
func (i MyInt) Double() int {

	// ここはintにキャストしないとエラーになる
	// なぜなら型がMyInt型であるから !== intである
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
