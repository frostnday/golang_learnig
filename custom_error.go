package main

import "fmt"

// エラーのstructを定義
type UserNotFound struct {
	Username string
}

// 定義したstructに対してerrorメソッドを実装する
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
	// return e.Username
}

// エラーを返すメソッド
func myFunc() error {
	ok := false
	if ok {
		return nil
	}

	// エラーだったらエラーのstructに値を格納
	return &UserNotFound{Username: "mike"}

	return nil
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
