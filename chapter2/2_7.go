package main

import (
	"fmt"
	"web_development_with_go/chapter2/strcon"
)

func main() {
	s := strcon.SwapCase("Gopher")
	fmt.Println("Converted string is :", s)
}
