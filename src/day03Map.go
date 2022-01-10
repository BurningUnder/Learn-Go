package main

import "fmt"

/*映射 键值对*/
func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]

	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)

	// returns 1
	total := len(lookup)
	fmt.Println(total)
	// has no return, can be called on a non-existing key
	delete(lookup, "goku")
}
