package main

import (
	"fmt"
)

func superFS(s *fieldsStruct) {
	s.Power += 10000
}

type funcSuper func(*fieldsStruct)

/*
	struct 里面的字段不能在创建后再定义
	使用type func 定义了一种函数类型，可以定义结构体内的函数了
	这里的操作时冗余的 直接 superFS(fs)就行了，不用fs.superFS(fs)
*/
type fieldsStruct struct {
	Super funcSuper
	Power int
}

func main() {
	fs := &fieldsStruct{
		Power: 9000,
	}
	fs.Super = superFS
	fs.Super(fs)
	//superFS(sf)
	fmt.Println(fs.Power)
}
