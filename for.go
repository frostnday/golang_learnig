package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")

			// これが実行されたらこれ以降の処理を飛ばして次の処理に向かう
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
}
