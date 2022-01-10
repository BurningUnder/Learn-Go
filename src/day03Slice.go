package main

import "fmt"

func main() {
	//	Slice就是不定义长度的数组，对slice的改变会反映到原本数组上
	//scores := []int{1,4,293,4,9}
	//虽然不定义长度，但是切片也是有长度的
	scores := make([]int, 0, 10) //这里定义了一个长度为0，容量为10 的切片
	fmt.Println(scores)
	//	切片长度为0 ，可以用其他方式来拓展切片
	scores = append(scores, 5)
	//	append会在切片容量不够的情况下创建新的切片并将原本切片的数据都复制到新切片上后append
	scores = scores[0:8]
	c := cap(scores)
	fmt.Println(scores, c)

}

/*
	names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10) -长度和容量都有效，可以直接用索引来获取数据
	var names []string	-当元素数量未知时与 append 连接
	scores := make([]int, 0, 20)
*/
