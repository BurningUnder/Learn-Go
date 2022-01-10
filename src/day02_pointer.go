package main

import "fmt"

/*
	在go里面也存在着指针传递实参的方法
	一般的函数只处理形参，获取的参数实际上是变量的副本，因此改变变量的副本并不改变原本变量
	需要改变原本变量则应使用指针来找到变量在内存中的位置进行修改，也就是使用&取地址符 和 * 指针操作
*/
type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{"power", 9000}
	Super(goku)
	fmt.Println(goku)
	gokuAdress := &Saiyan{"power", 9000}
	superS(gokuAdress)
	//	上面这个操作将修改变量的power
	fmt.Println(gokuAdress)
	//&{power 19000} fmt.println会将地址转化为数据
	print(gokuAdress)
	//0xc000004500 普通的print 则提取地址
	superSP(gokuAdress)
	//	上面这个操作什么也不会修改
	fmt.Println(gokuAdress)
	//&{power 19000}
}

func Super(s Saiyan) {
	s.Power += 10000
}

//当传递的参数是指针时，函数内的操作可以改变函数外的变量
//此时 goku := &Saiyan{} goku实际上是一个地址，函数参数标注的&会将其识别为指针
func superS(s *Saiyan) {
	s.Power += 10000
}

/*
	由于函数传参的机制是传递一个变量的副本
	所以即使修改地址的副本指向的值，也不会修改原来的变量
*/
func superSP(s *Saiyan) {
	s = &Saiyan{"changeAdreess", 1000}
	//	这里只是让这个副本的地址被指向了另一个变量
}
