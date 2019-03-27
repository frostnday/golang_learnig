package main

import "fmt"

// interface
type Human interface {
	Say() string
}

type Dog struct {
	Name string
}
type Person struct {
	Name string
}

func (p *Person) Say() {
	p.Name = "Mr. " + p.Name
	fmt.Println(p.Name)
}

func DriveCar(human Human) {
	// このように必ず特定のメソッドがないと困るときにInteefaceを使用する
	if human.Say() == "Mr Mike" {
		fmt.Println("Run")
	} else {
		fmt.Printlm("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}

	// これだとsaメソッドを実装していないのでエラーになる
	// var dog Dog = Dog{"dog"}
	// DriveCar(dog)
	DriveCar(mike)
}
