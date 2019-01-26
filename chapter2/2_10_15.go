package main

import "fmt"

func main() {
	var x[5]int
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50
	fmt.Println(x)

	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println(y)

	z := [...]int{10, 20, 30, 40, 50}
	fmt.Println(z)

	w := [5]int{2: 10, 4: 40}
	fmt.Println(w)
}