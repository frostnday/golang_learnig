package main

import "fmt"

// rangeとはforeachみたいなやつ
func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// foreachみたいにかける
	for i, v := range l {
		fmt.Println(i, v)
	}

	// foreachみたいにかける(要素数いらない)
	for _, v := range l {
		fmt.Println(v)
	}

	// 連想配列でももちろんforeachできるよ
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	// valueだけほしいときはkeyは _ としないといけない
	for _, v := range m {
		fmt.Println(v)
	}
}
