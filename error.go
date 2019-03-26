package main

import (
	"fmt"
	"log"
	"os"
)

// エラーハンドリングはtry catchではなく
// errがnいヵどうかではんていすることがおおい
func main() {
	file, err := os.Open("./defer.go")
	if err != nil {
		log.Fatal("Error")
	}

	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)

	// errが何度出てきてもerrは上書きされる
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	// os.Chdirは戻り値がerrのみなので
	// := ではなく = を使っている

	// ↑で :＝が使用できるのはerr以外にも戻り値があるためである

	// このように、errのみの場合は1行で書くことがおおい
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error")
	}
}
