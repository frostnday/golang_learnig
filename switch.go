package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()

	switch os {
	case "mac":
		fmt.Println("mac!!")

	case "windows":
		fmt.Println("windows!!")
	default:
		// 無くてもいい
		fmt.Println("Default!!")
	}

	// osという変数をこの処理の中でしか使用しないのであれば
	switch os2 := getOsName(); os2 {
	case "mac":
		fmt.Println("mac!!")

	case "windows":
		fmt.Println("windows!!")
	}

	t := time.Now()
	fmt.Println(t.Hour())

	// switchに変数の指定をしなくてもかける
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	}
}
