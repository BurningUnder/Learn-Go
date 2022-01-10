package main

func main() {
	println("it's day01")
	type Saiyan struct {
		Name  string
		Power int
	}
	goku := Saiyan{Name: "goku", Power: 9000}

	println(goku.Power)
}
