package main

import (
	"fmt"
	"web_development_with_go/chapter1/calc"
)

func main(){
	var x, y int = 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
