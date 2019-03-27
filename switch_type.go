package main

import "fmt"

// 空のinterfaceを引数にする
// こうすることでどんな型の値も受け取ることができる
func do(i interface{}) {
	// iはinterfaceなのでこのままでは計算できない
	// ii := i * 2

	// タイプアサーション
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// 文字列でもintでもどちらでも対応できるようにする

	// スイッチタイプ文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!!")
	default:
		fmt.Printf("I dont know %T\n", v)
	}

}

func main() {
	// 空のinterfaceに10を入れてる感じ
	// var i interface{} = 10
	// do(i)

	do(10)
	do("Mike")
	do(true)

	// キャストでは無くタイプコンバージョンという
	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
