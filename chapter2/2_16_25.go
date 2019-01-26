package main

import "fmt"

func main() {
	x := make([]int, 5, 10)
	fmt.Println(x)

	x = make([]int, 5)
	fmt.Println(x)

	x = []int{4:7}
	fmt.Println(x)

	x = []int{10,20,30}
	y := append(x, 40, 50)
	fmt.Println(x, y)

	x = []int{10, 20, 30}
	y = make([]int, 2)
	copy(y, x)
	fmt.Println(x, y)

	xx := make([]int, 2, 5)
	xx[0] = 10
	xx[1] = 20
	fmt.Println(xx)
	fmt.Println("Length is", len(xx))
	fmt.Println("Capacity is", cap(xx))

	xx = append(xx, 30, 40, 50)
	fmt.Println(xx)
	fmt.Println("Length is", len(xx))
	fmt.Println("Capacity is", cap(xx))

	xx = append(xx, 60)
	fmt.Println(xx)
	fmt.Println("Length is", len(xx))
	fmt.Println("Capacity is", cap(xx))

	x = []int{10,20,30,40,50}

	for k, v := range x {
		fmt.Printf("Index: %d Value: %d\n", k, v)
	}
}
