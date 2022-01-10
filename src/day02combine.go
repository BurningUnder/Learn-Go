package main

import "fmt"

/*
	java中拓展结构使用的是继承，但是继承需要重构方法
	go提供了一种不需要重构的写法
*/

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Println("I'm %s\n", p.Name)
}

type goPerson struct {
	*Person
	Power int
}

func main() {
	goku := goPerson{
		Person: &Person{"goku"},
		Power:  9000,
	}
	goku.Introduce() //隐式使用
}
