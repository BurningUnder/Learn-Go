package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 5)
	//copy(worst, scores[:5])
	//[X:] 是 从 X 到结尾 的简写，然而 [:X] 是 从开始到 X 的简写,不包含X
	copy(worst[2:4], scores[:5])
	fmt.Println(worst)
}
