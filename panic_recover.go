package main

import "fmt"

func thirdPartyConnectDB() {
	// 例外を投げる
	// goはあまりpanicを推奨していない
	panic("Unable to connect database")
}

func save() {
	defer func() {
		// パニックが起きた後にrecoverでシステムを強制終了しないようにできる
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("ok?")
}
