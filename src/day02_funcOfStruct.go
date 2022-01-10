package main

import (
	"fmt"
)

/*
	结构体的函数，想不明白为什么做的那么复杂
	好像没办法用直接 s.func = func(){}
*/

type structFunc struct {
	Name  string
	Power int
}

// 定义了 Super方法的接受者类型 为 *structFunc
func (s *structFunc) Super() {
	s.Power += 10000
}

func NewStruct(name string, power int) structFunc {
	return structFunc{
		Name:  name,
		Power: power,
	}
}

func main() {
	sf := &structFunc{"testFunc", 9000}
	sf.Super()
	fmt.Println(sf.Power)
}
